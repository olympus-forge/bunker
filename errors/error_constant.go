package errors

import "net/http"

type AppErrorCode int

const (
	StatusBadRequest AppErrorCode = iota + 1
	StatusUnAuthorized
	StatusPaymentRequired
	StatusForbidden
	StatusNotFound
	StatusMethodNotAllowed
	StatusNotAcceptable
	StatusProxyAuthRequired
	StatusRequestTimeout
	StatusConflict
	StatusGone
	StatusLengthRequired
	StatusPreconditionFailed
	StatusRequestEntityTooLarge
	StatusRequestURITooLong
	StatusUnsupportedMediaType
	StatusRequestedRangeNotSatisfiable
	StatusExpectationFailed
	StatusTeapot
	StatusMisdirectedRequest
	StatusUnprocessableEntity
	StatusLocked
	StatusFailedDependency
	StatusTooEarly
	StatusUpgradeRequired
	StatusPreconditionRequired
	StatusTooManyRequests
	StatusRequestHeaderFieldsTooLarge
	StatusUnavailableForLegalReasons
	StatusInternalServerError
	StatusNotImplemented
	StatusBadGateway
	StatusServiceUnavailable
	StatusGatewayTimeout
)

func httpStatusCode(code AppErrorCode) int {
	switch code {
	case StatusBadRequest:
		return http.StatusBadRequest
	case StatusUnAuthorized:
		return http.StatusUnauthorized
	case StatusPaymentRequired:
		return http.StatusPaymentRequired
	case StatusForbidden:
		return http.StatusForbidden
	case StatusNotFound:
		return http.StatusNotFound
	case StatusMethodNotAllowed:
		return http.StatusMethodNotAllowed
	case StatusNotAcceptable:
		return http.StatusNotAcceptable
	case StatusProxyAuthRequired:
		return http.StatusProxyAuthRequired
	case StatusRequestTimeout:
		return http.StatusRequestTimeout
	case StatusConflict:
		return http.StatusConflict
	case StatusGone:
		return http.StatusGone
	case StatusLengthRequired:
		return http.StatusLengthRequired
	case StatusPreconditionFailed:
		return http.StatusPreconditionFailed
	case StatusRequestEntityTooLarge:
		return http.StatusRequestEntityTooLarge
	case StatusRequestURITooLong:
		return http.StatusRequestURITooLong
	case StatusUnsupportedMediaType:
		return http.StatusUnsupportedMediaType
	case StatusRequestedRangeNotSatisfiable:
		return http.StatusRequestedRangeNotSatisfiable
	case StatusExpectationFailed:
		return http.StatusExpectationFailed
	case StatusTeapot:
		return http.StatusTeapot
	case StatusMisdirectedRequest:
		return http.StatusMisdirectedRequest
	case StatusUnprocessableEntity:
		return http.StatusUnprocessableEntity
	case StatusLocked:
		return http.StatusLocked
	case StatusFailedDependency:
		return http.StatusFailedDependency
	case StatusTooEarly:
		return http.StatusTooEarly
	case StatusUpgradeRequired:
		return http.StatusUpgradeRequired
	case StatusPreconditionRequired:
		return http.StatusPreconditionRequired
	case StatusTooManyRequests:
		return http.StatusTooManyRequests
	case StatusRequestHeaderFieldsTooLarge:
		return http.StatusRequestHeaderFieldsTooLarge
	case StatusUnavailableForLegalReasons:
		return http.StatusUnavailableForLegalReasons
	case StatusInternalServerError:
		return http.StatusInternalServerError
	case StatusNotImplemented:
		return http.StatusNotImplemented
	case StatusBadGateway:
		return http.StatusBadGateway
	case StatusServiceUnavailable:
		return http.StatusServiceUnavailable
	case StatusGatewayTimeout:
		return http.StatusGatewayTimeout
	default:
		return 0
	}
}

func errorCode(code AppErrorCode) string {
	switch code {
	case StatusBadRequest:
		return "IE-01"
	case StatusUnAuthorized:
		return "IE-02"
	case StatusPaymentRequired:
		return "IE-03"
	case StatusForbidden:
		return "IE-04"
	case StatusNotFound:
		return "IE-05"
	case StatusMethodNotAllowed:
		return "IE-06"
	case StatusNotAcceptable:
		return "IE-07"
	case StatusProxyAuthRequired:
		return "IE-08"
	case StatusRequestTimeout:
		return "IE-09"
	case StatusConflict:
		return "IE-10"
	case StatusGone:
		return "IE-11"
	case StatusLengthRequired:
		return "IE-12"
	case StatusPreconditionFailed:
		return "IE-13"
	case StatusRequestEntityTooLarge:
		return "IE-14"
	case StatusRequestURITooLong:
		return "IE-15"
	case StatusUnsupportedMediaType:
		return "IE-16"
	case StatusRequestedRangeNotSatisfiable:
		return "IE-17"
	case StatusExpectationFailed:
		return "IE-18"
	case StatusTeapot:
		return "IE-19"
	case StatusMisdirectedRequest:
		return "IE-20"
	case StatusUnprocessableEntity:
		return "IE-21"
	case StatusLocked:
		return "IE-22"
	case StatusFailedDependency:
		return "IE-23"
	case StatusTooEarly:
		return "IE-24"
	case StatusUpgradeRequired:
		return "IE-25"
	case StatusPreconditionRequired:
		return "IE-26"
	case StatusTooManyRequests:
		return "IE-27"
	case StatusRequestHeaderFieldsTooLarge:
		return "IE-28"
	case StatusUnavailableForLegalReasons:
		return "IE-29"
	case StatusInternalServerError:
		return "IE-30"
	case StatusNotImplemented:
		return "IE-31"
	case StatusBadGateway:
		return "IE-32"
	case StatusServiceUnavailable:
		return "IE-33"
	case StatusGatewayTimeout:
		return "IE-34"
	default:
		return ""
	}
}

func message(code AppErrorCode) string {
	switch code {
	case StatusBadRequest:
		return "Bad Request"
	case StatusUnAuthorized:
		return "Not Authorized"
	case StatusPaymentRequired:
		return "Payment Required"
	case StatusForbidden:
		return "Forbidden"
	case StatusNotFound:
		return "Not Found"
	case StatusMethodNotAllowed:
		return "Method Not Allowed"
	case StatusNotAcceptable:
		return "Not Acceptable"
	case StatusProxyAuthRequired:
		return "Proxy Auth Required"
	case StatusRequestTimeout:
		return "Request Timeout"
	case StatusConflict:
		return "Conflict"
	case StatusGone:
		return "Gone"
	case StatusLengthRequired:
		return "Length Required"
	case StatusPreconditionFailed:
		return "Precondition Failed"
	case StatusRequestEntityTooLarge:
		return "Request Entity Too Large"
	case StatusRequestURITooLong:
		return "Request URI Too Long"
	case StatusUnsupportedMediaType:
		return "Unsupported Media Type"
	case StatusRequestedRangeNotSatisfiable:
		return "Requested Range Not Satisfiable"
	case StatusExpectationFailed:
		return "Expectation Failed"
	case StatusTeapot:
		return "Teapot"
	case StatusMisdirectedRequest:
		return "Misdirected Request"
	case StatusUnprocessableEntity:
		return "Unprocessable Entity"
	case StatusLocked:
		return "Locked"
	case StatusFailedDependency:
		return "Failed Dependency"
	case StatusTooEarly:
		return "Too Early"
	case StatusUpgradeRequired:
		return "Upgrade Required"
	case StatusPreconditionRequired:
		return "Precondition Required"
	case StatusTooManyRequests:
		return "Too Many Requests"
	case StatusRequestHeaderFieldsTooLarge:
		return "Request Header Fields Too Large"
	case StatusUnavailableForLegalReasons:
		return "Unavailable For Legal Reasons"
	case StatusInternalServerError:
		return "Internal Server Error"
	case StatusNotImplemented:
		return "Not Implemented"
	case StatusBadGateway:
		return "Bad Gateway"
	case StatusServiceUnavailable:
		return "Service Unavailable"
	case StatusGatewayTimeout:
		return "Gateway Timeout"
	}
	return ""
}

func newApplicationError(code AppErrorCode) *ApplicationError {
	statusCode := httpStatusCode(code)
	if statusCode == 0 {
		return nil
	}
	return &ApplicationError{
		Code:           code,
		HttpStatusCode: statusCode,
		ErrorCode:      errorCode(code),
		Message:        message(code),
	}
}

var (
	BadRequest                   = newApplicationError(StatusBadRequest)
	UnAuthorized                 = newApplicationError(StatusUnAuthorized)
	PaymentRequired              = newApplicationError(StatusPaymentRequired)
	Forbidden                    = newApplicationError(StatusForbidden)
	NotFound                     = newApplicationError(StatusNotFound)
	MethodNotAllowed             = newApplicationError(StatusMethodNotAllowed)
	NotAcceptable                = newApplicationError(StatusNotAcceptable)
	ProxyAuthRequired            = newApplicationError(StatusProxyAuthRequired)
	RequestTimeout               = newApplicationError(StatusRequestTimeout)
	Conflict                     = newApplicationError(StatusConflict)
	Gone                         = newApplicationError(StatusGone)
	LengthRequired               = newApplicationError(StatusLengthRequired)
	PreconditionFailed           = newApplicationError(StatusPreconditionFailed)
	RequestEntityTooLarge        = newApplicationError(StatusRequestEntityTooLarge)
	RequestURITooLong            = newApplicationError(StatusRequestURITooLong)
	UnsupportedMediaType         = newApplicationError(StatusUnsupportedMediaType)
	RequestedRangeNotSatisfiable = newApplicationError(StatusRequestedRangeNotSatisfiable)
	ExpectationFailed            = newApplicationError(StatusExpectationFailed)
	Teapot                       = newApplicationError(StatusTeapot)
	MisdirectedRequest           = newApplicationError(StatusMisdirectedRequest)
	UnprocessableEntity          = newApplicationError(StatusUnprocessableEntity)
	Locked                       = newApplicationError(StatusLocked)
	FailedDependency             = newApplicationError(StatusFailedDependency)
	TooEarly                     = newApplicationError(StatusTooEarly)
	UpgradeRequired              = newApplicationError(StatusUpgradeRequired)
	PreconditionRequired         = newApplicationError(StatusPreconditionRequired)
	TooManyRequests              = newApplicationError(StatusTooManyRequests)
	RequestHeaderFieldsTooLarge  = newApplicationError(StatusRequestHeaderFieldsTooLarge)
	UnavailableForLegalReasons   = newApplicationError(StatusUnavailableForLegalReasons)
	InternalServerError          = newApplicationError(StatusInternalServerError)
	NotImplemented               = newApplicationError(StatusNotImplemented)
	BadGateway                   = newApplicationError(StatusBadGateway)
	ServiceUnavailable           = newApplicationError(StatusServiceUnavailable)
	GatewayTimeout               = newApplicationError(StatusGatewayTimeout)
)
