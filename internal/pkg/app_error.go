package pkg

type Error struct {
	StatusCode int
	Code       int
	Message    string
}

func NewError(statusCode, code int, message string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
	}
}

func (e *Error) Error() string {
	return e.Message
}
