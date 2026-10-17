package errors

import (
	"net/http"
	"testing"
)

func TestBunkerError(t *testing.T) {

	errBadRequest := New(StatusBadRequest)

	if errBadRequest.Code != StatusBadRequest {
		t.Errorf("Invalid status code [%v] not equal bad requests", errBadRequest.Code)
	}

	if errBadRequest.HttpStatusCode != http.StatusBadRequest {
		t.Errorf("Invalid http status code [%v] not equal http status bad requests", errBadRequest.HttpStatusCode)
	}

	errBadRequest.SetErrorCode("IE-01X")
	if errBadRequest.ErrorCode != "IE-01X" {
		t.Error("invalid error code")
	}

	msgError := BadRequest.Message
	errBadRequest.SetMessage(msgError)
	if errBadRequest.Message != msgError {
		t.Error("invalid message error")
	}

	newMsgError := "new error"
	errBadRequest.SetMessage(newMsgError)
	if errBadRequest.Message != newMsgError {
		t.Error("invalid message")
	}

	// Test SetErrors
	errBadRequest.SetErrors("underlying db failure")
	if errBadRequest.Errors != "underlying db failure" {
		t.Error("invalid errors field")
	}

	// Test global mutation isolation: mutating BadRequest variable must NOT affect New(StatusBadRequest)
	BadRequest.SetMessage("mutated globally")
	freshErr := New(StatusBadRequest)
	if freshErr.Message != "Bad Request" {
		t.Errorf("expected fresh error to have 'Bad Request', got: %s", freshErr.Message)
	}

	// Test 5xx server errors
	errInternal := New(StatusInternalServerError)
	if errInternal.HttpStatusCode != http.StatusInternalServerError || errInternal.ErrorCode != "IE-30" {
		t.Errorf("invalid 500 error: %+v", errInternal)
	}

	// Test nil receiver safety
	var nilErr *ApplicationError
	if nilErr.Error() != "" {
		t.Errorf("expected empty string from nil error, got: %s", nilErr.Error())
	}
}
