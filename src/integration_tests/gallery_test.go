package integrationtests

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"the-wedding-game-api/types"
)

func TestGetGallery1(t *testing.T) {
	if err := resetDatabase(); err != nil {
		t.Errorf("Error resetting database: %v", err)
		return
	}

	challenge1, err1 := createChallenge()
	challenge2, err2 := createChallenge()
	challenge3, err3 := createChallenge()
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

	err1 = createSubmission(challenge1.ID, user1.ID, "https://example.com/image1.jpg")
	err2 = createSubmission(challenge2.ID, user2.ID, "https://example.com/image2.jpg")
	err3 = createSubmission(challenge3.ID, user3.ID, "https://example.com/image3.jpg")
	if err1 != nil || err2 != nil || err3 != nil {
		t.Errorf("Error creating submissions")
		return
	}

	statusCode, body := makeRequestWithToken("GET", "/gallery", nil, accessToken.Token)
	if statusCode != 200 {
		t.Errorf("Invalid status code: %v", statusCode)
	}

	var response types.GalleryResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	expectedResponse := types.GalleryResponse{
		Images: []types.GalleryItem{
			{Url: "https://example.com/image3.jpg", SubmittedBy: user3.Username},
			{Url: "https://example.com/image2.jpg", SubmittedBy: user2.Username},
			{Url: "https://example.com/image1.jpg", SubmittedBy: user1.Username},
		},
	}
	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("Expected response: %v, got: %v", expectedResponse, response)
	}
}

func TestGetGallery2(t *testing.T) {
	if err := resetDatabase(); err != nil {
		t.Errorf("Error resetting database: %v", err)
		return
	}

	challenge1, err1 := createChallenge()
	challenge2, err2 := createChallenge()
	challenge3, err3 := createChallenge()
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

	_ = createSubmission(challenge1.ID, user1.ID, "https://example.com/image1.jpg")
	_ = createSubmission(challenge1.ID, user2.ID, "https://example.com/image2.jpg")
	_ = createSubmission(challenge1.ID, user3.ID, "https://example.com/image3.jpg")
	_ = createSubmission(challenge2.ID, user1.ID, "https://example.com/image4.jpg")
	_ = createSubmission(challenge2.ID, user2.ID, "https://example.com/image5.jpg")
	_ = createSubmission(challenge2.ID, user3.ID, "https://example.com/image6.jpg")
	_ = createSubmission(challenge3.ID, user1.ID, "https://example.com/image7.jpg")
	_ = createSubmission(challenge3.ID, user2.ID, "https://example.com/image8.jpg")
	_ = createSubmission(challenge3.ID, user3.ID, "https://example.com/image9.jpg")

	statusCode, body := makeRequestWithToken("GET", "/gallery", nil, accessToken.Token)
	if statusCode != 200 {
		t.Errorf("Invalid status code: %v", statusCode)
	}

	var response types.GalleryResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	expectedResponse := types.GalleryResponse{
		Images: []types.GalleryItem{
			{Url: "https://example.com/image9.jpg", SubmittedBy: user3.Username},
			{Url: "https://example.com/image8.jpg", SubmittedBy: user2.Username},
			{Url: "https://example.com/image7.jpg", SubmittedBy: user1.Username},
			{Url: "https://example.com/image6.jpg", SubmittedBy: user3.Username},
			{Url: "https://example.com/image5.jpg", SubmittedBy: user2.Username},
			{Url: "https://example.com/image4.jpg", SubmittedBy: user1.Username},
			{Url: "https://example.com/image3.jpg", SubmittedBy: user3.Username},
			{Url: "https://example.com/image2.jpg", SubmittedBy: user2.Username},
			{Url: "https://example.com/image1.jpg", SubmittedBy: user1.Username},
		},
	}
	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("Expected response: %v, got: %v", expectedResponse, response)
	}
}

func TestGalleryWithInvalidUrls(t *testing.T) {
	if err := resetDatabase(); err != nil {
		t.Errorf("Error resetting database: %v", err)
		return
	}

	challenge1, err1 := createChallenge()
	challenge2, err2 := createChallenge()
	challenge3, err3 := createChallenge()
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

	_ = createSubmission(challenge1.ID, user1.ID, "https://example.com/image1.jpg")
	_ = createSubmission(challenge1.ID, user2.ID, "https://example.com/image2.jpg")
	_ = createSubmission(challenge1.ID, user3.ID, "https://example.com/image3.jpg")
	_ = createSubmission(challenge2.ID, user1.ID, "invalid_url")
	_ = createSubmission(challenge2.ID, user2.ID, "https://example.com/image5.jpg")
	_ = createSubmission(challenge2.ID, user3.ID, "https://example.com/image6.jpg")
	_ = createSubmission(challenge3.ID, user1.ID, "https://example.com/image7.jpg")
	_ = createSubmission(challenge3.ID, user2.ID, "https://example.com/image8.jpg")
	_ = createSubmission(challenge3.ID, user3.ID, "invalid_url")

	statusCode, body := makeRequestWithToken("GET", "/gallery", nil, accessToken.Token)
	if statusCode != 200 {
		t.Errorf("Invalid status code: %v", statusCode)
	}

	var response types.GalleryResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	expectedResponse := types.GalleryResponse{
		Images: []types.GalleryItem{
			{Url: "https://example.com/image8.jpg", SubmittedBy: user2.Username},
			{Url: "https://example.com/image7.jpg", SubmittedBy: user1.Username},
			{Url: "https://example.com/image6.jpg", SubmittedBy: user3.Username},
			{Url: "https://example.com/image5.jpg", SubmittedBy: user2.Username},
			{Url: "https://example.com/image3.jpg", SubmittedBy: user3.Username},
			{Url: "https://example.com/image2.jpg", SubmittedBy: user2.Username},
			{Url: "https://example.com/image1.jpg", SubmittedBy: user1.Username},
		},
	}

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("Expected response: %v, got: %v", expectedResponse, response)
	}
}

func TestGalleryWithAnswerQuestionChallenges(t *testing.T) {
	if err := resetDatabase(); err != nil {
		t.Errorf("Error resetting database: %v", err)
		return
	}

	challenge1, err1 := createChallenge()
	challenge2, err2 := createAnswerQuestionChallenge()
	challenge3, err3 := createChallenge()
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

	_ = createSubmission(challenge1.ID, user1.ID, "https://example.com/image1.jpg")
	_ = createSubmission(challenge1.ID, user2.ID, "https://example.com/image2.jpg")
	_ = createSubmission(challenge1.ID, user3.ID, "https://example.com/image3.jpg")
	_ = createSubmission(challenge2.ID, user1.ID, "https://example.com/image4.jpg")
	_ = createSubmission(challenge2.ID, user2.ID, "https://example.com/image5.jpg")
	_ = createSubmission(challenge2.ID, user3.ID, "https://example.com/image6.jpg")
	_ = createSubmission(challenge3.ID, user1.ID, "https://example.com/image7.jpg")
	_ = createSubmission(challenge3.ID, user2.ID, "https://example.com/image8.jpg")
	_ = createSubmission(challenge3.ID, user3.ID, "https://example.com/image9.jpg")

	statusCode, body := makeRequestWithToken("GET", "/gallery", nil, accessToken.Token)
	if statusCode != 200 {
		t.Errorf("Invalid status code: %v", statusCode)
	}

	var response types.GalleryResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	expectedResponse := types.GalleryResponse{
		Images: []types.GalleryItem{
			{Url: "https://example.com/image9.jpg", SubmittedBy: user3.Username},
			{Url: "https://example.com/image8.jpg", SubmittedBy: user2.Username},
			{Url: "https://example.com/image7.jpg", SubmittedBy: user1.Username},
			{Url: "https://example.com/image3.jpg", SubmittedBy: user3.Username},
			{Url: "https://example.com/image2.jpg", SubmittedBy: user2.Username},
			{Url: "https://example.com/image1.jpg", SubmittedBy: user1.Username},
		},
	}

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("Expected response: %v, got: %v", expectedResponse, response)
	}
}

func TestGalleryWithNoSubmissions(t *testing.T) {
	if err := resetDatabase(); err != nil {
		t.Errorf("Error resetting database: %v", err)
		return
	}

	_, err1 := createChallenge()
	_, err2 := createChallenge()
	_, err3 := createChallenge()
	if err1 != nil || err2 != nil || err3 != nil {
		t.Errorf("Error creating challenges")
		return
	}

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, body := makeRequestWithToken("GET", "/gallery", nil, accessToken.Token)
	if statusCode != 200 {
		t.Errorf("Invalid status code: %v", statusCode)
		return
	}

	var response types.GalleryResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if len(response.Images) != 0 {
		t.Errorf("Expected empty leaderboard, got: %v", response.Images)
	}
}

func TestGalleryWithNoAccessToken(t *testing.T) {
	statusCode, body := makeRequest("GET", "/gallery", nil)
	if statusCode != 401 {
		t.Errorf("Invalid status code: %v", statusCode)
	}

	expectedBody := "{\"message\":\"access token is not provided\",\"status\":\"error\"}"
	if body != expectedBody {
		t.Errorf("Expected body: %v, got: %v", expectedBody, body)
	}
}

func TestGalleryInvalidAccessToken(t *testing.T) {
	statusCode, body := makeRequestWithToken("GET", "/gallery", nil, "invalid_token")
	if statusCode != 403 {
		t.Errorf("Invalid status code: %v", statusCode)
	}

	expectedBody := "{\"message\":\"access denied\",\"status\":\"error\"}"
	if body != expectedBody {
		t.Errorf("Expected body: %v, got: %v", expectedBody, body)
	}
}
