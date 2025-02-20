package apperrors

import "testing"

func TestNewDatabaseError(t *testing.T) {
	databaseError := NewDatabaseError("database error")
	if databaseError.Message != "database error" {
		t.Errorf("expected database error but got %s", databaseError.Message)
	}
	if databaseError.code != "DatabaseError" {
		t.Errorf("expected DatabaseError but got %s", databaseError.code)
	}
	if !IsDatabaseError(databaseError) {
		t.Errorf("expected true but got false")
	}

	databaseError = NewDatabaseError("hello there")
	if databaseError.Message != "hello there" {
		t.Errorf("expected hello there but got %s", databaseError.Message)
	}
	if databaseError.code != "DatabaseError" {
		t.Errorf("expected DatabaseError but got %s", databaseError.code)
	}
	if !IsDatabaseError(databaseError) {
		t.Errorf("expected true but got false")
	}
}

func TestDatabaseErrorMessage(t *testing.T) {
	databaseError := NewDatabaseError("hello there")
	if databaseError.Error() != "hello there" {
		t.Errorf("expected hello there but got %s", databaseError.Error())
	}

	databaseError = NewDatabaseError("random message")
	if databaseError.Error() != "random message" {
		t.Errorf("expected database error but got %s", databaseError.Error())
	}
}

func TestIsDatabaseError(t *testing.T) {
	databaseError := NewDatabaseError("hello there")
	if !IsDatabaseError(databaseError) {
		t.Errorf("expected true but got false")
	}

	databaseError = NewDatabaseError("random message")
	if !IsDatabaseError(databaseError) {
		t.Errorf("expected true but got false")
	}
}

func TestIsDatabaseErrorNegative(t *testing.T) {
	authorizationError := NewAuthorizationError()
	if IsDatabaseError(authorizationError) {
		t.Errorf("expected false but got true")
	}
}
