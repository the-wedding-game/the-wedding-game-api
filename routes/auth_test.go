package routes

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	integrationtests "the-wedding-game-api/_test/integration"
	"the-wedding-game-api/db"
	"the-wedding-game-api/models"
	"the-wedding-game-api/types"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	integrationtests.Setup()
	router = GetRouter()

	code := m.Run()

	integrationtests.TearDown()
	os.Exit(code)
}

func makeRequest(method string, path string, bodyObj interface{}) (int, string) {
	body, err := json.Marshal(bodyObj)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	return resp.Code, resp.Body.String()
}

func TestLogin(t *testing.T) {
	db.ResetConnection()

	user := &models.User{
		Username: "test_user_for_login",
		Role:     types.Player,
	}
	_, err := user.Save()
	if err != nil {
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

func TestLoginWithAdmin(t *testing.T) {
	db.ResetConnection()
	user := models.User{
		Username: "test_admin_for_login",
		Role:     types.Admin,
	}
	_, err := user.Save()
	if err != nil {
		t.Errorf("Error creating user")
		return
	}

	loginRequest := types.LoginRequest{
		Username: "test_admin_for_login",
		Password: "testpassword",
	}
	statusCode, body := makeRequest("POST", "/auth/login", loginRequest)

	if statusCode != http.StatusForbidden {
		t.Errorf("Expected status code 403, got %v", statusCode)
	}
	expectedBody := `{"message":"access denied","status":"error"}`
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
	db.ResetConnection()

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
	db.ResetConnection()

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
