package apperrors

import (
	"errors"
)

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

func NewValidationError(message string) ValidationError {
	return ValidationError{Message: message}
}

func IsValidationError(err error) bool {
	var validationError ValidationError
	if errors.As(err, &validationError) {
		return true
	}

	return false
}
