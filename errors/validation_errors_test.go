package apperrors

import "testing"

func TestValidationError(t *testing.T) {
	validationError := NewValidationError("validation failed")
	if validationError.Message != "validation failed" {
		t.Errorf("expected validation failed but got %s", validationError.Message)
	}
	if validationError.code != "ValidationError" {
		t.Errorf("expected ValidationError but got %s", validationError.code)
	}
	if !IsValidationError(validationError) {
		t.Errorf("expected true but got false")
	}

	validationError = NewValidationError("hello there")
	if validationError.Message != "hello there" {
		t.Errorf("expected hello there but got %s", validationError.Message)
	}
	if validationError.code != "ValidationError" {
		t.Errorf("expected ValidationError but got %s", validationError.code)
	}
	if !IsValidationError(validationError) {
		t.Errorf("expected true but got false")
	}
}

func TestValidationErrorMessage(t *testing.T) {
	validationError := NewValidationError("hello there")
	if validationError.Error() != "hello there" {
		t.Errorf("expected hello there but got %s", validationError.Error())
	}

	validationError = NewValidationError("random message")
	if validationError.Error() != "random message" {
		t.Errorf("expected validation failed but got %s", validationError.Error())
	}
}

func TestIsValidationError(t *testing.T) {
	validationError := NewValidationError("hello there")
	if !IsValidationError(validationError) {
		t.Errorf("expected true but got false")
	}

	validationError = NewValidationError("random message")
	if !IsValidationError(validationError) {
		t.Errorf("expected true but got false")
	}
}

func TestIsValidationErrorNegative(t *testing.T) {
	authorizationError := NewAuthorizationError()
	if IsValidationError(authorizationError) {
		t.Errorf("expected false but got true")
	}
}
