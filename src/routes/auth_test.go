package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"the-wedding-game-api/models"
	"the-wedding-game-api/types"
)

func TestLogin(t *testing.T) {
	models.ResetConnection()

	user := &models.User{
		Username: "test_user_for_login",
		Role:     types.Player,
	}
	_, err := user.Save()
	if err != nil {
		t.Errorf("Error creating user")
		return
	}

	loginRequest := types.LoginRequest{
		Username: "test_user_for_login",
		Password: "testpassword",
	}
	statusCode, body := makeRequest("POST", "/auth/login", loginRequest)

	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	var loginResponse types.LoginResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&loginResponse); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if loginResponse.User.Username != "test_user_for_login" {
		t.Errorf("Expected username test_user_for_login, got %v", loginResponse.User.Username)
	}
	if loginResponse.User.Role != types.Player {
		t.Errorf("Expected role PLAYER, got %v", loginResponse.User.Role)
	}
	if loginResponse.AccessToken == "" {
		t.Errorf("Expected access token to not be empty")
	}
	if len(loginResponse.AccessToken) != 36 {
		t.Errorf("Expected access token to be a UUID")
	}
}

func TestLoginWithAdminValidPassword(t *testing.T) {
	models.ResetConnection()
	user := models.User{
		Username: "test_admin_for_login1",
		Role:     types.Admin,
	}
	_, err := user.Save()
	if err != nil {
		t.Errorf("Error creating user")
		return
	}

	err = os.Setenv("ADMIN_PASSWORD", "testpassword")
	if err != nil {
		t.Errorf("Error setting environment variable: %s", err.Error())
		return
	}
	loginRequest := types.LoginRequest{
		Username: "test_admin_for_login1",
		Password: "testpassword",
	}
	statusCode, body := makeRequest("POST", "/auth/login", loginRequest)

	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	var loginResponse types.LoginResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&loginResponse); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if loginResponse.User.Username != "test_admin_for_login1" {
		t.Errorf("Expected username test_admin_for_login1, got %v", loginResponse.User.Username)
	}
	if loginResponse.User.Role != types.Admin {
		t.Errorf("Expected role ADMIN, got %v", loginResponse.User.Role)
	}
	if loginResponse.AccessToken == "" {
		t.Errorf("Expected access token to not be empty")
	}
	if len(loginResponse.AccessToken) != 36 {
		t.Errorf("Expected access token to be a UUID")
	}

}

func TestLoginWithAdminInvalidPassword(t *testing.T) {
	models.ResetConnection()
	user := models.User{
		Username: "test_admin_for_login2",
		Role:     types.Admin,
	}
	_, err := user.Save()
	if err != nil {
		t.Errorf("Error creating user")
		return
	}

	err = os.Setenv("ADMIN_PASSWORD", "testpassword")
	if err != nil {
		t.Errorf("Error setting environment variable: %s", err.Error())
		return
	}

	loginRequest := types.LoginRequest{
		Username: "test_admin_for_login2",
		Password: "invalidpassword",
	}
	statusCode, body := makeRequest("POST", "/auth/login", loginRequest)

	if statusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code 401, got %v", statusCode)
	}
	expectedBody := `{"message":"invalid password","status":"error"}`
	if body != expectedBody {
		t.Errorf("Expected body %v, got %v", expectedBody, body)
	}
}

func TestLoginWithMissingBody(t *testing.T) {
	statusCode, body := makeRequest("POST", "/auth/login", nil)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}
	expectedBody := `{"message":"Key: 'LoginRequest.Username' Error:Field validation for 'Username' failed on the 'required' tag","status":"error"}`
	if body != expectedBody {
		t.Errorf("Expected body %v, got %v", expectedBody, body)
	}
}

func TestLoginWithEmptyBody(t *testing.T) {
	statusCode, body := makeRequest("POST", "/auth/login", map[string]string{})

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}
	expectedBody := `{"message":"Key: 'LoginRequest.Username' Error:Field validation for 'Username' failed on the 'required' tag","status":"error"}`
	if body != expectedBody {
		t.Errorf("Expected body %v, got %v", expectedBody, body)
	}
}

func TestLoginWithEmptyUsername(t *testing.T) {
	loginRequest := types.LoginRequest{
		Username: "",
		Password: "testpassword",
	}
	statusCode, body := makeRequest("POST", "/auth/login", loginRequest)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	expectedBody := `{"message":"Key: 'LoginRequest.Username' Error:Field validation for 'Username' failed on the 'required' tag","status":"error"}`
	if body != expectedBody {
		t.Errorf("Expected body %v, got %v", expectedBody, body)
	}
}

func TestLoginWithEmptyPassword(t *testing.T) {
	models.ResetConnection()

	user := &models.User{
		Username: "test_user_for_login_with_empty_password",
		Role:     types.Player,
	}
	_, err := user.Save()
	if err != nil {
		t.Errorf("Error creating user")
		return
	}

	loginRequest := types.LoginRequest{
		Username: "test_user_for_login_with_empty_password",
		Password: "",
	}
	statusCode, body := makeRequest("POST", "/auth/login", loginRequest)

	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	var loginResponse types.LoginResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&loginResponse); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if loginResponse.User.Username != "test_user_for_login_with_empty_password" {
		t.Errorf("Expected username test_user_for_login_with_empty_password, got %v", loginResponse.User.Username)
	}
	if loginResponse.User.Role != types.Player {
		t.Errorf("Expected role PLAYER, got %v", loginResponse.User.Role)
	}
}

func TestLoginWithNonexistentUser(t *testing.T) {
	models.ResetConnection()

	loginRequest := types.LoginRequest{
		Username: "nonexistent_user",
		Password: "testpassword",
	}
	statusCode, body := makeRequest("POST", "/auth/login", loginRequest)

	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	var loginResponse types.LoginResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&loginResponse); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if loginResponse.User.Username != "nonexistent_user" {
		t.Errorf("Expected username nonexistent_user, got %v", loginResponse.User.Username)
	}
	if loginResponse.User.Role != types.Player {
		t.Errorf("Expected role PLAYER, got %v", loginResponse.User.Role)
	}
	if loginResponse.AccessToken == "" {
		t.Errorf("Expected access token to not be empty")
	}
	if len(loginResponse.AccessToken) != 36 {
		t.Errorf("Expected access token to be a UUID")
	}
}

func TestGetCurrentUser(t *testing.T) {
	models.ResetConnection()

	user := models.User{
		Username: "test_user_for_get_current_user",
		Role:     types.Player,
	}
	user, err := user.Save()
	if err != nil {
		t.Errorf("Error creating user")
		return
	}

	accessToken, err := models.LinkAccessTokenToUser(user.ID)
	if err != nil {
		t.Errorf("Error creating access token")
		return
	}

	req, err := http.NewRequest("GET", "/auth/current-user", nil)
	if err != nil {
		t.Errorf("Error creating request")
		return
	}
	req.Header.Set("Authorization", "Bearer "+accessToken.Token)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", resp.Code)
	}

	var userResponse types.UserResponse
	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&userResponse); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if userResponse.Username != "test_user_for_get_current_user" {
		t.Errorf("Expected username test_user_for_get_current_user, got %v", userResponse.Username)
	}
	if userResponse.Role != types.Player {
		t.Errorf("Expected role PLAYER, got %v", userResponse.Role)
	}
}

func TestGetCurrentUserWithMissingAccessToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/auth/current-user", nil)
	if err != nil {
		t.Errorf("Error creating request")
		return
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code 401, got %v", resp.Code)
	}

	expectedBody := `{"message":"access token is not provided","status":"error"}`
	if resp.Body.String() != expectedBody {
		t.Errorf("Expected body %v, got %v", expectedBody, resp.Body.String())
	}
}

func TestGetCurrentUserWithNonExistentAccessToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/auth/current-user", nil)
	if err != nil {
		t.Errorf("Error creating request")
		return
	}
	req.Header.Set("Authorization", "Bearer invalid_token")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusForbidden {
		t.Errorf("Expected status code 403, got %v", resp.Code)
	}

	expectedBody := `{"message":"access denied","status":"error"}`
	if resp.Body.String() != expectedBody {
		t.Errorf("Expected body %v, got %v", expectedBody, resp.Body.String())
	}
}

func TestGetCurrentUserWithInvalidAccessToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/auth/current-user", nil)
	if err != nil {
		t.Errorf("Error creating request")
		return
	}
	req.Header.Set("Authorization", "invalid_token")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code 401, got %v", resp.Code)
	}

	expectedBody := `{"message":"invalid access token format","status":"error"}`
	if resp.Body.String() != expectedBody {
		t.Errorf("Expected body %v, got %v", expectedBody, resp.Body.String())
	}
}
