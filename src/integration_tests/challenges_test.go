package integrationtests

import (
	"bytes"
	"encoding/json"
	"net/http"
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

	if response.Challenges[1].Completed != false {
		t.Errorf("Expected completed false, got %v", response.Challenges[1].Completed)
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
