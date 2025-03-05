package apperrors

import (
	"errors"
)

type NotFoundError struct {
	code    string
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

func NewNotFoundError(entity string, key string) NotFoundError {
	return NotFoundError{
		code:    "NotFoundError",
		Message: entity + " with key " + key + " not found.",
	}
}

func IsNotFoundError(err error) bool {
	var notFoundError NotFoundError
	ok := errors.As(err, &notFoundError)
	return ok
}
