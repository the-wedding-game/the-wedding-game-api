package integrationtests

import (
	"bytes"
	"encoding/json"
	"testing"
	"the-wedding-game-api/types"
)

func TestGetCurrentUserPoints(t *testing.T) {
	user, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge1, err1 := createChallengeWithPoints(100)
	challenge2, err2 := createChallengeWithPoints(200)
	challenge3, err3 := createChallengeWithPoints(300)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Errorf("Error creating challenges")
		return
	}

	err1 = completeChallenge(challenge1.ID, user.ID)
	err2 = completeChallenge(challenge2.ID, user.ID)
	err3 = completeChallenge(challenge3.ID, user.ID)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Errorf("Error completing challenges")
		return
	}

	statusCode, body := makeRequestWithToken("GET", "/points/me", nil, accessToken.Token)
	if statusCode != 200 {
		t.Errorf("Invalid status code: %v", statusCode)
	}

	var response types.CurrentUserPointsResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	expectedResponse := types.CurrentUserPointsResponse{
		Points: 600,
	}
	if response != expectedResponse {
		t.Errorf("Expected response: %v, got: %v", expectedResponse, response)
	}
}

func TestGetCurrentUserPointsNoPoints(t *testing.T) {
	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, body := makeRequestWithToken("GET", "/points/me", nil, accessToken.Token)
	if statusCode != 200 {
		t.Errorf("Invalid status code: %v", statusCode)
	}

	var response types.CurrentUserPointsResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	expectedResponse := types.CurrentUserPointsResponse{
		Points: 0,
	}
	if response != expectedResponse {
		t.Errorf("Expected response: %v, got: %v", expectedResponse, response)
	}
}

func TestGetCurrentUserPointsInvalidAccessToken(t *testing.T) {
	statusCode, body := makeRequestWithToken("GET", "/points/me", nil, "invalid_token")
	if statusCode != 403 {
		t.Errorf("Invalid status code: %v", statusCode)
	}

	expectedBody := "{\"message\":\"access denied\",\"status\":\"error\"}"
	if body != expectedBody {
		t.Errorf("Expected body: %v, got: %v", expectedBody, body)
	}
}

func TestGetCurrentUserPointsNoAccessToken(t *testing.T) {
	statusCode, body := makeRequest("GET", "/points/me", nil)
	if statusCode != 401 {
		t.Errorf("Invalid status code: %v", statusCode)
	}

	expectedBody := "{\"message\":\"access token is not provided\",\"status\":\"error\"}"
	if body != expectedBody {
		t.Errorf("Expected body: %v, got: %v", expectedBody, body)
	}
}
