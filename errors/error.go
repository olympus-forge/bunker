package bunker

type ApplicationError struct {
	Code           AppErrorCode `json:"code"`
	HttpStatusCode int          `json:"http_status_code"`
	ErrorCode      string       `json:"error_code"`
	Errors         error        `json:"error,omitempty"`
	Message        string       `json:"message"`
}

func New(code AppErrorCode) *ApplicationError {
	return getApplicationError(code)
}

func (ae *ApplicationError) SetError(err error) *ApplicationError {
	if err != nil {
		ae.Errors = err
	}
	return ae
}

func (ae *ApplicationError) SetMessage(msg string) *ApplicationError {
	if msg != "" {
		ae.Message = msg
	}
	return ae
}

func (ae *ApplicationError) Error() string {
	return ae.Message
}
