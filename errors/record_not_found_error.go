package apperrors

import "errors"

type RecordNotFoundError struct {
	code    string
	Message string
}

func (e RecordNotFoundError) Error() string {
	return e.Message
}

func NewRecordNotFoundError(message string) RecordNotFoundError {
	return RecordNotFoundError{
		code:    "RecordNotFoundError",
		Message: message,
	}
}

func IsRecordNotFoundError(err error) bool {
	var recordNotFoundError RecordNotFoundError
	ok := errors.As(err, &recordNotFoundError)
	return ok
}
