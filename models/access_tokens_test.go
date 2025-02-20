package models

import (
	"errors"
	"testing"
	test "the-wedding-game-api/_test"
	"the-wedding-game-api/db"
	apperrors "the-wedding-game-api/errors"
	"time"
)

var (
	testAccessToken = AccessToken{Token: "test_token", UserID: 1, ExpiresOn: 1}
)

func createTestAccessToken(accessToken AccessToken) {
	database := db.GetConnection()
	database.Create(&accessToken)
}

func createTestUser(user User) {
	database := db.GetConnection()
	database.Create(&user)
}

func TestLinkAccessTokenToUser(t *testing.T) {
	test.SetupMockDb()
	accessToken, err := LinkAccessTokenToUser(1)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if len(accessToken.Token) != 36 {
		t.Errorf("expected 36 but got %d", len(accessToken.Token))
	}

	if accessToken.UserID != 1 {
		t.Errorf("expected 1 but got %d", accessToken.UserID)
	}

	in24Hours := time.Now().Add(24 * time.Hour).Unix()
	if in24Hours-accessToken.ExpiresOn > 1 {
		t.Errorf("Access token expiry time invalid")
	}
}

func TestLinkAccessTokenToUserNegative(t *testing.T) {
	mockDb := test.SetupMockDb()

	mockDb.Error = errors.New("test_error")

	_, err := LinkAccessTokenToUser(1)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if !apperrors.IsDatabaseError(err) {
		t.Errorf("expected database error but got %s", err.Error())
	}
	if err.Error() != "test_error" {
		t.Errorf("expected error but got %s", err.Error())
	}
}

func TestGetUserByAccessToken(t *testing.T) {
	test.SetupMockDb()
	createTestAccessToken(testAccessToken)
	createTestUser(testUser)

	user, err := GetUserByAccessToken("test_access_token")
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if user.ID != testUser.ID {
		t.Errorf("expected 1 but got %d", user.ID)
	}
}

func TestGetUserByAccessTokenNotFound(t *testing.T) {
	test.SetupMockDb()

	_, err := GetUserByAccessToken("test_access_token")
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if !apperrors.IsAccessTokenNotFoundError(err) {
		t.Errorf("expected access token not found error but got %s", err.Error())
	}
	if err.Error() != "access token not found" {
		t.Errorf("expected access token not found but got %s", err.Error())
	}
}

func TestGetUserByAccessTokenUserNotFound(t *testing.T) {
	test.SetupMockDb()
	createTestAccessToken(testAccessToken)

	_, err := GetUserByAccessToken("test_access_token")
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if !apperrors.IsNotFoundError(err) {
		t.Errorf("expected not found error but got %s", err.Error())
	}
	if err.Error() != "User with key 1 not found." {
		t.Errorf("expected User with key 1 not found. but got %s", err.Error())
	}
}

func TestGetUserByAccessTokenError(t *testing.T) {
	mockDb := test.SetupMockDb()
	createTestAccessToken(testAccessToken)
	createTestUser(testUser)

	mockDb.Error = errors.New("test_error")

	_, err := GetUserByAccessToken("test_access_token")
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if !apperrors.IsDatabaseError(err) {
		t.Errorf("expected database error but got %s", err.Error())
	}
	if err.Error() != "test_error" {
		t.Errorf("expected error but got %s", err.Error())
	}
}
