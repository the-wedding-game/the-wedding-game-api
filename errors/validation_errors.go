package apperrors

import (
	"errors"
)

type ValidationError struct {
	code    string
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

func NewValidationError(message string) ValidationError {
	return ValidationError{
		code:    "ValidationError",
		Message: message,
	}
}

func IsValidationError(err error) bool {
	var validationError ValidationError
	if errors.As(err, &validationError) {
		return true
	}

	return false
}
