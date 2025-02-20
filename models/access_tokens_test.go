package models

import (
	"errors"
	"fmt"
	"testing"
	test "the-wedding-game-api/_test"
	"the-wedding-game-api/db"
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
	fmt.Println(in24Hours)
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

	if err.Error() != "record not found: *models.AccessToken" {
		t.Errorf("expected record not found: *models.AccessToken but got %s", err.Error())
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

	if err.Error() != "record not found: *models.User" {
		t.Errorf("expected record not found: *models.User but got %s", err.Error())
	}
}
