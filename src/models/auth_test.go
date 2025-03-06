package models

import (
	"errors"
	"os"
	"testing"
	test "the-wedding-game-api/_tests"
	"the-wedding-game-api/db"
	apperrors "the-wedding-game-api/errors"
	"the-wedding-game-api/types"
)

var (
	testUser      = User{Username: "test_username", Role: types.Player}
	testUserAdmin = User{Username: "test_admin_username", Role: types.Admin}
)

func TestNewUser(t *testing.T) {
	user := NewUser(testUser.Username)
	if user.Username != testUser.Username {
		t.Errorf("expected %s but got %s", testUser.Username, user.Username)
	}
	if user.Role != types.Player {
		t.Errorf("expected PLAYER but got %s", user.Role)
	}

	user = NewUser(testUserAdmin.Username)
	if user.Username != testUserAdmin.Username {
		t.Errorf("expected %s but got %s", testUserAdmin.Username, user.Username)
	}
	if user.Role != types.Player {
		t.Errorf("expected PLAYER but got %s", user.Role)
	}
}

func TestDoesUserExist(t *testing.T) {
	test.SetupMockDb()
	createTestUser(testUser)

	exists, user, err := DoesUserExist(testUser.Username)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if !exists {
		t.Errorf("expected true but got false")
	}
	if user.Username != testUser.Username {
		t.Errorf("expected %s but got %s", testUser.Username, user.Username)
	}
	if user.Role != types.Player {
		t.Errorf("expected PLAYER but got %s", user.Role)
	}
}

func TestDoesUserExistNotFound(t *testing.T) {
	test.SetupMockDb()

	exists, _, err := DoesUserExist(testUserAdmin.Username)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if exists {
		t.Errorf("expected false but got true")
	}
}

func TestDoesUserExistError(t *testing.T) {
	mockDb := test.SetupMockDb()
	createTestUser(testUser)

	mockDb.Error = errors.New("test_error")
	exists, _, err := DoesUserExist(testUserAdmin.Username)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if !apperrors.IsDatabaseError(err) {
		t.Errorf("expected true but got false")
	}
	if err.Error() != "test_error" {
		t.Errorf("expected test_error but got %s", err.Error())
	}
	if exists {
		t.Errorf("expected false but got true")
	}
}

func TestUserSave(t *testing.T) {
	test.SetupMockDb()

	user := NewUser(testUser.Username)
	savedUser, err := user.Save()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if savedUser.Username != testUser.Username {
		t.Errorf("expected %s but got %s", testUser.Username, savedUser.Username)
	}
	if savedUser.Role != types.Player {
		t.Errorf("expected PLAYER but got %s", savedUser.Role)
	}

	mockDb := db.GetConnection()
	var userFromDb User
	err = mockDb.First(&userFromDb, savedUser.ID).GetError()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}
	if userFromDb.Username != testUser.Username {
		t.Errorf("expected %s but got %s", testUser.Username, userFromDb.Username)
	}
	if userFromDb.Role != types.Player {
		t.Errorf("expected PLAYER but got %s", userFromDb.Role)
	}
}

func TestUserSaveError(t *testing.T) {
	mockDb := test.SetupMockDb()

	user := NewUser(testUser.Username)
	mockDb.Error = errors.New("test_error")
	_, err := user.Save()
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if !apperrors.IsDatabaseError(err) {
		t.Errorf("expected true but got false")
	}
	if err.Error() != "test_error" {
		t.Errorf("expected test_error but got %s", err.Error())
	}
}

func TestValidatePassword1(t *testing.T) {
	err := os.Setenv("ADMIN_PASSWORD", "test_password")
	if err != nil {
		t.Errorf("Failed to set environment variable: %s", err.Error())
	}

	err = ValidatePassword("test_password")
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}
}

func TestValidatePassword2(t *testing.T) {
	err := os.Setenv("ADMIN_PASSWORD", "another password")
	if err != nil {
		t.Errorf("Failed to set environment variable: %s", err.Error())
	}

	err = ValidatePassword("another password")
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}
}

func TestValidatePasswordError(t *testing.T) {
	err := os.Setenv("ADMIN_PASSWORD", "test_password")
	if err != nil {
		t.Errorf("Failed to set environment variable: %s", err.Error())
	}

	err = ValidatePassword("wrong_password")
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if !apperrors.IsAuthenticationError(err) {
		t.Errorf("expected true but got false")
	}
	if err.Error() != "invalid password" {
		t.Errorf("expected invalid password but got %s", err.Error())
	}
}

func TestGetPoints(t *testing.T) {
	test.SetupMockDb()
	createTestUser(testUser)

	user := NewUser(testUser.Username)
	points, err := user.GetPoints()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if points != 100 {
		t.Errorf("expected 100 but got %d", points)
	}
}

func TestGetPointsError(t *testing.T) {
	mockDb := test.SetupMockDb()
	createTestUser(testUser)

	mockDb.Error = errors.New("test_error")
	user := NewUser(testUser.Username)
	_, err := user.GetPoints()
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if !apperrors.IsDatabaseError(err) {
		t.Errorf("expected true but got false")
	}
	if err.Error() != "test_error" {
		t.Errorf("expected test_error but got %s", err.Error())
	}
}
