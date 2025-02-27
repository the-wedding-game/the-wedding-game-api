package apperrors

import (
	"errors"
	"testing"
)

func TestNewStorageError(t *testing.T) {
	storageError := NewStorageError("storage error")
	if storageError.Message != "storage error" {
		t.Errorf("expected storage error but got %s", storageError.Message)
	}
	if storageError.code != "StorageError" {
		t.Errorf("expected StorageError but got %s", storageError.code)
	}
	if !errors.As(storageError, &StorageError{}) {
		t.Errorf("expected true but got false")
	}

	storageError = NewStorageError("hello there")
	if storageError.Message != "hello there" {
		t.Errorf("expected hello there but got %s", storageError.Message)
	}
	if storageError.code != "StorageError" {
		t.Errorf("expected StorageError but got %s", storageError.code)
	}
	if !errors.As(storageError, &StorageError{}) {
		t.Errorf("expected true but got false")
	}
}

func TestStorageErrorMessage(t *testing.T) {
	storageError := NewStorageError("hello there")
	if storageError.Error() != "hello there" {
		t.Errorf("expected hello there but got %s", storageError.Error())
	}

	storageError = NewStorageError("random message")
	if storageError.Error() != "random message" {
		t.Errorf("expected storage error but got %s", storageError.Error())
	}
}
