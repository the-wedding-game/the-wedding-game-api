package apperrors

import "testing"

func TestNewRecordNotFoundError(t *testing.T) {
	recordNotFoundError := NewRecordNotFoundError("record not found")
	if recordNotFoundError.Message != "record not found" {
		t.Errorf("expected record not found but got %s", recordNotFoundError.Message)
	}
	if recordNotFoundError.code != "RecordNotFoundError" {
		t.Errorf("expected RecordNotFoundError but got %s", recordNotFoundError.code)
	}
	if !IsRecordNotFoundError(recordNotFoundError) {
		t.Errorf("expected true but got false")
	}

	recordNotFoundError = NewRecordNotFoundError("hello there")
	if recordNotFoundError.Message != "hello there" {
		t.Errorf("expected hello there but got %s", recordNotFoundError.Message)
	}
	if recordNotFoundError.code != "RecordNotFoundError" {
		t.Errorf("expected RecordNotFoundError but got %s", recordNotFoundError.code)
	}
	if !IsRecordNotFoundError(recordNotFoundError) {
		t.Errorf("expected true but got false")
	}
}

func TestRecordNotFoundErrorMessage(t *testing.T) {
	recordNotFoundError := NewRecordNotFoundError("hello there")
	if recordNotFoundError.Error() != "hello there" {
		t.Errorf("expected hello there but got %s", recordNotFoundError.Error())
	}

	recordNotFoundError = NewRecordNotFoundError("random message")
	if recordNotFoundError.Error() != "random message" {
		t.Errorf("expected record not found but got %s", recordNotFoundError.Error())
	}
}

func TestIsRecordNotFoundError(t *testing.T) {
	recordNotFoundError := NewRecordNotFoundError("hello there")
	if !IsRecordNotFoundError(recordNotFoundError) {
		t.Errorf("expected true but got false")
	}

	recordNotFoundError = NewRecordNotFoundError("random message")
	if !IsRecordNotFoundError(recordNotFoundError) {
		t.Errorf("expected true but got false")
	}
}

func TestIsRecordNotFoundErrorNegative(t *testing.T) {
	authorizationError := NewAuthorizationError()
	if IsRecordNotFoundError(authorizationError) {
		t.Errorf("expected false but got true")
	}
}
