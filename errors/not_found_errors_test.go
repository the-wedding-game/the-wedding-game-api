package apperrors

import "testing"

func TestCreateNotFoundError(t *testing.T) {
	notFoundError := NewNotFoundError("entity", "key")
	if notFoundError.Message != "entity with key key not found." {
		t.Errorf("expected entity with key key not found. but got %s", notFoundError.Message)
	}
	if notFoundError.code != "NotFoundError" {
		t.Errorf("expected NotFoundError but got %s", notFoundError.code)
	}
	if !IsNotFoundError(notFoundError) {
		t.Errorf("expected true but got false")
	}

	notFoundError = NewNotFoundError("fake entity", "fake key")
	if notFoundError.Message != "fake entity with key fake key not found." {
		t.Errorf("expected fake entity with key fake key not found. but got %s", notFoundError.Message)
	}
	if notFoundError.code != "NotFoundError" {
		t.Errorf("expected NotFoundError but got %s", notFoundError.code)
	}
	if !IsNotFoundError(notFoundError) {
		t.Errorf("expected true but got false")
	}
}

func TestNotFoundErrorMessage(t *testing.T) {
	notFoundError := NewNotFoundError("test entity", "test key")
	if notFoundError.Error() != "test entity with key test key not found." {
		t.Errorf("expected test entity with key test key not found. but got %s", notFoundError.Error())
	}
}

func TestIsNotFoundError(t *testing.T) {
	notFoundError := NewNotFoundError("test entity", "test key")
	if !IsNotFoundError(notFoundError) {
		t.Errorf("expected true but got false")
	}

	notFoundError = NewNotFoundError("another entity", "another key")
	if !IsNotFoundError(notFoundError) {
		t.Errorf("expected true but got false")
	}
}

func TestIsNotFoundErrorNegative(t *testing.T) {
	authorizationError := NewAuthorizationError()
	if IsNotFoundError(authorizationError) {
		t.Errorf("expected false but got true")
	}
}
