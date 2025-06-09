package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
	"testing"
	"the-wedding-game-api/models"
	"the-wedding-game-api/types"
)

func TestGetChallengeByIdUploadPhoto(t *testing.T) {
	models.ResetConnection()

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err := challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challengeId := challenge.ID
	statusCode, body := makeRequestWithToken("GET", "/challenges/"+strconv.Itoa(int(challengeId)), nil, accessToken.Token)

	if statusCode != 200 {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	var response types.GetChallengeResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if response.Name != "test_challenge" {
		t.Errorf("Expected name test_challenge, got %v", response.Name)
	}

	if response.Description != "test_description" {
		t.Errorf("Expected description test_description, got %v", response.Description)
	}

	if response.Points != 10 {
		t.Errorf("Expected points 10, got %v", response.Points)
	}

	if response.Image != "test_image" {
		t.Errorf("Expected image test_image, got %v", response.Image)
	}

	if response.Type != types.UploadPhotoChallenge {
		t.Errorf("Expected type UPLOAD_PHOTO, got %v", response.Type)
	}

	if response.Status != types.ActiveChallenge {
		t.Errorf("Expected status ACTIVE, got %v", response.Status)
	}

	if response.Completed != false {
		t.Errorf("Expected completed false, got %v", response.Completed)
	}
}

func TestGetChallengeByIdAnswerQuestion(t *testing.T) {
	models.ResetConnection()

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err := challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challengeId := challenge.ID
	statusCode, body := makeRequestWithToken("GET", "/challenges/"+strconv.Itoa(int(challengeId)), nil, accessToken.Token)

	if statusCode != 200 {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	var response types.GetChallengeResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if response.Name != "test_challenge" {
		t.Errorf("Expected name test_challenge, got %v", response.Name)
	}

	if response.Description != "test_description" {
		t.Errorf("Expected description test_description, got %v", response.Description)
	}

	if response.Points != 10 {
		t.Errorf("Expected points 10, got %v", response.Points)
	}

	if response.Image != "test_image" {
		t.Errorf("Expected image test_image, got %v", response.Image)
	}

	if response.Type != types.AnswerQuestionChallenge {
		t.Errorf("Expected type ANSWER_QUESTION, got %v", response.Type)
	}

	if response.Status != types.ActiveChallenge {
		t.Errorf("Expected status ACTIVE, got %v", response.Status)
	}

	if response.Completed != false {
		t.Errorf("Expected completed false, got %v", response.Completed)
	}
}

func TestGetChallengeByIdCompleted(t *testing.T) {
	models.ResetConnection()

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err := challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	user, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challengeId := challenge.ID
	submission := models.NewSubmission(user.ID, challengeId, "test_answer")
	_, err = submission.Save()
	if err != nil {
		t.Errorf("Error saving submission: %v", err)
		return
	}

	statusCode, body := makeRequestWithToken("GET", "/challenges/"+strconv.Itoa(int(challengeId)), nil, accessToken.Token)

	if statusCode != 200 {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	var response types.GetChallengeResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if response.Name != "test_challenge" {
		t.Errorf("Expected name test_challenge, got %v", response.Name)
	}

	if response.Description != "test_description" {
		t.Errorf("Expected description test_description, got %v", response.Description)
	}

	if response.Points != 10 {
		t.Errorf("Expected points 10, got %v", response.Points)
	}

	if response.Image != "test_image" {
		t.Errorf("Expected image test_image, got %v", response.Image)
	}

	if response.Type != types.AnswerQuestionChallenge {
		t.Errorf("Expected type ANSWER_QUESTION, got %v", response.Type)
	}

	if response.Status != types.ActiveChallenge {
		t.Errorf("Expected status ACTIVE, got %v", response.Status)
	}

	if response.Completed != true {
		t.Errorf("Expected completed true, got %v", response.Completed)
	}
}

func TestGetChallengeByWithNoAccessToken(t *testing.T) {
	models.ResetConnection()

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err := challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	statusCode, responseBody := makeRequest("GET", "/challenges/"+strconv.Itoa(int(challenge.ID)), nil)

	if statusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code 401, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"access token is not provided\",\"status\":\"error\"}" {
		t.Errorf("Expected response Unauthorized, got %v", responseBody)
	}
}

func TestGetChallengeByIdInvalidAccessToken(t *testing.T) {
	models.ResetConnection()

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err := challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	statusCode, responseBody := makeRequestWithToken("GET", "/challenges/"+strconv.Itoa(int(challenge.ID)), nil, "invalid_access_token")

	if statusCode != http.StatusForbidden {
		t.Errorf("Expected status code 403, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"access denied\",\"status\":\"error\"}" {
		t.Errorf("Expected response Unauthorized, got %v", responseBody)
	}
}

func TestGetChallengeByIdNonExistentChallenge(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, responseBody := makeRequestWithToken("GET", "/challenges/999", nil, accessToken.Token)

	if statusCode != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Challenge with key 999 not found.\",\"status\":\"error\"}" {
		t.Errorf("Expected response Challenge not found, got %v", responseBody)
	}
}

func TestGetChallengeByIdInvalidChallengeId(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, responseBody := makeRequestWithToken("GET", "/challenges/invalid", nil, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"invalid challenge id\",\"status\":\"error\"}" {
		t.Errorf("Expected response Invalid challenge id, got %v", responseBody)
	}
}

func TestGetChallengeByIdWithInactiveChallenge(t *testing.T) {
	models.ResetConnection()

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.InactiveChallenge,
	}
	challenge, err := challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, responseBody := makeRequestWithToken("GET", "/challenges/"+strconv.Itoa(int(challenge.ID)), nil, accessToken.Token)

	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	var challengeResponse types.GetChallengeResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(responseBody)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&challengeResponse); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if challengeResponse.Name != "test_challenge" {
		t.Errorf("Expected name test_challenge, got %v", challengeResponse.Name)
	}

	if challengeResponse.Description != "test_description" {
		t.Errorf("Expected description test_description, got %v", challengeResponse.Description)
	}

	if challengeResponse.Points != 10 {
		t.Errorf("Expected points 10, got %v", challengeResponse.Points)
	}

	if challengeResponse.Image != "test_image" {
		t.Errorf("Expected image test_image, got %v", challengeResponse.Image)
	}

	if challengeResponse.Type != types.AnswerQuestionChallenge {
		t.Errorf("Expected type ANSWER_QUESTION, got %v", challengeResponse.Type)
	}

	if challengeResponse.Status != types.InactiveChallenge {
		t.Errorf("Expected status INACTIVE, got %v", challengeResponse.Status)
	}

	if challengeResponse.Completed != false {
		t.Errorf("Expected completed false, got %v", challengeResponse.Completed)
	}
}

func TestCreateChallengeUploadPhoto(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challengeRequest := types.CreateChallengeRequest{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "https://test_image.com",
		Type:        types.UploadPhotoChallenge,
	}
	statusCode, body := makeRequestWithToken("POST", "/challenges", challengeRequest, accessToken.Token)

	if statusCode != http.StatusCreated {
		t.Errorf("Expected status code 201, got %v", statusCode)
	}

	var response types.ChallengeCreatedResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if response.Name != "test_challenge" {
		t.Errorf("Expected name test_challenge, got %v", response.Name)
	}

	if response.Description != "test_description" {
		t.Errorf("Expected description test_description, got %v", response.Description)
	}

	if response.Points != 10 {
		t.Errorf("Expected points 10, got %v", response.Points)
	}

	if response.Image != "https://test_image.com" {
		t.Errorf("Expected image test_image, got %v", response.Image)
	}

	if response.Type != types.UploadPhotoChallenge {
		t.Errorf("Expected type UPLOAD_PHOTO, got %v", response.Type)
	}

	if response.Status != types.ActiveChallenge {
		t.Errorf("Expected status ACTIVE, got %v", response.Status)
	}
}

func TestCreateChallengeAnswerQuestion(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challengeRequest := types.CreateChallengeRequest{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "https://test_image.com",
		Type:        types.AnswerQuestionChallenge,
		Answer:      "test_answer",
	}
	statusCode, body := makeRequestWithToken("POST", "/challenges", challengeRequest, accessToken.Token)

	if statusCode != http.StatusCreated {
		t.Errorf("Expected status code 201, got %v", statusCode)
	}

	var response types.ChallengeCreatedResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if response.Name != "test_challenge" {
		t.Errorf("Expected name test_challenge, got %v", response.Name)
	}

	if response.Description != "test_description" {
		t.Errorf("Expected description test_description, got %v", response.Description)
	}

	if response.Points != 10 {
		t.Errorf("Expected points 10, got %v", response.Points)
	}

	if response.Image != "https://test_image.com" {
		t.Errorf("Expected image test_image, got %v", response.Image)
	}

	if response.Type != types.AnswerQuestionChallenge {
		t.Errorf("Expected type ANSWER_QUESTION, got %v", response.Type)
	}

	if response.Status != types.ActiveChallenge {
		t.Errorf("Expected status ACTIVE, got %v", response.Status)
	}
}

func TestCreateChallengeAsNotAdmin(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challengeRequest := types.CreateChallengeRequest{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "https://test_image.com",
		Type:        types.UploadPhotoChallenge,
	}
	statusCode, responseBody := makeRequestWithToken("POST", "/challenges", challengeRequest, accessToken.Token)

	if statusCode != http.StatusForbidden {
		t.Errorf("Expected status code 403, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"access denied\",\"status\":\"error\"}" {
		t.Errorf("Expected response Unauthorized, got %v", responseBody)
	}
}

func TestCreateChallengeWithInvalidAccessToken(t *testing.T) {
	models.ResetConnection()

	challengeRequest := types.CreateChallengeRequest{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "https://test_image.com",
		Type:        types.UploadPhotoChallenge,
	}
	statusCode, responseBody := makeRequestWithToken("POST", "/challenges", challengeRequest, "invalid_access_token")

	if statusCode != http.StatusForbidden {
		t.Errorf("Expected status code 403, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"access denied\",\"status\":\"error\"}" {
		t.Errorf("Expected response Unauthorized, got %v", responseBody)
	}
}

func TestCreateChallengeWithMissingAccessToken(t *testing.T) {
	models.ResetConnection()

	challengeRequest := types.CreateChallengeRequest{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "https://test_image.com",
		Type:        types.UploadPhotoChallenge,
	}
	statusCode, responseBody := makeRequest("POST", "/challenges", challengeRequest)

	if statusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code 401, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"access token is not provided\",\"status\":\"error\"}" {
		t.Errorf("Expected response Unauthorized, got %v", responseBody)
	}
}

func TestCreateChallengeWithMissingBody(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, responseBody := makeRequestWithToken("POST", "/challenges", nil, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'CreateChallengeRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag\\nKey: 'CreateChallengeRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag\\nKey: 'CreateChallengeRequest.Points' Error:Field validation for 'Points' failed on the 'required' tag\\nKey: 'CreateChallengeRequest.Image' Error:Field validation for 'Image' failed on the 'required' tag\\nKey: 'CreateChallengeRequest.Type' Error:Field validation for 'Type' failed on the 'required' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestCreateChallengeWithEmptyBody(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, responseBody := makeRequestWithToken("POST", "/challenges", map[string]string{}, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'CreateChallengeRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag\\nKey: 'CreateChallengeRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag\\nKey: 'CreateChallengeRequest.Points' Error:Field validation for 'Points' failed on the 'required' tag\\nKey: 'CreateChallengeRequest.Image' Error:Field validation for 'Image' failed on the 'required' tag\\nKey: 'CreateChallengeRequest.Type' Error:Field validation for 'Type' failed on the 'required' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestCreateChallengeWithEmptyName(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challengeRequest := types.CreateChallengeRequest{
		Name:        "",
		Description: "test_description",
		Points:      10,
		Image:       "https://test_image.com",
		Type:        types.UploadPhotoChallenge,
	}
	statusCode, responseBody := makeRequestWithToken("POST", "/challenges", challengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'CreateChallengeRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestCreateChallengeWithEmptyDescription(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challengeRequest := types.CreateChallengeRequest{
		Name:        "test_challenge",
		Description: "",
		Points:      10,
		Image:       "https://test_image.com",
		Type:        types.UploadPhotoChallenge,
	}
	statusCode, responseBody := makeRequestWithToken("POST", "/challenges", challengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'CreateChallengeRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestCreateChallengeWithEmptyPoints(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challengeRequest := map[string]interface{}{
		"name":        "test_challenge",
		"description": "test_description",
		"points":      "",
		"image":       "https://test_image.com",
		"type":        types.UploadPhotoChallenge,
	}
	statusCode, responseBody := makeRequestWithToken("POST", "/challenges", challengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"json: cannot unmarshal string into Go struct field CreateChallengeRequest.points of type uint\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestCreateChallengeWithZeroPoints(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challengeRequest := types.CreateChallengeRequest{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      0,
		Image:       "https://test_image.com",
		Type:        types.UploadPhotoChallenge,
	}
	statusCode, responseBody := makeRequestWithToken("POST", "/challenges", challengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'CreateChallengeRequest.Points' Error:Field validation for 'Points' failed on the 'required' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestCreateChallengeWithNegativePoints(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challengeRequest := map[string]any{
		"Name":        "test_challenge",
		"Description": "test_description",
		"Points":      -10,
		"Image":       "https://test_image.com",
		"Type":        string(types.UploadPhotoChallenge),
	}
	statusCode, responseBody := makeRequestWithToken("POST", "/challenges", challengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"json: cannot unmarshal number -10 into Go struct field CreateChallengeRequest.points of type uint\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestCreateChallengeWithEmptyImage(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challengeRequest := types.CreateChallengeRequest{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "",
		Type:        types.UploadPhotoChallenge,
	}
	statusCode, responseBody := makeRequestWithToken("POST", "/challenges", challengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'CreateChallengeRequest.Image' Error:Field validation for 'Image' failed on the 'required' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestCreateChallengeWithInvalidImage(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challengeRequest := types.CreateChallengeRequest{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "invalid_image",
		Type:        types.UploadPhotoChallenge,
	}
	statusCode, responseBody := makeRequestWithToken("POST", "/challenges", challengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'CreateChallengeRequest.Image' Error:Field validation for 'Image' failed on the 'url' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestCreateChallengeWithInvalidType(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challengeRequest := types.CreateChallengeRequest{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "https://test_image.com",
		Type:        "invalid_type",
	}
	statusCode, responseBody := makeRequestWithToken("POST", "/challenges", challengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"invalid challenge type\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestCreateChallengeWithEmptyType(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challengeRequest := types.CreateChallengeRequest{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "https://test_image.com",
		Type:        "",
	}
	statusCode, responseBody := makeRequestWithToken("POST", "/challenges", challengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'CreateChallengeRequest.Type' Error:Field validation for 'Type' failed on the 'required' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestCreateChallengeWithEmptyAnswerForAnswerQuestionChallenge(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challengeRequest := types.CreateChallengeRequest{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "https://test_image.com",
		Type:        types.AnswerQuestionChallenge,
		Answer:      "",
	}
	statusCode, responseBody := makeRequestWithToken("POST", "/challenges", challengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"answer is required for answer question challenges\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestCreateChallengeWithoutAnswerForAnswerQuestionChallenge(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challengeRequest := types.CreateChallengeRequest{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "https://test_image.com",
		Type:        types.AnswerQuestionChallenge,
	}
	statusCode, responseBody := makeRequestWithToken("POST", "/challenges", challengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"answer is required for answer question challenges\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestGetAllChallengesEmpty(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, body := makeRequestWithToken("GET", "/challenges", nil, accessToken.Token)

	if statusCode != 200 {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	var response types.GetChallengesResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if len(response.Challenges) != 0 {
		t.Errorf("Expected 0 challenges, got %v", len(response.Challenges))
	}
}

func TestGetAllChallenges(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	challenge1 := models.Challenge{
		Name:        "test_challenge_1",
		Description: "test_description_1",
		Points:      10,
		Image:       "test_image_1",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge1, err := challenge1.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	challenge2 := models.Challenge{
		Name:        "test_challenge_2",
		Description: "test_description_2",
		Points:      20,
		Image:       "test_image_2",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge2, err = challenge2.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, body := makeRequestWithToken("GET", "/challenges", nil, accessToken.Token)

	if statusCode != 200 {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	var response types.GetChallengesResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if len(response.Challenges) != 2 {
		t.Errorf("Expected 2 challenges, got %v", len(response.Challenges))
		return
	}

	expectedResponseChallenge1 := types.GetChallengeResponse{
		Id:          challenge1.ID,
		Name:        "test_challenge_1",
		Description: "test_description_1",
		Points:      10,
		Image:       "test_image_1",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
		Completed:   false,
	}

	if response.Challenges[0] != expectedResponseChallenge1 {
		t.Errorf("Expected challenge 1 to be %v, got %v", expectedResponseChallenge1, response.Challenges[0])
		return
	}

	expectedResponseChallenge2 := types.GetChallengeResponse{
		Id:          challenge2.ID,
		Name:        "test_challenge_2",
		Description: "test_description_2",
		Points:      20,
		Image:       "test_image_2",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
		Completed:   false,
	}

	if response.Challenges[1] != expectedResponseChallenge2 {
		t.Errorf("Expected challenge 2 to be %v, got %v", expectedResponseChallenge2, response.Challenges[1])
		return
	}
}

func TestGetAllChallengesWithInactiveChallenges(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	challenge1 := models.Challenge{
		Name:        "test_challenge_1",
		Description: "test_description_1",
		Points:      10,
		Image:       "test_image_1",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge1, err := challenge1.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	challenge2 := models.Challenge{
		Name:        "test_challenge_2",
		Description: "test_description_2",
		Points:      20,
		Image:       "test_image_2",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.InactiveChallenge,
	}
	challenge2, err = challenge2.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, body := makeRequestWithToken("GET", "/challenges", nil, accessToken.Token)

	if statusCode != 200 {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	var response types.GetChallengesResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if len(response.Challenges) != 1 {
		t.Errorf("Expected 1 challenge, got %v", len(response.Challenges))
	}

	if response.Challenges[0].Name != "test_challenge_1" {
		t.Errorf("Expected name test_challenge_1, got %v", response.Challenges[0].Name)
	}

	if response.Challenges[0].Description != "test_description_1" {
		t.Errorf("Expected description test_description_1, got %v", response.Challenges[0].Description)
	}

	if response.Challenges[0].Points != 10 {
		t.Errorf("Expected points 10, got %v", response.Challenges[0].Points)
	}

	if response.Challenges[0].Image != "test_image_1" {
		t.Errorf("Expected image test_image_1, got %v", response.Challenges[0].Image)
	}

	if response.Challenges[0].Type != types.UploadPhotoChallenge {
		t.Errorf("Expected type UPLOAD_PHOTO, got %v", response.Challenges[0].Type)
	}

	if response.Challenges[0].Status != types.ActiveChallenge {
		t.Errorf("Expected status ACTIVE, got %v", response.Challenges[0].Status)
	}

	if response.Challenges[0].Completed != false {
		t.Errorf("Expected completed false, got %v", response.Challenges[0].Completed)
	}
}

func TestGetAllChallengesWithCompleted(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	challenge1 := models.Challenge{
		Name:        "test_challenge_1",
		Description: "test_description_1",
		Points:      10,
		Image:       "test_image_1",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge1, err := challenge1.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	challenge2 := models.Challenge{
		Name:        "test_challenge_2",
		Description: "test_description_2",
		Points:      20,
		Image:       "test_image_2",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge2, err = challenge2.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	user, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	submission := models.NewSubmission(user.ID, challenge2.ID, "test_answer")
	_, err = submission.Save()
	if err != nil {
		t.Errorf("Error saving submission: %v", err)
		return
	}

	statusCode, body := makeRequestWithToken("GET", "/challenges", nil, accessToken.Token)

	if statusCode != 200 {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	var response types.GetChallengesResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if len(response.Challenges) != 2 {
		t.Errorf("Expected 2 challenges, got %v", len(response.Challenges))
	}

	if response.Challenges[0].Name != "test_challenge_1" {
		t.Errorf("Expected name test_challenge_1, got %v", response.Challenges[0].Name)
	}

	if response.Challenges[0].Description != "test_description_1" {
		t.Errorf("Expected description test_description_1, got %v", response.Challenges[0].Description)
	}

	if response.Challenges[0].Points != 10 {
		t.Errorf("Expected points 10, got %v", response.Challenges[0].Points)
	}

	if response.Challenges[0].Image != "test_image_1" {
		t.Errorf("Expected image test_image_1, got %v", response.Challenges[0].Image)
	}

	if response.Challenges[0].Type != types.UploadPhotoChallenge {
		t.Errorf("Expected type UPLOAD_PHOTO, got %v", response.Challenges[0].Type)
	}

	if response.Challenges[0].Status != types.ActiveChallenge {
		t.Errorf("Expected status ACTIVE, got %v", response.Challenges[0].Status)
	}

	if response.Challenges[0].Completed != false {
		t.Errorf("Expected completed false, got %v", response.Challenges[0].Completed)
	}

	if response.Challenges[1].Name != "test_challenge_2" {
		t.Errorf("Expected name test_challenge_2, got %v", response.Challenges[1].Name)
	}

	if response.Challenges[1].Description != "test_description_2" {
		t.Errorf("Expected description test_description_2, got %v", response.Challenges[1].Description)
	}

	if response.Challenges[1].Points != 20 {
		t.Errorf("Expected points 20, got %v", response.Challenges[1].Points)
	}

	if response.Challenges[1].Image != "test_image_2" {
		t.Errorf("Expected image test_image_2, got %v", response.Challenges[1].Image)
	}

	if response.Challenges[1].Type != types.AnswerQuestionChallenge {
		t.Errorf("Expected type ANSWER_QUESTION, got %v", response.Challenges[1].Type)
	}

	if response.Challenges[1].Status != types.ActiveChallenge {
		t.Errorf("Expected status ACTIVE, got %v", response.Challenges[1].Status)
	}

	if response.Challenges[1].Completed != true {
		t.Errorf("Expected completed true, got %v", response.Challenges[1].Completed)
	}
}

func TestGetAllChallengesWithoutAccessToken(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	statusCode, responseBody := makeRequest("GET", "/challenges", nil)

	if statusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code 401, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"access token is not provided\",\"status\":\"error\"}" {
		t.Errorf("Expected response Unauthorized, got %v", responseBody)
	}
}

func TestGetAllChallengesWithInvalidAccessToken(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	statusCode, responseBody := makeRequestWithToken("GET", "/challenges", nil, "invalid_access_token")

	if statusCode != http.StatusForbidden {
		t.Errorf("Expected status code 403, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"access denied\",\"status\":\"error\"}" {
		t.Errorf("Expected response Unauthorized, got %v", responseBody)
	}
}

func TestVerifyAnswerQuestionChallengeCorrectAnswer(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	answer := models.Answer{
		ChallengeID: challenge.ID,
		Value:       "test_answer",
	}
	answer, err = answer.Save()
	if err != nil {
		t.Errorf("Error saving answer: %v", err)
		return
	}

	isCompleted, err := models.IsChallengeCompleted(accessToken.UserID, challenge.ID)
	if err != nil {
		t.Errorf("Error checking if challenge is completed: %v", err)
		return
	}
	if isCompleted {
		t.Errorf("Expected challenge not completed, got %v", isCompleted)
		return
	}

	verifyAnswerRequest := types.VerifyAnswerRequest{
		Answer: "test_answer",
	}

	statusCode, responseBody := makeRequestWithToken("POST", "/challenges/"+strconv.Itoa(int(challenge.ID))+"/verify", verifyAnswerRequest, accessToken.Token)

	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	var response types.VerifyAnswerResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(responseBody)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if response.Correct != true {
		t.Errorf("Expected correct true, got %v", response.Correct)
	}

	isCompleted, err = models.IsChallengeCompleted(accessToken.UserID, challenge.ID)
	if err != nil {
		t.Errorf("Error checking if challenge is completed: %v", err)
		return
	}
	if !isCompleted {
		t.Errorf("Expected challenge completed, got %v", isCompleted)
		return
	}

}

func TestVerifyAnswerQuestionChallengeIncorrectAnswer(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	answer := models.Answer{
		ChallengeID: challenge.ID,
		Value:       "test_answer",
	}
	answer, err = answer.Save()
	if err != nil {
		t.Errorf("Error saving answer: %v", err)
		return
	}

	isCompleted, err := models.IsChallengeCompleted(accessToken.UserID, challenge.ID)
	if err != nil {
		t.Errorf("Error checking if challenge is completed: %v", err)
		return
	}
	if isCompleted {
		t.Errorf("Expected challenge not completed, got %v", isCompleted)
		return
	}

	verifyAnswerRequest := types.VerifyAnswerRequest{
		Answer: "incorrect_answer",
	}

	statusCode, responseBody := makeRequestWithToken("POST", "/challenges/"+strconv.Itoa(int(challenge.ID))+"/verify", verifyAnswerRequest, accessToken.Token)

	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	var response types.VerifyAnswerResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(responseBody)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if response.Correct != false {
		t.Errorf("Expected correct false, got %v", response.Correct)
	}

	isCompleted, err = models.IsChallengeCompleted(accessToken.UserID, challenge.ID)
	if err != nil {
		t.Errorf("Error checking if challenge is completed: %v", err)
		return
	}
	if isCompleted {
		t.Errorf("Expected challenge not completed, got %v", isCompleted)
		return
	}
}

func TestVerifyAnswerChallengeDoesNotExist(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	verifyAnswerRequest := types.VerifyAnswerRequest{
		Answer: "test_answer",
	}

	statusCode, responseBody := makeRequestWithToken("POST", "/challenges/1/verify", verifyAnswerRequest, accessToken.Token)

	if statusCode != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Challenge with key 1 not found.\",\"status\":\"error\"}" {
		t.Errorf("Expected response Not Found, got %v", responseBody)
	}
}

func TestVerifyAnswerChallengeUploadPhotoChallenge(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	verifyAnswerRequest := types.VerifyAnswerRequest{
		Answer: "https://images.com/test.jpg",
	}

	statusCode, responseBody := makeRequestWithToken("POST", "/challenges/"+strconv.Itoa(int(challenge.ID))+"/verify", verifyAnswerRequest, accessToken.Token)

	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	var response types.VerifyAnswerResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(responseBody)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if response.Correct != true {
		t.Errorf("Expected correct true, got %v", response.Correct)
	}

	isCompleted, err := models.IsChallengeCompleted(accessToken.UserID, challenge.ID)
	if err != nil {
		t.Errorf("Error checking if challenge is completed: %v", err)
		return
	}
	if !isCompleted {
		t.Errorf("Expected challenge completed, got %v", isCompleted)
		return
	}
}

func TestVerifyAnswerChallengeUploadPhotoChallengeInvalidUrl(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	verifyAnswerRequest := types.VerifyAnswerRequest{
		Answer: "invalid_url",
	}

	statusCode, responseBody := makeRequestWithToken("POST", "/challenges/"+strconv.Itoa(int(challenge.ID))+"/verify", verifyAnswerRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"invalid image url\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestVerifyAnswerChallengeAlreadyCompleted(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	answer := models.Answer{
		ChallengeID: challenge.ID,
		Value:       "test_answer",
	}
	answer, err = answer.Save()
	if err != nil {
		t.Errorf("Error saving answer: %v", err)
		return
	}

	submission := models.NewSubmission(accessToken.UserID, challenge.ID, "test_answer")
	_, err = submission.Save()
	if err != nil {
		t.Errorf("Error saving submission: %v", err)
		return
	}

	verifyAnswerRequest := types.VerifyAnswerRequest{
		Answer: "test_answer",
	}

	statusCode, responseBody := makeRequestWithToken("POST", "/challenges/"+strconv.Itoa(int(challenge.ID))+"/verify", verifyAnswerRequest, accessToken.Token)

	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	var response types.VerifyAnswerResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(responseBody)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if response.Correct != true {
		t.Errorf("Expected correct true, got %v", response.Correct)
	}
}

func TestGetAllChallengesAdmin(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	challenge1 := models.Challenge{
		Name:        "test_challenge_1",
		Description: "test_description_1",
		Points:      10,
		Image:       "test_image_1",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge1, err := challenge1.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	challenge2 := models.Challenge{
		Name:        "test_challenge_2",
		Description: "test_description_2",
		Points:      20,
		Image:       "test_image_2",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge2, err = challenge2.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, body := makeRequestWithToken("GET", "/admin/challenges", nil, accessToken.Token)

	if statusCode != 200 {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	var response types.GetChallengesAdminResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if len(response.Challenges) != 2 {
		t.Errorf("Expected 2 challenges, got %v", len(response.Challenges))
	}

	if response.Challenges[0].Name != "test_challenge_1" {
		t.Errorf("Expected name test_challenge_1, got %v", response.Challenges[0].Name)
	}

	if response.Challenges[0].Description != "test_description_1" {
		t.Errorf("Expected description test_description_1, got %v", response.Challenges[0].Description)
	}

	if response.Challenges[0].Points != 10 {
		t.Errorf("Expected points 10, got %v", response.Challenges[0].Points)
	}

	if response.Challenges[0].Image != "test_image_1" {
		t.Errorf("Expected image test_image_1, got %v", response.Challenges[0].Image)
	}

	if response.Challenges[0].Type != types.UploadPhotoChallenge {
		t.Errorf("Expected type UPLOAD_PHOTO, got %v", response.Challenges[0].Type)
	}

	if response.Challenges[0].Status != types.ActiveChallenge {
		t.Errorf("Expected status ACTIVE, got %v", response.Challenges[0].Status)
	}

	if response.Challenges[1].Name != "test_challenge_2" {
		t.Errorf("Expected name test_challenge_2, got %v", response.Challenges[1].Name)
	}

	if response.Challenges[1].Description != "test_description_2" {
		t.Errorf("Expected description test_description_2, got %v", response.Challenges[1].Description)
	}

	if response.Challenges[1].Points != 20 {
		t.Errorf("Expected points 20, got %v", response.Challenges[1].Points)
	}

	if response.Challenges[1].Image != "test_image_2" {
		t.Errorf("Expected image test_image_2, got %v", response.Challenges[1].Image)
	}

	if response.Challenges[1].Type != types.AnswerQuestionChallenge {
		t.Errorf("Expected type ANSWER_QUESTION, got %v", response.Challenges[1].Type)
	}

	if response.Challenges[1].Status != types.ActiveChallenge {
		t.Errorf("Expected status ACTIVE, got %v", response.Challenges[1].Status)
	}
}

func TestGetAllChallengesAdminNotAdmin(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, responseBody := makeRequestWithToken("GET", "/admin/challenges", nil, accessToken.Token)

	if statusCode != http.StatusForbidden {
		t.Errorf("Expected status code 403, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"access denied\",\"status\":\"error\"}" {
		t.Errorf("Expected response Forbidden, got %v", responseBody)
	}
}

func TestGetAllChallengesAdminWithoutAccessToken(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	statusCode, responseBody := makeRequest("GET", "/admin/challenges", nil)

	if statusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code 401, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"access token is not provided\",\"status\":\"error\"}" {
		t.Errorf("Expected response Unauthorized, got %v", responseBody)
	}
}

func TestGetAllChallengesAdminWithInvalidAccessToken(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	statusCode, responseBody := makeRequestWithToken("GET", "/admin/challenges", nil, "invalid_access_token")

	if statusCode != http.StatusForbidden {
		t.Errorf("Expected status code 403, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"access denied\",\"status\":\"error\"}" {
		t.Errorf("Expected response Unauthorized, got %v", responseBody)
	}
}

func TestUpdateChallenge(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	updateChallengeRequest := types.UpdateChallengeRequest{
		Name:        "updated_name",
		Description: "updated_description",
		Points:      20,
		Image:       "https://example.com/image.jpg",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.InactiveChallenge,
		Answer:      "test_answer",
	}

	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/"+strconv.Itoa(int(challenge.ID)), updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
		return
	}

	var response types.UpdateChallengeResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(responseBody)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if response.Name != "updated_name" {
		t.Errorf("Expected name updated_name, got %v", response.Name)
	}

	if response.Description != "updated_description" {
		t.Errorf("Expected description updated_description, got %v", response.Description)
	}

	if response.Points != 20 {
		t.Errorf("Expected points 20, got %v", response.Points)
	}

	if response.Image != "https://example.com/image.jpg" {
		t.Errorf("Expected image updated_image, got %v", response.Image)
	}

	if response.Type != types.AnswerQuestionChallenge {
		t.Errorf("Expected type ANSWER_QUESTION, got %v", response.Type)
	}

	if response.Status != types.InactiveChallenge {
		t.Errorf("Expected status INACTIVE, got %v", response.Status)
	}

	verify, err := models.VerifyAnswer(challenge.ID, "test_answer")
	if err != nil {
		t.Errorf("Error verifying answer: %v", err)
		return
	}
	if !verify {
		t.Errorf("Expected answer to be verified")
	}

}

func TestUpdateChallengeWithInvalidId(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	updateChallengeRequest := types.UpdateChallengeRequest{
		Name:        "updated_name",
		Description: "updated_description",
		Points:      20,
		Image:       "https://example.com/image.jpg",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.InactiveChallenge,
		Answer:      "test_answer",
	}

	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/sdds", updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"invalid challenge id\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestUpdateChallengeDoesNotExist(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	updateChallengeRequest := types.UpdateChallengeRequest{
		Name:        "updated_name",
		Description: "updated_description",
		Points:      20,
		Image:       "https://example.com/image.jpg",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.InactiveChallenge,
		Answer:      "test_answer",
	}

	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/99999", updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Challenge with key 99999 not found.\",\"status\":\"error\"}" {
		t.Errorf("Expected response Not Found, got %v", responseBody)
	}
}

func TestUpdateChallengeWithMissingName(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	updateChallengeRequest := map[string]interface{}{
		"description": "updated_description",
		"points":      20,
		"image":       "https://example.com/image.jpg",
		"type":        types.AnswerQuestionChallenge,
		"answer":      "test_answer",
		"status":      types.InactiveChallenge,
	}
	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/1", updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'UpdateChallengeRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestUpdateChallengeWithMissingDescription(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	updateChallengeRequest := map[string]interface{}{
		"name":   "updated_name",
		"points": 20,
		"image":  "https://example.com/image.jpg",
		"type":   types.AnswerQuestionChallenge,
		"answer": "test_answer",
		"status": types.InactiveChallenge,
	}
	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/1", updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'UpdateChallengeRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestUpdateChallengeWithMissingPoints(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	updateChallengeRequest := map[string]interface{}{
		"name":        "updated_name",
		"description": "updated_description",
		"image":       "https://example.com/image.jpg",
		"type":        types.AnswerQuestionChallenge,
		"answer":      "test_answer",
		"status":      types.InactiveChallenge,
	}
	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/1", updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'UpdateChallengeRequest.Points' Error:Field validation for 'Points' failed on the 'required' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestUpdateChallengeWithMissingImage(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	updateChallengeRequest := map[string]interface{}{
		"name":        "updated_name",
		"description": "updated_description",
		"points":      20,
		"type":        types.AnswerQuestionChallenge,
		"answer":      "test_answer",
		"status":      types.InactiveChallenge,
	}
	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/1", updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'UpdateChallengeRequest.Image' Error:Field validation for 'Image' failed on the 'required' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestUpdateChallengeWithMissingType(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	updateChallengeRequest := map[string]interface{}{
		"name":        "updated_name",
		"description": "updated_description",
		"points":      20,
		"image":       "https://example.com/image.jpg",
		"answer":      "test_answer",
		"status":      types.InactiveChallenge,
	}
	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/1", updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'UpdateChallengeRequest.Type' Error:Field validation for 'Type' failed on the 'required' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestUpdateChallengeWithMissingStatus(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	updateChallengeRequest := map[string]interface{}{
		"name":        "updated_name",
		"description": "updated_description",
		"points":      20,
		"image":       "https://example.com/image.jpg",
		"type":        types.AnswerQuestionChallenge,
		"answer":      "test_answer",
	}
	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/1", updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'UpdateChallengeRequest.Status' Error:Field validation for 'Status' failed on the 'required' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestUpdateChallengeWithInvalidType(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	updateChallengeRequest := map[string]interface{}{
		"name":        "updated_name",
		"description": "updated_description",
		"points":      20,
		"image":       "https://example.com/image.jpg",
		"type":        "invalid_type",
		"status":      types.InactiveChallenge,
		"answer":      "test_answer",
	}
	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/1", updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'UpdateChallengeRequest.Type' Error:Field validation for 'Type' failed on the 'oneof' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestUpdateChallengeWithInvalidStatus(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	updateChallengeRequest := map[string]interface{}{
		"name":        "updated_name",
		"description": "updated_description",
		"points":      20,
		"image":       "https://example.com/image.jpg",
		"type":        types.AnswerQuestionChallenge,
		"status":      "invalid_status",
		"answer":      "test_answer",
	}
	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/1", updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'UpdateChallengeRequest.Status' Error:Field validation for 'Status' failed on the 'oneof' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestUpdateChallengeWithMissingAnswerQuestionChallengeAnswer(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	updateChallengeRequest := map[string]interface{}{
		"name":        "updated_name",
		"description": "updated_description",
		"points":      20,
		"image":       "https://example.com/image.jpg",
		"type":        types.AnswerQuestionChallenge,
		"status":      types.InactiveChallenge,
	}
	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/"+strconv.Itoa(int(challenge.ID)), updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Answer cannot be empty when changing to AnswerQuestion challenge type\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}

}

func TestUpdateChallengeWithEmptyName(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	updateChallengeRequest := map[string]interface{}{
		"name":        "",
		"description": "updated_description",
		"points":      20,
		"image":       "https://example.com/image.jpg",
		"type":        types.AnswerQuestionChallenge,
		"status":      types.InactiveChallenge,
		"answer":      "test_answer",
	}
	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/1", updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'UpdateChallengeRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestUpdateChallengeWithEmptyDescription(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	updateChallengeRequest := map[string]interface{}{
		"name":        "updated_name",
		"description": "",
		"points":      20,
		"image":       "https://example.com/image.jpg",
		"type":        types.AnswerQuestionChallenge,
		"status":      types.InactiveChallenge,
		"answer":      "test_answer",
	}
	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/1", updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'UpdateChallengeRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestUpdateChallengeWithEmptyImage(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	updateChallengeRequest := map[string]interface{}{
		"name":        "updated_name",
		"description": "updated_description",
		"points":      20,
		"image":       "",
		"type":        types.AnswerQuestionChallenge,
		"status":      types.InactiveChallenge,
		"answer":      "test_answer",
	}
	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/1", updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'UpdateChallengeRequest.Image' Error:Field validation for 'Image' failed on the 'required' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestUpdateChallengeWithEmptyAnswerQuestionChallengeAnswer(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	updateChallengeRequest := map[string]interface{}{
		"name":        "updated_name",
		"description": "updated_description",
		"points":      20,
		"image":       "https://example.com/image.jpg",
		"type":        types.AnswerQuestionChallenge,
		"status":      types.InactiveChallenge,
		"answer":      "",
	}
	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/"+strconv.Itoa(int(challenge.ID)), updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Answer cannot be empty when changing to AnswerQuestion challenge type\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestUpdateChallengeWithInvalidImage(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	updateChallengeRequest := map[string]interface{}{
		"name":        "updated_name",
		"description": "updated_description",
		"points":      20,
		"image":       "invalid_image",
		"type":        types.AnswerQuestionChallenge,
		"status":      types.InactiveChallenge,
		"answer":      "test_answer",
	}
	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/1", updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Key: 'UpdateChallengeRequest.Image' Error:Field validation for 'Image' failed on the 'url' tag\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestUpdateChallengeHasSubmissions(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	submission := models.NewSubmission(accessToken.UserID, challenge.ID, "test_answer")
	_, err = submission.Save()
	if err != nil {
		t.Errorf("Error saving submission: %v", err)
		return
	}

	updateChallengeRequest := types.UpdateChallengeRequest{
		Name:        "updated_name",
		Description: "updated_description",
		Points:      20,
		Image:       "https://example.com/image.jpg",
		Type:        types.UploadPhotoChallenge,
		Status:      types.InactiveChallenge,
	}

	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/"+strconv.Itoa(int(challenge.ID)), updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	var response types.UpdateChallengeResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(responseBody)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	expectedResponse := types.UpdateChallengeResponse{
		Id:          challenge.ID,
		Name:        "updated_name",
		Description: "updated_description",
		Points:      20,
		Image:       "https://example.com/image.jpg",
		Type:        types.UploadPhotoChallenge,
		Status:      types.InactiveChallenge,
	}
	if response != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, response)
	}
}

func TestUpdateChallengeUpdateTypeButHasSubmissions1(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	submission := models.NewSubmission(accessToken.UserID, challenge.ID, "test_answer")
	_, err = submission.Save()
	if err != nil {
		t.Errorf("Error saving submission: %v", err)
		return
	}

	updateChallengeRequest := types.UpdateChallengeRequest{
		Name:        "updated_name",
		Description: "updated_description",
		Points:      20,
		Image:       "https://example.com/image.jpg",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.InactiveChallenge,
		Answer:      "test_answer",
	}

	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/"+strconv.Itoa(int(challenge.ID)), updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Cannot update challenge type if submissions exist\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestUpdateChallengeUpdateTypeButHasSubmissions2(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	submission := models.NewSubmission(accessToken.UserID, challenge.ID, "test_answer")
	_, err = submission.Save()
	if err != nil {
		t.Errorf("Error saving submission: %v", err)
		return
	}

	updateChallengeRequest := types.UpdateChallengeRequest{
		Name:        "updated_name",
		Description: "updated_description",
		Points:      20,
		Image:       "https://example.com/image.jpg",
		Type:        types.UploadPhotoChallenge,
		Status:      types.InactiveChallenge,
	}

	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/"+strconv.Itoa(int(challenge.ID)), updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	if responseBody != "{\"message\":\"Cannot update challenge type if submissions exist\",\"status\":\"error\"}" {
		t.Errorf("Expected response Bad Request, got %v", responseBody)
	}
}

func TestUpdateChallengeUpdateAnswerButSubmissionsExistAnswerDoesNotExist(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	submission := models.NewSubmission(accessToken.UserID, challenge.ID, "test_answer")
	_, err = submission.Save()
	if err != nil {
		t.Errorf("Error saving submission: %v", err)
		return
	}

	updateChallengeRequest := types.UpdateChallengeRequest{
		Name:        "updated_name",
		Description: "updated_description",
		Points:      20,
		Image:       "https://example.com/image.jpg",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.InactiveChallenge,
		Answer:      "updated_answer",
	}

	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/"+strconv.Itoa(int(challenge.ID)), updateChallengeRequest, accessToken.Token)

	if statusCode != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %v", statusCode)
	}

	expectedResponse := "{\"message\":\"error verifying answer: Answer with Challenge with key " + strconv.Itoa(int(challenge.ID)) + " not found.\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
	}
}

func TestUpdateChallengeUpdateAnswerButSubmissionsExistAnswerExists(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	submission := models.NewSubmission(accessToken.UserID, challenge.ID, "test_answer")
	_, err = submission.Save()
	if err != nil {
		t.Errorf("Error saving submission: %v", err)
		return
	}

	answer := models.Answer{
		ChallengeID: challenge.ID,
		Value:       "test_answer",
	}
	answer, err = answer.Save()
	if err != nil {
		t.Errorf("Error saving answer: %v", err)
		return
	}

	updateChallengeRequest := types.UpdateChallengeRequest{
		Name:        "updated_name",
		Description: "updated_description",
		Points:      20,
		Image:       "https://example.com/image.jpg",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.InactiveChallenge,
		Answer:      "updated_answer",
	}

	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/"+strconv.Itoa(int(challenge.ID)), updateChallengeRequest, accessToken.Token)
	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	expectedResponse := "{\"message\":\"Cannot update answer if submissions exist\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
	}
}

func TestUpdateChallengeUpdateAnswerButSubmissionsExistAnswerExistsAndIsTheSame(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	submission := models.NewSubmission(accessToken.UserID, challenge.ID, "test_answer")
	_, err = submission.Save()
	if err != nil {
		t.Errorf("Error saving submission: %v", err)
		return
	}

	answer := models.Answer{
		ChallengeID: challenge.ID,
		Value:       "test_answer",
	}
	answer, err = answer.Save()
	if err != nil {
		t.Errorf("Error saving answer: %v", err)
		return
	}

	updateChallengeRequest := types.UpdateChallengeRequest{
		Name:        "updated_name",
		Description: "updated_description",
		Points:      20,
		Image:       "https://example.com/image.jpg",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.InactiveChallenge,
		Answer:      "test_answer",
	}

	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/"+strconv.Itoa(int(challenge.ID)), updateChallengeRequest, accessToken.Token)
	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	expectedResponse := types.UpdateChallengeResponse{
		Id:          challenge.ID,
		Name:        "updated_name",
		Description: "updated_description",
		Points:      20,
		Image:       "https://example.com/image.jpg",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.InactiveChallenge,
	}
	var response types.UpdateChallengeResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(responseBody)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if response != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, response)
	}

	//Check if answer is saved
	verify, err := models.VerifyAnswer(challenge.ID, "test_answer")
	if err != nil {
		t.Errorf("Error verifying answer: %v", err)
		return
	}
	if !verify {
		t.Errorf("Expected answer to be verified")
	}
}

func TestUpdateChallengeUploadToAnswer(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	updateChallengeRequest := types.UpdateChallengeRequest{
		Name:        "updated_name",
		Description: "updated_description",
		Points:      20,
		Image:       "https://example.com/image.jpg",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.InactiveChallenge,
		Answer:      "test_answer",
	}

	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/"+strconv.Itoa(int(challenge.ID)), updateChallengeRequest, accessToken.Token)
	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	expectedResponse := types.UpdateChallengeResponse{
		Id:          challenge.ID,
		Name:        "updated_name",
		Description: "updated_description",
		Points:      20,
		Image:       "https://example.com/image.jpg",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.InactiveChallenge,
	}
	var response types.UpdateChallengeResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(responseBody)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if response != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, response)
	}

	//Check if answer is created
	verify, err := models.VerifyAnswer(challenge.ID, "test_answer")
	if err != nil {
		t.Errorf("Error verifying answer: %v", err)
		return
	}
	if !verify {
		t.Errorf("Expected answer to be verified")
	}
}

func TestUpdateChallengeAnswerToUpload(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	answer := models.Answer{
		ChallengeID: challenge.ID,
		Value:       "test_answer",
	}
	answer, err = answer.Save()
	if err != nil {
		t.Errorf("Error saving answer: %v", err)
		return
	}

	updateChallengeRequest := types.UpdateChallengeRequest{
		Name:        "updated_name",
		Description: "updated_description",
		Points:      20,
		Image:       "https://example.com/image.jpg",
		Type:        types.UploadPhotoChallenge,
		Status:      types.InactiveChallenge,
	}

	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/"+strconv.Itoa(int(challenge.ID)), updateChallengeRequest, accessToken.Token)
	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
	}

	expectedResponse := types.UpdateChallengeResponse{
		Id:          challenge.ID,
		Name:        "updated_name",
		Description: "updated_description",
		Points:      20,
		Image:       "https://example.com/image.jpg",
		Type:        types.UploadPhotoChallenge,
		Status:      types.InactiveChallenge,
	}
	var response types.UpdateChallengeResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(responseBody)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if response != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, response)
		return
	}
}

func TestUpdateChallengeUploadToAnswerNoAnswer(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	updateChallengeRequest := types.UpdateChallengeRequest{
		Name:        "updated_name",
		Description: "updated_description",
		Points:      20,
		Image:       "https://example.com/image.jpg",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.InactiveChallenge,
	}

	statusCode, responseBody := makeRequestWithToken("PUT", "/challenges/"+strconv.Itoa(int(challenge.ID)), updateChallengeRequest, accessToken.Token)
	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	expectedResponse := "{\"message\":\"Answer cannot be empty when changing to AnswerQuestion challenge type\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestGetSubmissions(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	user1, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	user2, _, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	submission1 := models.NewSubmission(accessToken.UserID, challenge.ID, "test_answer")
	_, err = submission1.Save()
	if err != nil {
		t.Errorf("Error saving submission: %v", err)
		return
	}

	submission2 := models.NewSubmission(user2.ID, challenge.ID, "test_answer2")
	_, err = submission2.Save()
	if err != nil {
		t.Errorf("Error saving submission: %v", err)
		return
	}

	statusCode, responseBody := makeRequestWithToken("GET", "/challenges/"+strconv.Itoa(int(challenge.ID))+"/submissions", nil, accessToken.Token)
	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
		return
	}

	var response types.GetSubmissionsResponse
	err = json.Unmarshal([]byte(responseBody), &response)
	if err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	var expectedSubmission1 = types.SubmissionForChallenge{
		Id:            submission1.ID,
		Answer:        submission1.Answer,
		ChallengeId:   challenge.ID,
		ChallengeName: challenge.Name,
		UserId:        user1.ID,
		Username:      user1.Username,
	}

	var expectedSubmission2 = types.SubmissionForChallenge{
		Id:            submission2.ID,
		Answer:        submission2.Answer,
		ChallengeId:   challenge.ID,
		ChallengeName: challenge.Name,
		UserId:        user2.ID,
		Username:      user2.Username,
	}

	var expectedResponse = types.GetSubmissionsResponse{
		Submissions: []types.SubmissionForChallenge{expectedSubmission1, expectedSubmission2},
	}

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("Expected response %v, got %v", expectedResponse, response)
		return
	}
}

func TestGetSubmissionsWithInvalid(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, responseBody := makeRequestWithToken("GET", "/challenges/invalid/submissions", nil, accessToken.Token)
	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
		return
	}

	expectedResponse := "{\"message\":\"invalid challenge id\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestGetSubmissionsWithChallengeDoesNotExist(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, responseBody := makeRequestWithToken("GET", "/challenges/999999/submissions", nil, accessToken.Token)
	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
		return
	}

	//expect empty
	var expectedResponse = types.GetSubmissionsResponse{
		Submissions: []types.SubmissionForChallenge{},
	}
	var response types.GetSubmissionsResponse
	err = json.Unmarshal([]byte(responseBody), &response)
	if err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestGetSubmissionsNoSubmissions(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	statusCode, responseBody := makeRequestWithToken("GET", "/challenges/"+strconv.Itoa(int(challenge.ID))+"/submissions", nil, accessToken.Token)
	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
		return
	}

	var expectedResponse = types.GetSubmissionsResponse{
		Submissions: []types.SubmissionForChallenge{},
	}
	var response types.GetSubmissionsResponse
	err = json.Unmarshal([]byte(responseBody), &response)
	if err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestGetSubmissionsDatabaseError(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	dropSubmissionsTable()
	defer setupTestDb()

	statusCode, responseBody := makeRequestWithToken("GET", "/challenges/"+strconv.Itoa(int(challenge.ID))+"/submissions", nil, accessToken.Token)
	if statusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code 500, got %v", statusCode)
		return
	}

	expectedResponse := "{\"message\":\"An unexpected error occurred.\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestGetSubmissionsWithInvalidToken(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err := challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	statusCode, responseBody := makeRequestWithToken("GET", "/challenges/"+strconv.Itoa(int(challenge.ID))+"/submissions", nil, "invalid_token")
	if statusCode != http.StatusForbidden {
		t.Errorf("Expected status code 401, got %v", statusCode)
		return
	}

	expectedResponse := "{\"message\":\"access denied\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestGetSubmissionsWithoutToken(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err := challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	statusCode, responseBody := makeRequest("GET", "/challenges/"+strconv.Itoa(int(challenge.ID))+"/submissions", nil)
	if statusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code 401, got %v", statusCode)
		return
	}

	expectedResponse := "{\"message\":\"access token is not provided\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestGetAnswer(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	answer := models.Answer{
		ChallengeID: challenge.ID,
		Value:       "test_answer_23487",
	}
	answer, err = answer.Save()
	if err != nil {
		t.Errorf("Error saving answer: %v", err)
		return
	}

	statusCode, responseBody := makeRequestWithToken("GET", "/challenges/"+strconv.Itoa(int(challenge.ID))+"/answer", nil, accessToken.Token)
	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
		return
	}

	var expectedResponse = types.GetAnswerResponse{
		Answer: answer.Value,
	}
	var response types.GetAnswerResponse
	err = json.Unmarshal([]byte(responseBody), &response)
	if err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestGetAnswerWithoutAccessToken(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err := challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	statusCode, responseBody := makeRequest("GET", "/challenges/"+strconv.Itoa(int(challenge.ID))+"/answer", nil)
	if statusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code 401, got %v", statusCode)
		return
	}

	expectedResponse := "{\"message\":\"access token is not provided\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestGetAnswerWithInvalidAccessToken(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err := challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	statusCode, responseBody := makeRequestWithToken("GET", "/challenges/"+strconv.Itoa(int(challenge.ID))+"/answer", nil, "invalid_token")
	if statusCode != http.StatusForbidden {
		t.Errorf("Expected status code 403, got %v", statusCode)
		return
	}

	expectedResponse := "{\"message\":\"access denied\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestGetAnswerWithNonAdminAccessToken(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	answer := models.Answer{
		ChallengeID: challenge.ID,
		Value:       "test_answer_23487",
	}
	answer, err = answer.Save()
	if err != nil {
		t.Errorf("Error saving answer: %v", err)
		return
	}

	statusCode, responseBody := makeRequestWithToken("GET", "/challenges/"+strconv.Itoa(int(challenge.ID))+"/answer", nil, accessToken.Token)
	if statusCode != http.StatusForbidden {
		t.Errorf("Expected status code 200, got %v", statusCode)
		return
	}

	expectedResponse := "{\"message\":\"access denied\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestGetAnswerWithChallengeDoesNotExist(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, responseBody := makeRequestWithToken("GET", "/challenges/999999/answer", nil, accessToken.Token)
	if statusCode != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %v", statusCode)
		return
	}

	expectedResponse := "{\"message\":\"Answer with Challenge with key 999999 not found.\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestGetAnswerWithChallengeIsNotAnswerQuestion(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	statusCode, responseBody := makeRequestWithToken("GET", "/challenges/"+strconv.Itoa(int(challenge.ID))+"/answer", nil, accessToken.Token)
	if statusCode != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %v", statusCode)
		return
	}

	expectedResponse := "{\"message\":\"Answer with Challenge with key 7 not found.\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestGetAnswerAnswerDoesNotExist(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	statusCode, responseBody := makeRequestWithToken("GET", "/challenges/"+strconv.Itoa(int(challenge.ID))+"/answer", nil, accessToken.Token)
	if statusCode != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %v", statusCode)
		return
	}

	expectedResponse := "{\"message\":\"Answer with Challenge with key " + strconv.Itoa(int(challenge.ID)) + " not found.\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestGetAnswerWithInvalidChallengeId(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, responseBody := makeRequestWithToken("GET", "/challenges/invalid/answer", nil, accessToken.Token)
	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
		return
	}

	expectedResponse := "{\"message\":\"invalid challenge id\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestGetAnswerWithDatabaseError(t *testing.T) {
	models.ResetConnection()
	deleteAllChallenges()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	dropAnswersTable()
	defer setupTestDb()

	statusCode, responseBody := makeRequestWithToken("GET", "/challenges/"+strconv.Itoa(int(challenge.ID))+"/answer", nil, accessToken.Token)
	if statusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code 500, got %v", statusCode)
		return
	}

	expectedResponse := "{\"message\":\"An unexpected error occurred.\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestDeleteChallenge(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	statusCode, responseBody := makeRequestWithToken("DELETE", "/challenges/"+strconv.Itoa(int(challenge.ID)), nil, accessToken.Token)
	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
		return
	}

	var response types.DeleteChallengeResponse
	err = json.Unmarshal([]byte(responseBody), &response)
	decoder := json.NewDecoder(bytes.NewReader([]byte(responseBody)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	expectedResponse := types.DeleteChallengeResponse{
		Id: challenge.ID,
	}

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("Expected response: %v, got: %v", expectedResponse, response)
	}

	_, err = models.GetChallengeByID(challenge.ID)
	if err == nil {
		t.Errorf("Expected challenge to be deleted, but it still exists")
		return
	}

	if err.Error() != "Challenge with key 1 not found." {
		t.Errorf("Expected Challenge with key 1 not found., got %v", err)
		return
	}
}

func TestDeleteChallengeWithAnswer(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	answer := models.Answer{
		ChallengeID: challenge.ID,
		Value:       "test_answer",
	}
	answer, err = answer.Save()
	if err != nil {
		t.Errorf("Error saving answer: %v", err)
		return
	}

	statusCode, responseBody := makeRequestWithToken("DELETE", "/challenges/"+strconv.Itoa(int(challenge.ID)), nil, accessToken.Token)
	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
		return
	}

	var response types.DeleteChallengeResponse
	err = json.Unmarshal([]byte(responseBody), &response)
	if err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	expectedResponse := types.DeleteChallengeResponse{
		Id: challenge.ID,
	}

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("Expected response: %v, got: %v", expectedResponse, response)
		return
	}

	answerExists, err := checkIfAnswerExists(answer.ID)
	if err != nil {
		t.Errorf("Error checking if answer exists: %v", err)
		return
	}

	if answerExists {
		t.Errorf("Expected answer to be deleted, but it still exists")
		return
	}

	answerExists, err = checkIfAnswerExistsForChallenge(challenge.ID)
	if err != nil {
		t.Errorf("Error checking if answer exists for challenge: %v", err)
		return
	}

	if answerExists {
		t.Errorf("Expected answer to be deleted for challenge, but it still exists")
		return
	}
}

func TestDeleteChallengeWithSubmissions(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	submission := models.NewSubmission(accessToken.UserID, challenge.ID, "test_answer")
	_, err = submission.Save()
	if err != nil {
		t.Errorf("Error saving submission: %v", err)
		return
	}

	statusCode, responseBody := makeRequestWithToken("DELETE", "/challenges/"+strconv.Itoa(int(challenge.ID)), nil, accessToken.Token)
	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
		return
	}

	var response types.DeleteChallengeResponse
	err = json.Unmarshal([]byte(responseBody), &response)
	if err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	expectedResponse := types.DeleteChallengeResponse{
		Id: challenge.ID,
	}

	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("Expected response: %v, got: %v", expectedResponse, response)
		return
	}

	submissionExists, err := checkIfSubmissionExists(submission.ID)
	if err != nil {
		t.Errorf("Error checking if submission exists: %v", err)
		return
	}

	if submissionExists {
		t.Errorf("Expected submission to be deleted, but it still exists")
		return
	}

	submissionExists, err = checkIfSubmissionExistsForChallenge(challenge.ID)
	if err != nil {
		t.Errorf("Error checking if answer exists for challenge: %v", err)
		return
	}

	if submissionExists {
		t.Errorf("Expected answer to be deleted for challenge, but it still exists")
		return
	}
}

func TestDeleteChallengeDoesNotExist(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, responseBody := makeRequestWithToken("DELETE", "/challenges/999999", nil, accessToken.Token)
	if statusCode != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %v", statusCode)
		return
	}

	expectedResponse := "{\"message\":\"Challenge with key 999999 not found.\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestDeleteChallengeInvalidId(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, responseBody := makeRequestWithToken("DELETE", "/challenges/invalid", nil, accessToken.Token)
	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
		return
	}

	expectedResponse := "{\"message\":\"invalid challenge id\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestDeleteChallengeWithEmptyToken(t *testing.T) {
	models.ResetConnection()

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err := challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	statusCode, responseBody := makeRequest("DELETE", "/challenges/"+strconv.Itoa(int(challenge.ID)), nil)
	if statusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code 401, got %v", statusCode)
		return
	}

	expectedResponse := "{\"message\":\"access token is not provided\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestDeleteChallengeInvalidToken(t *testing.T) {
	models.ResetConnection()

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err := challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	statusCode, responseBody := makeRequestWithToken("DELETE", "/challenges/"+strconv.Itoa(int(challenge.ID)), nil, "invalid_token")
	if statusCode != http.StatusForbidden {
		t.Errorf("Expected status code 403, got %v", statusCode)
		return
	}

	expectedResponse := "{\"message\":\"access denied\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}

func TestDeleteChallengeNotAdmin(t *testing.T) {
	models.ResetConnection()

	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	challenge, err = challenge.Save()
	if err != nil {
		t.Errorf("Error saving challenge: %v", err)
		return
	}

	statusCode, responseBody := makeRequestWithToken("DELETE", "/challenges/"+strconv.Itoa(int(challenge.ID)), nil, accessToken.Token)
	if statusCode != http.StatusForbidden {
		t.Errorf("Expected status code 403, got %v", statusCode)
		return
	}

	expectedResponse := "{\"message\":\"access denied\",\"status\":\"error\"}"
	if responseBody != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, responseBody)
		return
	}
}
