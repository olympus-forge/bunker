package errors

type ApplicationError struct {
	Code           AppErrorCode `json:"code"`
	HttpStatusCode int          `json:"http_status_code"`
	ErrorCode      string       `json:"error_code"`
	Errors         string       `json:"error,omitempty"`
	Message        string       `json:"message"`
}

func New(code AppErrorCode) *ApplicationError {
	return newApplicationError(code)
}

func (ae *ApplicationError) SetErrorCode(code string) *ApplicationError {
	if ae != nil {
		ae.ErrorCode = code
	}
	return ae
}

func (ae *ApplicationError) SetMessage(msg string) *ApplicationError {
	if ae != nil {
		ae.Message = msg
	}
	return ae
}

func (ae *ApplicationError) SetErrors(err string) *ApplicationError {
	if ae != nil {
		ae.Errors = err
	}
	return ae
}

func (ae *ApplicationError) Error() string {
	if ae == nil {
		return ""
	}
	return ae.Message
}
