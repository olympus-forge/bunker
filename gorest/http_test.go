package gorest

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type samplePayload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestBunkerHttp(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/echo-headers":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"origin": r.Header["X-App-Origin"],
				"auth":   r.Header.Get("Authorization"),
				"query":  r.URL.Query().Get("data"),
			})
		case "/post-json":
			body, _ := io.ReadAll(r.Body)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write(body)
		default:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		}
	}))
	defer ts.Close()

	t.Run("GET with Debug and Body Reading", func(t *testing.T) {
		req := New(ts.URL).Get().
			AddPath("/echo-headers").
			SetDebug(true).
			Query(map[string]string{"data": `{"id":1}`}).
			SetHeader("X-App-Origin", "xman", "wolverine").
			SetToken("secret-token").
			Do()

		if req.HaveError() {
			t.Fatalf("unexpected error: %v", req.Errors)
		}

		// Body should still be readable even when Debug was enabled
		bodyStr := req.BodyString()
		if bodyStr == "" {
			t.Fatal("expected non-empty body string")
		}

		var result map[string]interface{}
		if err := req.Bind(&result); err != nil {
			t.Fatalf("failed to bind json: %v", err)
		}

		if result["auth"] != "Bearer secret-token" {
			t.Errorf("expected bearer token, got: %v", result["auth"])
		}
	})

	t.Run("POST with Pointer to Struct and Slices", func(t *testing.T) {
		payload := &samplePayload{Name: "Gopher", Age: 10}
		req := New(ts.URL).Post().
			AddPath("/post-json").
			SetPayload(payload).
			Do()

		if req.HaveError() {
			t.Fatalf("unexpected error: %v", req.Errors)
		}

		var echoed samplePayload
		if err := req.Bind(&echoed); err != nil {
			t.Fatalf("bind failed: %v", err)
		}
		if echoed.Name != "Gopher" || echoed.Age != 10 {
			t.Errorf("unexpected echoed payload: %+v", echoed)
		}

		// Slice payload
		slicePayload := []string{"item1", "item2"}
		sliceReq := New(ts.URL).Post().
			AddPath("/post-json").
			SetPayload(slicePayload).
			Do()

		if sliceReq.HaveError() {
			t.Fatalf("unexpected slice error: %v", sliceReq.Errors)
		}
		var echoedSlice []string
		if err := sliceReq.Bind(&echoedSlice); err != nil {
			t.Fatalf("bind slice failed: %v", err)
		}
		if len(echoedSlice) != 2 || echoedSlice[0] != "item1" {
			t.Errorf("unexpected slice result: %v", echoedSlice)
		}
	})

	t.Run("Headers and Query Options", func(t *testing.T) {
		req := New(ts.URL).Get().
			SetHeaders(map[string]string{"X-Custom": "val1"}).
			Query("?page=2&limit=50")

		if req.HaveError() {
			t.Fatalf("header/query error: %v", req.Errors)
		}
		if req.QueryData.Get("page") != "2" || req.QueryData.Get("limit") != "50" {
			t.Errorf("query param parsing failed: %v", req.QueryData)
		}
		if req.Header["X-Custom"][0] != "val1" {
			t.Errorf("custom header not set: %v", req.Header)
		}
	})

	t.Run("Unsupported Method Does Not Panic With Debug", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("panic caught: %v", r)
			}
		}()
		req := New(ts.URL).SetDebug(true)
		req.Method = "UNKNOWN_METHOD"
		req.Do()

		if !req.HaveError() {
			t.Errorf("expected error for unknown method")
		}
	})
}
