package routes

import (
	"bytes"
	"encoding/json"
	"reflect"
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

func TestGetCurrentUserPointsWithInactiveChallenges(t *testing.T) {
	if err := resetDatabase(); err != nil {
		t.Errorf("Error resetting database: %v", err)
		return
	}

	user, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge1, err1 := createChallengeWithPoints(100)
	challenge2, err2 := createInactiveChallengeWithPoints(200)
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
		Points: 400,
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

func TestGetLeaderboard1(t *testing.T) {
	if err := resetDatabase(); err != nil {
		t.Errorf("Error resetting database: %v", err)
		return
	}

	challenge1, err1 := createChallengeWithPoints(100)
	challenge2, err2 := createChallengeWithPoints(200)
	challenge3, err3 := createChallengeWithPoints(300)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Errorf("Error creating challenges")
		return
	}

	user1, _, err1 := createUserAndGetAccessToken()
	user2, _, err2 := createUserAndGetAccessToken()
	user3, accessToken, err3 := createUserAndGetAccessToken()
	if err1 != nil || err2 != nil || err3 != nil {
		t.Errorf("Error creating users")
		return
	}

	err1 = completeChallenge(challenge1.ID, user1.ID)
	err2 = completeChallenge(challenge2.ID, user2.ID)
	err3 = completeChallenge(challenge3.ID, user3.ID)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Errorf("Error completing challenges")
		return
	}

	statusCode, body := makeRequestWithToken("GET", "/leaderboard", nil, accessToken.Token)
	if statusCode != 200 {
		t.Errorf("Invalid status code: %v", statusCode)
	}

	var response types.GetLeaderboardResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	expectedResponse := types.GetLeaderboardResponse{
		Leaderboard: []types.LeaderboardEntry{
			{Username: user3.Username, Points: 300},
			{Username: user2.Username, Points: 200},
			{Username: user1.Username, Points: 100},
		},
	}
	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("Expected response: %v, got: %v", expectedResponse, response)
	}
}

func TestGetLeaderboard2(t *testing.T) {
	if err := resetDatabase(); err != nil {
		t.Errorf("Error resetting database: %v", err)
		return
	}

	challenge1, err1 := createChallengeWithPoints(100)
	challenge2, err2 := createChallengeWithPoints(200)
	challenge3, err3 := createChallengeWithPoints(300)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Errorf("Error creating challenges")
		return
	}

	user1, _, err1 := createUserAndGetAccessToken()
	user2, _, err2 := createUserAndGetAccessToken()
	user3, accessToken, err3 := createUserAndGetAccessToken()
	if err1 != nil || err2 != nil || err3 != nil {
		t.Errorf("Error creating users")
		return
	}

	err1 = completeChallenge(challenge1.ID, user1.ID)
	err2 = completeChallenge(challenge2.ID, user2.ID)
	err3 = completeChallenge(challenge3.ID, user3.ID)
	err4 := completeChallenge(challenge1.ID, user2.ID)
	err5 := completeChallenge(challenge3.ID, user2.ID)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		t.Errorf("Error completing challenges")
		return
	}

	statusCode, body := makeRequestWithToken("GET", "/leaderboard", nil, accessToken.Token)
	if statusCode != 200 {
		t.Errorf("Invalid status code: %v", statusCode)
	}

	var response types.GetLeaderboardResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	expectedResponse := types.GetLeaderboardResponse{
		Leaderboard: []types.LeaderboardEntry{
			{Username: user2.Username, Points: 600},
			{Username: user3.Username, Points: 300},
			{Username: user1.Username, Points: 100},
		},
	}
	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("Expected response: %v, got: %v", expectedResponse, response)
	}
}

func TestGetLeaderboardWithInactiveChallenges(t *testing.T) {
	if err := resetDatabase(); err != nil {
		t.Errorf("Error resetting database: %v", err)
		return
	}

	challenge1, err1 := createChallengeWithPoints(100)
	challenge2, err2 := createInactiveChallengeWithPoints(200)
	challenge3, err3 := createChallengeWithPoints(300)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Errorf("Error creating challenges")
		return
	}

	user1, _, err1 := createUserAndGetAccessToken()
	user2, _, err2 := createUserAndGetAccessToken()
	user3, accessToken, err3 := createUserAndGetAccessToken()
	if err1 != nil || err2 != nil || err3 != nil {
		t.Errorf("Error creating users")
		return
	}

	err1 = completeChallenge(challenge1.ID, user1.ID)
	err2 = completeChallenge(challenge2.ID, user2.ID)
	err3 = completeChallenge(challenge3.ID, user3.ID)
	err4 := completeChallenge(challenge1.ID, user2.ID)
	err5 := completeChallenge(challenge2.ID, user3.ID)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		t.Errorf("Error completing challenges")
		return
	}

	statusCode, body := makeRequestWithToken("GET", "/leaderboard", nil, accessToken.Token)
	if statusCode != 200 {
		t.Errorf("Invalid status code: %v", statusCode)
	}

	var response types.GetLeaderboardResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	expectedResponse := types.GetLeaderboardResponse{
		Leaderboard: []types.LeaderboardEntry{
			{Username: user3.Username, Points: 300},
			{Username: user1.Username, Points: 100},
			{Username: user2.Username, Points: 100},
		},
	}
	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("Expected response: %v, got: %v", expectedResponse, response)
	}
}

func TestGetLeaderboardNoChallenges(t *testing.T) {
	if err := resetDatabase(); err != nil {
		t.Errorf("Error resetting database: %v", err)
		return
	}

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, body := makeRequestWithToken("GET", "/leaderboard", nil, accessToken.Token)
	if statusCode != 200 {
		t.Errorf("Invalid status code: %v", statusCode)
	}

	var response types.GetLeaderboardResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if len(response.Leaderboard) != 0 {
		t.Errorf("Expected empty leaderboard, got: %v", response.Leaderboard)
	}
}

func TestGetLeaderboardNoSubmissions(t *testing.T) {
	if err := resetDatabase(); err != nil {
		t.Errorf("Error resetting database: %v", err)
		return
	}

	_, err1 := createChallengeWithPoints(100)
	_, err2 := createChallengeWithPoints(200)
	_, err3 := createChallengeWithPoints(300)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Errorf("Error creating challenges")
		return
	}

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, body := makeRequestWithToken("GET", "/leaderboard", nil, accessToken.Token)
	if statusCode != 200 {
		t.Errorf("Invalid status code: %v", statusCode)
	}

	var response types.GetLeaderboardResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if len(response.Leaderboard) != 0 {
		t.Errorf("Expected empty leaderboard, got: %v", response.Leaderboard)
	}
}

func TestGetLeaderboardNoAccessToken(t *testing.T) {
	statusCode, body := makeRequest("GET", "/leaderboard", nil)
	if statusCode != 401 {
		t.Errorf("Invalid status code: %v", statusCode)
	}

	expectedBody := "{\"message\":\"access token is not provided\",\"status\":\"error\"}"
	if body != expectedBody {
		t.Errorf("Expected body: %v, got: %v", expectedBody, body)
	}
}

func TestGetLeaderboardInvalidAccessToken(t *testing.T) {
	statusCode, body := makeRequestWithToken("GET", "/leaderboard", nil, "invalid_token")
	if statusCode != 403 {
		t.Errorf("Invalid status code: %v", statusCode)
	}

	expectedBody := "{\"message\":\"access denied\",\"status\":\"error\"}"
	if body != expectedBody {
		t.Errorf("Expected body: %v, got: %v", expectedBody, body)
	}
}
