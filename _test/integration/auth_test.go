package integration_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"the-wedding-game-api/db"
	"the-wedding-game-api/models"
	"the-wedding-game-api/routes"
	"the-wedding-game-api/types"
)

func TestMain(m *testing.M) {
	dockerComposeUp()
	setupTestDb()

	code := m.Run()

	dockerComposeDown()
	os.Exit(code)
}

func makeRequest(method string, path string, bodyObj interface{}) (int, string) {
	router := routes.GetRouter()

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
	user := &models.User{
		Username: "test_user_for_login",
		Role:     types.Player,
	}
	database := db.GetConnection().Create(&user)
	if database.GetError() != nil {
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
