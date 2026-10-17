package gorest

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/olympus-forge/bunker"

	"golang.org/x/net/publicsuffix"
)

// HTTP methods we support
const (
	POST    string = http.MethodPost
	GET     string = http.MethodGet
	HEAD    string = http.MethodHead
	PUT     string = http.MethodPut
	DELETE  string = http.MethodDelete
	PATCH   string = http.MethodPatch
	OPTIONS string = http.MethodOptions
)

// ShortContentTypes defines some short content types.
const (
	Json       string = "application/json"
	UrlEncoded string = "application/x-www-form-urlencoded"
	Form       string = "multipart/form-data"
	FormData   string = "multipart/form-data"
)

const (
	Auth   string = "Authorization"
	Bearer string = "Bearer "
)

type Requester struct {
	path    string
	BaseUrl string
	Method  string

	token         string
	Header        map[string][]string
	basicAuth     map[string]string
	hasBasicAuth  bool
	basicAuthUser string
	basicAuthPass string

	FormData  url.Values
	QueryData url.Values

	Client   *http.Client
	Response *http.Response
	Request  *http.Request

	TimeOut time.Duration

	Errors []error

	Context context.Context

	Debug              bool
	insecureSkipVerify bool

	timeRequest time.Duration
	timeIn      time.Time

	body      io.Reader
	bodyBytes []byte
}

func New(host string) *Requester {
	return &Requester{
		BaseUrl:   host,
		Header:    make(map[string][]string),
		FormData:  make(url.Values),
		QueryData: make(url.Values),
		basicAuth: make(map[string]string),
	}
}

func (base *Requester) HaveError() bool {
	return len(base.Errors) != 0
}

func (base *Requester) SetHeader(param string, values ...string) *Requester {
	base.Header[param] = values
	return base
}

func (base *Requester) SetContext(ctx context.Context) *Requester {
	base.Context = ctx
	return base
}

func (base *Requester) SetBasicAuth(username, password string) *Requester {
	base.hasBasicAuth = true
	base.basicAuthUser = username
	base.basicAuthPass = password
	base.basicAuth[username] = password
	return base
}

func (base *Requester) SetToken(token string) *Requester {
	base.token = token
	return base
}

func (base *Requester) SetHeaders(headers interface{}) *Requester {
	if headers == nil {
		return base
	}
	switch h := headers.(type) {
	case map[string]string:
		for k, v := range h {
			base.Header[k] = []string{v}
		}
	case map[string][]string:
		for k, v := range h {
			base.Header[k] = v
		}
	case http.Header:
		for k, v := range h {
			base.Header[k] = v
		}
	case map[string]interface{}:
		for k, v := range h {
			switch val := v.(type) {
			case string:
				base.Header[k] = []string{val}
			case []string:
				base.Header[k] = val
			default:
				base.Header[k] = []string{fmt.Sprint(val)}
			}
		}
	default:
		base.Errors = append(base.Errors, fmt.Errorf("unsupported type of %T", headers))
	}
	return base
}

func (base *Requester) setMapStringHeaders(headers map[string][]string) *Requester {
	for k, v := range headers {
		base.Header[k] = v
	}
	return base
}

func (base *Requester) setMapAnyHeaders(headers map[string]interface{}) *Requester {
	for k, v := range headers {
		switch value := v.(type) {
		case string:
			base.Header[k] = []string{value}
		case []string:
			base.Header[k] = value
		default:
			base.Header[k] = []string{fmt.Sprint(value)}
		}
	}
	return base
}

func (base *Requester) SetContentType(contentType string) *Requester {
	if bunker.IsEmptyString(contentType) {
		base.Header["Content-Type"] = []string{Json}
	} else {
		base.Header["Content-Type"] = []string{contentType}
	}
	return base
}

func (base *Requester) initClient() *Requester {
	if base.Client != nil {
		base.Client.Timeout = base.TimeOut
		return base
	}
	jar, errCookie := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if errCookie != nil {
		base.Errors = append(base.Errors, errCookie)
		return base
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: base.insecureSkipVerify},
		},
		Jar:     jar,
		Timeout: base.TimeOut,
	}
	base.Client = client
	return base
}

func (base *Requester) SetTimeout(timeOut time.Duration) *Requester {
	base.TimeOut = timeOut
	return base
}

func (base *Requester) initRequestClient() bool {
	if bunker.IsEmptyString(base.Method) {
		base.Errors = append(base.Errors, errors.New("the request method has not been selected"))
		return false
	}

	if base.initRequest().HaveError() {
		return false
	}

	if (base.body != nil || len(base.bodyBytes) > 0) && bunker.IsEmptyString(base.Request.Header.Get("Content-Type")) {
		base.Request.Header.Set("Content-Type", Json)
	}

	if base.initClient().HaveError() {
		return false
	}

	return true
}

func (base *Requester) Do() *Requester {
	defer func() {
		if base.HaveError() && (base.Debug || os.Getenv("debug") == "true") {
			bunker.PrintErr(fmt.Sprint(base.Errors))
		}
	}()

	startTime := time.Now()
	if !base.initRequestClient() {
		return base
	}

	switch base.Method {
	case GET, HEAD, DELETE, OPTIONS, POST, PUT, PATCH:
		response, errRequestClient := base.Client.Do(base.Request)
		if errRequestClient != nil {
			base.Errors = append(base.Errors, errRequestClient)
			return base
		}
		base.Response = response
	default:
		base.Errors = append(base.Errors, fmt.Errorf("unsupported method of %s", base.Method))
		return base
	}
	base.timeRequest = time.Since(startTime)
	base.timeIn = time.Now()
	base.printDebug()
	return base
}

func (base *Requester) initRequest() *Requester {
	var bodyReader io.Reader
	if len(base.bodyBytes) > 0 {
		bodyReader = bytes.NewReader(base.bodyBytes)
	} else if base.body != nil {
		bodyReader = base.body
	}

	request, errRequest := http.NewRequest(base.Method, base.BaseUrl, bodyReader)
	if errRequest != nil {
		base.Errors = append(base.Errors, errRequest)
		return base
	}
	if base.Context != nil {
		request = request.WithContext(base.Context)
	}
	if base.Header != nil {
		for k, v := range base.Header {
			request.Header[k] = v
		}
	}
	if base.hasBasicAuth {
		request.SetBasicAuth(base.basicAuthUser, base.basicAuthPass)
	} else if len(base.basicAuth) > 0 {
		for user, pass := range base.basicAuth {
			request.SetBasicAuth(user, pass)
			break
		}
	}
	if !bunker.IsEmptyString(base.token) {
		request.Header.Set(Auth, Bearer+base.token)
	}

	reqUrl := request.URL.Query()
	for param, values := range base.QueryData {
		for _, value := range values {
			reqUrl.Add(param, value)
		}
	}
	request.URL.RawQuery = reqUrl.Encode()

	if !bunker.IsEmptyString(base.path) {
		p := base.path
		if bunker.IsEmptyString(request.URL.Path) || request.URL.Path == "/" {
			if !strings.HasPrefix(p, "/") {
				request.URL.Path = "/" + p
			} else {
				request.URL.Path = p
			}
		} else {
			request.URL.Path = strings.TrimRight(request.URL.Path, "/") + "/" + strings.TrimLeft(p, "/")
		}
	}
	base.Request = request
	return base
}

func (base *Requester) Get() *Requester {
	base.Method = GET
	return base
}

func (base *Requester) Post() *Requester {
	base.Method = POST
	return base
}

func (base *Requester) Put() *Requester {
	base.Method = PUT
	return base
}

func (base *Requester) Patch() *Requester {
	base.Method = PATCH
	return base
}

func (base *Requester) Delete() *Requester {
	base.Method = DELETE
	return base
}

func (base *Requester) Head() *Requester {
	base.Method = HEAD
	return base
}

func (base *Requester) Options() *Requester {
	base.Method = OPTIONS
	return base
}

func (base *Requester) SetDebug(debug bool) *Requester {
	base.Debug = debug
	return base
}

func (base *Requester) SkipVerify(verify bool) *Requester {
	base.insecureSkipVerify = verify
	return base
}

func (base *Requester) printDebug() {
	if base.Debug || os.Getenv("debug") == "true" {
		debugMessage := base.debug()
		if !base.HaveError() {
			bunker.LogInfo(debugMessage)
		}
	}
}

func (base *Requester) readAll(input io.ReadCloser) []byte {
	if input != nil {
		result, err := io.ReadAll(input)
		if err != nil {
			base.Errors = append(base.Errors, err)
			return nil
		}
		return result
	}
	return nil
}

func (base *Requester) readBodyRequest() []byte {
	if len(base.bodyBytes) > 0 {
		return base.bodyBytes
	}
	if base.Request != nil && base.Request.Body != nil {
		result, err := io.ReadAll(base.Request.Body)
		if err != nil {
			base.Errors = append(base.Errors, err)
			return nil
		}
		base.Request.Body = io.NopCloser(bytes.NewBuffer(result))
		base.bodyBytes = result
		return result
	}
	return nil
}

func (base *Requester) SetPayload(body interface{}) *Requester {
	if body == nil {
		return base
	}
	switch v := body.(type) {
	case string:
		base.bodyBytes = []byte(v)
		base.body = strings.NewReader(v)
	case []byte:
		base.bodyBytes = v
		base.body = bytes.NewReader(v)
	case io.Reader:
		data, err := io.ReadAll(v)
		if err != nil {
			base.Errors = append(base.Errors, err)
			return base
		}
		base.bodyBytes = data
		base.body = bytes.NewReader(data)
	default:
		dataMarshal, errMarshal := json.Marshal(body)
		if errMarshal != nil {
			base.Errors = append(base.Errors, errMarshal)
			return base
		}
		base.bodyBytes = dataMarshal
		base.body = bytes.NewReader(dataMarshal)
	}
	return base
}

func (base *Requester) AddPath(path string) *Requester {
	if bunker.IsEmptyString(path) {
		return base
	}
	if bunker.IsEmptyString(base.path) {
		base.path = path
	} else {
		base.path = strings.TrimRight(base.path, "/") + "/" + strings.TrimLeft(path, "/")
	}
	return base
}

func (base *Requester) Query(query interface{}) *Requester {
	if query == nil {
		return base
	}
	switch input := query.(type) {
	case string:
		base.queryString(input)
	case url.Values:
		for k, vals := range input {
			for _, v := range vals {
				base.QueryData.Add(k, v)
			}
		}
	case map[string][]string:
		for k, vals := range input {
			for _, v := range vals {
				base.QueryData.Add(k, v)
			}
		}
	case map[string]string:
		base.queryMap(input)
	case map[string]interface{}:
		for k, v := range input {
			base.QueryData.Add(k, fmt.Sprint(v))
		}
	default:
		switch reflect.ValueOf(query).Kind() {
		case reflect.String:
			base.queryString(fmt.Sprint(query))
		case reflect.Map:
			base.Errors = append(base.Errors, fmt.Errorf("unsupported type of %T", query))
		default:
			base.Errors = append(base.Errors, fmt.Errorf("unsupported type of %T", query))
		}
	}
	return base
}

func (base *Requester) queryString(input string) *Requester {
	cleanInput := strings.TrimPrefix(input, "?")
	values, errParse := url.ParseQuery(cleanInput)
	if errParse != nil {
		base.Errors = append(base.Errors, errParse)
		return base
	}
	for param, vals := range values {
		for _, v := range vals {
			base.QueryData.Add(param, v)
		}
	}
	return base
}

func (base *Requester) queryMap(input map[string]string) *Requester {
	for k, v := range input {
		base.QueryData.Add(k, v)
	}
	return base
}

func (base *Requester) setBodyMap(input interface{}) {
	base.SetPayload(input)
}

func (base *Requester) setBodyMapString(input interface{}) *Requester {
	return base.SetPayload(input)
}

func (base *Requester) setBodyStruct(input interface{}) *Requester {
	return base.SetPayload(input)
}

func (base *Requester) setBodyString(input interface{}) *Requester {
	return base.SetPayload(input)
}

func (base *Requester) debug() string {
	if base.Request != nil {
		var responseBody string
		if base.Response != nil && base.Response.Body != nil {
			bodyBytes := base.readAll(base.Response.Body)
			responseBody = string(bodyBytes)
			base.Response.Body.Close()
			base.Response.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
		status := "No Response"
		if base.Response != nil {
			status = base.Response.Status
		}
		return fmt.Sprintf(
			`
	REQUEST
	==========================================================
	%s / %s / %s
	URL             : %s
	HEADERS         : %v
	BODY REQUEST    : 
	%v

	RESPONSE
	==========================================================
	STATUS          : %s
	RECEIVED AT     : %v
	RESPONSE TIME   : %v
	
	BODY RESPONSE   :
	%v

	==========================================================
	`,
			time.Now().Format("2006/01/02 15:04:05"),
			base.Method,
			base.Request.Proto,
			base.Request.URL,
			base.headerToString(),
			string(base.readBodyRequest()),
			status,
			base.timeIn.Format(time.RFC1123),
			base.timeRequest.Seconds(),
			responseBody,
		)
	} else {
		base.Errors = append(base.Errors, errors.New("can't debug before request"))
		return ""
	}
}

func (base *Requester) headerToString() (result string) {
	if base.Request != nil && base.Request.Header != nil {
		header, _ := json.Marshal(base.Request.Header)
		return string(header)
	}
	return ""
}

// BodyBytes returns the response body as bytes, preserving the body for further reads.
func (base *Requester) BodyBytes() []byte {
	if base.Response == nil || base.Response.Body == nil {
		return nil
	}
	data, err := io.ReadAll(base.Response.Body)
	if err != nil {
		base.Errors = append(base.Errors, err)
		return nil
	}
	base.Response.Body.Close()
	base.Response.Body = io.NopCloser(bytes.NewBuffer(data))
	return data
}

// BodyString returns the response body as a string.
func (base *Requester) BodyString() string {
	return string(base.BodyBytes())
}

// Bind unmarshals the response JSON payload into dest.
func (base *Requester) Bind(dest interface{}) error {
	data := base.BodyBytes()
	if len(data) == 0 {
		return errors.New("empty response body")
	}
	return json.Unmarshal(data, dest)
}
