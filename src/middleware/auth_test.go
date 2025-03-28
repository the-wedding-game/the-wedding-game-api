package middleware

import (
	"testing"
	test "the-wedding-game-api/_tests"
	apperrors "the-wedding-game-api/errors"
	"the-wedding-game-api/models"
	"the-wedding-game-api/types"
)

var (
	testAccessToken = models.AccessToken{Token: "test_token", UserID: 1, ExpiresOn: 1}
	testUser        = models.User{Username: "test_username", Role: types.Player}
	testUserAdmin   = models.User{Username: "test_username", Role: types.Admin}
)

func SetupMockDb() *models.MockDB {
	mockDB := &models.MockDB{}
	models.GetConnection = func() models.DatabaseInterface {
		return mockDB
	}
	return mockDB
}

func createTestAccessToken(accessToken models.AccessToken) {
	database := models.GetConnection()
	database.Create(&accessToken)
}

func createTestUser(user models.User) {
	database := models.GetConnection()
	database.Create(&user)
}

func TestGetCurrentUser(t *testing.T) {
	SetupMockDb()
	createTestAccessToken(testAccessToken)
	createTestUser(testUser)

	request := test.GenerateBasicRequest()
	request.Request.Header.Set("Authorization", "Bearer token")
	user, err := GetCurrentUser(request)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if user.Username != testUser.Username {
		t.Errorf("expected test_username but got %s", user.Username)
	}

	if user.Role != testUser.Role {
		t.Errorf("expected Player but got %s", user.Role)
	}
}

func TestGetCurrentUserNoAccessToken(t *testing.T) {
	request := test.GenerateBasicRequest()
	_, err := GetCurrentUser(request)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if !apperrors.IsAuthenticationError(err) {
		t.Errorf("expected authentication error but got %s", err.Error())
	}

	if err.Error() != "access token is not provided" {
		t.Errorf("expected access token is not provided but got %s", err.Error())
	}
}

func TestGetCurrentUserInvalidAccessTokenFormat(t *testing.T) {
	request := test.GenerateBasicRequest()
	request.Request.Header.Set("Authorization", "token")
	_, err := GetCurrentUser(request)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if !apperrors.IsAuthenticationError(err) {
		t.Errorf("expected authentication error but got %s", err.Error())
	}

	if err.Error() != "invalid access token format" {
		t.Errorf("expected invalid access token format but got %s", err.Error())
	}
}

func TestCheckIsAdmin(t *testing.T) {
	SetupMockDb()
	createTestAccessToken(testAccessToken)
	createTestUser(testUserAdmin)

	request := test.GenerateBasicRequest()
	request.Request.Header.Set("Authorization", "Bearer test_token")

	err := CheckIsAdmin(request)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
	}
}

func TestCheckIsAdminNotAdmin(t *testing.T) {
	SetupMockDb()
	createTestAccessToken(testAccessToken)
	createTestUser(testUser)

	request := test.GenerateBasicRequest()
	request.Request.Header.Set("Authorization", "Bearer test_token")

	err := CheckIsAdmin(request)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if !apperrors.IsAuthorizationError(err) {
		t.Errorf("expected authorization error but got %s", err.Error())
	}

	if err.Error() != "access denied" {
		t.Errorf("expected access denied but got %s", err.Error())
	}
}

func TestCheckIsAdminNoAccessToken(t *testing.T) {
	request := test.GenerateBasicRequest()
	err := CheckIsAdmin(request)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if !apperrors.IsAuthenticationError(err) {
		t.Errorf("expected authorization error but got %s", err.Error())
	}

	if err.Error() != "access token is not provided" {
		t.Errorf("expected access token is not provided but got %s", err.Error())
	}
}

func TestCheckIsAdminInvalidAccessTokenFormat(t *testing.T) {
	request := test.GenerateBasicRequest()
	request.Request.Header.Set("Authorization", "token")

	err := CheckIsAdmin(request)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if !apperrors.IsAuthenticationError(err) {
		t.Errorf("expected authentication error but got %s", err.Error())
	}

	if err.Error() != "invalid access token format" {
		t.Errorf("expected invalid access token format but got %s", err.Error())
	}
}
