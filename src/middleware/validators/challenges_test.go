package validators

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"os"
	"testing"
	apperrors "the-wedding-game-api/errors"
	"the-wedding-game-api/types"
)

var (
	validCreateChallengeRequestUpload types.CreateChallengeRequest
	validCreateChallengeRequestAnswer types.CreateChallengeRequest
	validVerifyAnswerRequest          types.VerifyAnswerRequest
)

func TestMain(m *testing.M) {
	validCreateChallengeRequestUpload = types.CreateChallengeRequest{
		Name:        "testUpload",
		Description: "testUploadDescription",
		Points:      10,
		Image:       "https://fake.com/upload.jpg",
		Type:        "UPLOAD_PHOTO",
	}

	validCreateChallengeRequestAnswer = types.CreateChallengeRequest{
		Name:        "testAnswer",
		Description: "testAnswerDescription",
		Points:      10,
		Image:       "https://fake.com/answer.jpg",
		Type:        "ANSWER_QUESTION",
		Answer:      "answer",
	}

	validVerifyAnswerRequest = types.VerifyAnswerRequest{
		Answer: "test_answer",
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}

func generateRequestWithBodyOnly(requestData map[string]interface{}) *gin.Context {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	jsonData, _ := json.Marshal(requestData)
	body := bytes.NewBuffer(jsonData)
	c.Request = httptest.NewRequest("POST", "/challenges", body)
	return c
}

func generateRequestWithParamsOnly(params map[string]string) *gin.Context {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/challenges", nil)
	for key, value := range params {
		c.Params = append(c.Params, gin.Param{Key: key, Value: value})
	}
	return c
}

func generateRequestWithBodyAndParams(requestData map[string]interface{}, params map[string]string) *gin.Context {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	jsonData, _ := json.Marshal(requestData)
	body := bytes.NewBuffer(jsonData)
	c.Request = httptest.NewRequest("POST", "/challenges", body)

	for key, value := range params {
		c.Params = append(c.Params, gin.Param{Key: key, Value: value})
	}
	return c
}

func TestValidateCreateChallengeRequestValidUploadPhoto(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestUpload.Name,
		"description": validCreateChallengeRequestUpload.Description,
		"points":      validCreateChallengeRequestUpload.Points,
		"Image":       validCreateChallengeRequestUpload.Image,
		"type":        validCreateChallengeRequestUpload.Type,
	}
	c := generateRequestWithBodyOnly(requestData)

	createChallengeRequest, err := ValidateCreateChallengeRequest(c)
	if err != nil {
		t.Error("Expected no error, got", err)
	}

	if createChallengeRequest.Name != validCreateChallengeRequestUpload.Name {
		t.Error("Expected name to be", validCreateChallengeRequestUpload.Name, "got", createChallengeRequest.Name)
	}
	if createChallengeRequest.Description != validCreateChallengeRequestUpload.Description {
		t.Error("Expected description to be", validCreateChallengeRequestUpload.Description, "got", createChallengeRequest.Description)
	}
	if createChallengeRequest.Points != validCreateChallengeRequestUpload.Points {
		t.Error("Expected points to be", validCreateChallengeRequestUpload.Points, "got", createChallengeRequest.Points)
	}
	if createChallengeRequest.Image != validCreateChallengeRequestUpload.Image {
		t.Error("Expected image to be", validCreateChallengeRequestUpload.Image, "got", createChallengeRequest.Image)
	}
	if createChallengeRequest.Type != validCreateChallengeRequestUpload.Type {
		t.Error("Expected type to be", validCreateChallengeRequestUpload.Type, "got", createChallengeRequest.Type)
	}
}

func TestValidateCreateChallengeRequestValidAnswerQuestion(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestAnswer.Name,
		"description": validCreateChallengeRequestAnswer.Description,
		"points":      validCreateChallengeRequestAnswer.Points,
		"Image":       validCreateChallengeRequestAnswer.Image,
		"type":        validCreateChallengeRequestAnswer.Type,
		"answer":      validCreateChallengeRequestAnswer.Answer,
	}
	c := generateRequestWithBodyOnly(requestData)

	createChallengeRequest, err := ValidateCreateChallengeRequest(c)
	if err != nil {
		t.Error("Expected no error, got", err)
	}

	if createChallengeRequest.Name != validCreateChallengeRequestAnswer.Name {
		t.Error("Expected name to be", validCreateChallengeRequestAnswer.Name, "got", createChallengeRequest.Name)
	}
	if createChallengeRequest.Description != validCreateChallengeRequestAnswer.Description {
		t.Error("Expected description to be", validCreateChallengeRequestAnswer.Description, "got", createChallengeRequest.Description)
	}
	if createChallengeRequest.Points != validCreateChallengeRequestAnswer.Points {
		t.Error("Expected points to be", validCreateChallengeRequestAnswer.Points, "got", createChallengeRequest.Points)
	}
	if createChallengeRequest.Image != validCreateChallengeRequestAnswer.Image {
		t.Error("Expected image to be", validCreateChallengeRequestAnswer.Image, "got", createChallengeRequest.Image)
	}
	if createChallengeRequest.Type != validCreateChallengeRequestAnswer.Type {
		t.Error("Expected type to be", validCreateChallengeRequestAnswer.Type, "got", createChallengeRequest.Type)
	}
	if createChallengeRequest.Answer != validCreateChallengeRequestAnswer.Answer {
		t.Error("Expected answer to be", validCreateChallengeRequestAnswer.Answer, "got", createChallengeRequest.Answer)
	}
}

func TestValidateCreateChallengeRequestWithoutName(t *testing.T) {
	requestData := map[string]interface{}{
		"description": validCreateChallengeRequestUpload.Description,
		"points":      validCreateChallengeRequestUpload.Points,
		"Image":       validCreateChallengeRequestUpload.Image,
		"type":        validCreateChallengeRequestUpload.Type,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	if !apperrors.IsValidationError(err) {
		t.Error("Expected validation error, got", err)
	}

	if err.Error() != "Key: 'CreateChallengeRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag" {
		t.Error("Expected error message to be 'Key: 'CreateChallengeRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag', got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithoutDescription(t *testing.T) {
	requestData := map[string]interface{}{
		"name":   validCreateChallengeRequestUpload.Name,
		"points": validCreateChallengeRequestUpload.Points,
		"Image":  validCreateChallengeRequestUpload.Image,
		"type":   validCreateChallengeRequestUpload.Type,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	if !apperrors.IsValidationError(err) {
		t.Error("Expected validation error, got", err)
	}

	if err.Error() != "Key: 'CreateChallengeRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag" {
		t.Error("Expected error message to be 'Key: 'CreateChallengeRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag', got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithoutPoints(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestUpload.Name,
		"description": validCreateChallengeRequestUpload.Description,
		"Image":       validCreateChallengeRequestUpload.Image,
		"type":        validCreateChallengeRequestUpload.Type,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	if !apperrors.IsValidationError(err) {
		t.Error("Expected validation error, got", err)
	}

	if err.Error() != "Key: 'CreateChallengeRequest.Points' Error:Field validation for 'Points' failed on the 'required' tag" {
		t.Error("Expected error message to be 'Key: 'CreateChallengeRequest.Points' Error:Field validation for 'Points' failed on the 'required' tag', got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithoutImage(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestUpload.Name,
		"description": validCreateChallengeRequestUpload.Description,
		"points":      validCreateChallengeRequestUpload.Points,
		"type":        validCreateChallengeRequestUpload.Type,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	if !apperrors.IsValidationError(err) {
		t.Error("Expected validation error, got", err)
	}

	if err.Error() != "Key: 'CreateChallengeRequest.Image' Error:Field validation for 'Image' failed on the 'required' tag" {
		t.Error("Expected error message to be 'Key: 'CreateChallengeRequest.Image' Error:Field validation for 'Image' failed on the 'required' tag', got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithoutType(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestUpload.Name,
		"description": validCreateChallengeRequestUpload.Description,
		"points":      validCreateChallengeRequestUpload.Points,
		"Image":       validCreateChallengeRequestUpload.Image,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	if !apperrors.IsValidationError(err) {
		t.Error("Expected validation error, got", err)
	}

	if err.Error() != "Key: 'CreateChallengeRequest.Type' Error:Field validation for 'Type' failed on the 'required' tag" {
		t.Error("Expected error message to be 'Key: 'CreateChallengeRequest.Type' Error:Field validation for 'Type' failed on the 'required' tag', got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithoutAnswer(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestAnswer.Name,
		"description": validCreateChallengeRequestAnswer.Description,
		"points":      validCreateChallengeRequestAnswer.Points,
		"Image":       validCreateChallengeRequestAnswer.Image,
		"type":        validCreateChallengeRequestAnswer.Type,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	if !apperrors.IsValidationError(err) {
		t.Error("Expected validation error, got", err)
	}

	if err.Error() != "answer is required for answer question challenges" {
		t.Error("Expected error message to be 'answer is required for answer question challenges', got", err.Error())
	}
}

func TestValidateCreateChallengeRequestInvalidChallengeType(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestUpload.Name,
		"description": validCreateChallengeRequestUpload.Description,
		"points":      validCreateChallengeRequestUpload.Points,
		"Image":       validCreateChallengeRequestUpload.Image,
		"type":        "UPLOAD_ANSWER",
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	if !apperrors.IsValidationError(err) {
		t.Error("Expected validation error, got", err)
	}

	if err.Error() != "invalid challenge type" {
		t.Error("Expected error message to be 'invalid challenge type', got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithStringPoints(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestUpload.Name,
		"description": validCreateChallengeRequestUpload.Description,
		"points":      "10",
		"Image":       validCreateChallengeRequestUpload.Image,
		"type":        validCreateChallengeRequestUpload.Type,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	if !apperrors.IsValidationError(err) {
		t.Error("Expected validation error, got", err)
	}

	if err.Error() != "json: cannot unmarshal string into Go struct field CreateChallengeRequest.points of type uint" {
		t.Error("Expected error message to be 'json: cannot unmarshal string into Go struct field CreateChallengeRequest.points of type uint', got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithNegativePoints(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestUpload.Name,
		"description": validCreateChallengeRequestUpload.Description,
		"points":      -10,
		"Image":       validCreateChallengeRequestUpload.Image,
		"type":        validCreateChallengeRequestUpload.Type,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	if !apperrors.IsValidationError(err) {
		t.Error("Expected validation error, got", err)
	}

	expectedError := "json: cannot unmarshal number -10 into Go struct field CreateChallengeRequest.points of type uint"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithFloatPoints(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestUpload.Name,
		"description": validCreateChallengeRequestUpload.Description,
		"points":      10.5,
		"Image":       validCreateChallengeRequestUpload.Image,
		"type":        validCreateChallengeRequestUpload.Type,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	if !apperrors.IsValidationError(err) {
		t.Error("Expected validation error, got", err)
	}

	expectedError := "json: cannot unmarshal number 10.5 into Go struct field CreateChallengeRequest.points of type uint"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithZeroPoints(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestUpload.Name,
		"description": validCreateChallengeRequestUpload.Description,
		"points":      0,
		"Image":       validCreateChallengeRequestUpload.Image,
		"type":        validCreateChallengeRequestUpload.Type,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected no error, got nil")
		return
	}

	if !apperrors.IsValidationError(err) {
		t.Error("Expected validation error, got", err)
	}

	expectedError := "Key: 'CreateChallengeRequest.Points' Error:Field validation for 'Points' failed on the 'required' tag"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithEmptyBody(t *testing.T) {
	requestData := map[string]interface{}{}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	if !apperrors.IsValidationError(err) {
		t.Error("Expected validation error, got", err)
	}

	expectedError := "Key: 'CreateChallengeRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'CreateChallengeRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag\nKey: 'CreateChallengeRequest.Points' Error:Field validation for 'Points' failed on the 'required' tag\nKey: 'CreateChallengeRequest.Image' Error:Field validation for 'Image' failed on the 'required' tag\nKey: 'CreateChallengeRequest.Type' Error:Field validation for 'Type' failed on the 'required' tag"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithEmptyName(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        "",
		"description": validCreateChallengeRequestUpload.Description,
		"points":      validCreateChallengeRequestUpload.Points,
		"Image":       validCreateChallengeRequestUpload.Image,
		"type":        validCreateChallengeRequestUpload.Type,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	if !apperrors.IsValidationError(err) {
		t.Error("Expected validation error, got", err)
	}

	expectedError := "Key: 'CreateChallengeRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithEmptyDescription(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestUpload.Name,
		"description": "",
		"points":      validCreateChallengeRequestUpload.Points,
		"Image":       validCreateChallengeRequestUpload.Image,
		"type":        validCreateChallengeRequestUpload.Type,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	if !apperrors.IsValidationError(err) {
		t.Error("Expected validation error, got", err)
	}

	expectedError := "Key: 'CreateChallengeRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithEmptyImage(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestUpload.Name,
		"description": validCreateChallengeRequestUpload.Description,
		"points":      validCreateChallengeRequestUpload.Points,
		"Image":       "",
		"type":        validCreateChallengeRequestUpload.Type,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	if !apperrors.IsValidationError(err) {
		t.Error("Expected validation error, got", err)
	}

	expectedError := "Key: 'CreateChallengeRequest.Image' Error:Field validation for 'Image' failed on the 'required' tag"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithEmptyType(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestUpload.Name,
		"description": validCreateChallengeRequestUpload.Description,
		"points":      validCreateChallengeRequestUpload.Points,
		"Image":       validCreateChallengeRequestUpload.Image,
		"type":        "",
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	if !apperrors.IsValidationError(err) {
		t.Error("Expected validation error, got", err)
	}

	expectedError := "Key: 'CreateChallengeRequest.Type' Error:Field validation for 'Type' failed on the 'required' tag"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithEmptyAnswer(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestAnswer.Name,
		"description": validCreateChallengeRequestAnswer.Description,
		"points":      validCreateChallengeRequestAnswer.Points,
		"Image":       validCreateChallengeRequestAnswer.Image,
		"type":        validCreateChallengeRequestAnswer.Type,
		"answer":      "",
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	if !apperrors.IsValidationError(err) {
		t.Error("Expected validation error, got", err)
	}

	expectedError := "answer is required for answer question challenges"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithInvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBuffer([]byte("{"))
	c.Request = httptest.NewRequest("POST", "/challenges", body)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "unexpected EOF"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithNameAsNumber(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        10,
		"description": validCreateChallengeRequestUpload.Description,
		"points":      validCreateChallengeRequestUpload.Points,
		"Image":       validCreateChallengeRequestUpload.Image,
		"type":        validCreateChallengeRequestUpload.Type,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "json: cannot unmarshal number into Go struct field CreateChallengeRequest.name of type string"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithDescriptionAsNumber(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestUpload.Name,
		"description": 10,
		"points":      validCreateChallengeRequestUpload.Points,
		"Image":       validCreateChallengeRequestUpload.Image,
		"type":        validCreateChallengeRequestUpload.Type,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "json: cannot unmarshal number into Go struct field CreateChallengeRequest.description of type string"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithImageAsNumber(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestUpload.Name,
		"description": validCreateChallengeRequestUpload.Description,
		"points":      validCreateChallengeRequestUpload.Points,
		"Image":       10,
		"type":        validCreateChallengeRequestUpload.Type,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "json: cannot unmarshal number into Go struct field CreateChallengeRequest.image of type string"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithTypeAsNumber(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestUpload.Name,
		"description": validCreateChallengeRequestUpload.Description,
		"points":      validCreateChallengeRequestUpload.Points,
		"Image":       validCreateChallengeRequestUpload.Image,
		"type":        10,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "json: cannot unmarshal number into Go struct field CreateChallengeRequest.type of type types.ChallengeType"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithAnswerAsNumber(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestAnswer.Name,
		"description": validCreateChallengeRequestAnswer.Description,
		"points":      validCreateChallengeRequestAnswer.Points,
		"Image":       validCreateChallengeRequestAnswer.Image,
		"type":        validCreateChallengeRequestAnswer.Type,
		"answer":      10,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "json: cannot unmarshal number into Go struct field CreateChallengeRequest.answer of type string"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithInvalidImageUrl1(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestUpload.Name,
		"description": validCreateChallengeRequestUpload.Description,
		"points":      validCreateChallengeRequestUpload.Points,
		"Image":       "invalid",
		"type":        validCreateChallengeRequestUpload.Type,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "Key: 'CreateChallengeRequest.Image' Error:Field validation for 'Image' failed on the 'url' tag"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithInvalidImageUrl2(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestUpload.Name,
		"description": validCreateChallengeRequestUpload.Description,
		"points":      validCreateChallengeRequestUpload.Points,
		"Image":       "http://invalid",
		"type":        validCreateChallengeRequestUpload.Type,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "invalid image url"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateCreateChallengeRequestWithInvalidImageUrl3(t *testing.T) {
	requestData := map[string]interface{}{
		"name":        validCreateChallengeRequestUpload.Name,
		"description": validCreateChallengeRequestUpload.Description,
		"points":      validCreateChallengeRequestUpload.Points,
		"Image":       "image.com/fake.jpg",
		"type":        validCreateChallengeRequestUpload.Type,
	}
	c := generateRequestWithBodyOnly(requestData)

	_, err := ValidateCreateChallengeRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "Key: 'CreateChallengeRequest.Image' Error:Field validation for 'Image' failed on the 'url' tag"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateGetChallengeByIdRequestValid1(t *testing.T) {
	params := map[string]string{"id": "1"}
	c := generateRequestWithParamsOnly(params)

	id, err := ValidateGetChallengeByIdRequest(c)
	if err != nil {
		t.Error("Expected no error, got", err)
		return
	}

	if id != 1 {
		t.Error("Expected id to be 1, got", id)
	}
}

func TestValidateGetChallengeByIdRequestValid2(t *testing.T) {
	params := map[string]string{"id": "45345"}
	c := generateRequestWithParamsOnly(params)

	id, err := ValidateGetChallengeByIdRequest(c)
	if err != nil {
		t.Error("Expected no error, got", err)
		return
	}

	if id != 45345 {
		t.Error("Expected id to be 45345, got", id)
	}
}

func TestValidateGetChallengeByIdRequestInvalid(t *testing.T) {
	params := map[string]string{"id": "invalid"}
	c := generateRequestWithParamsOnly(params)

	_, err := ValidateGetChallengeByIdRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "invalid challenge id"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateGetChallengeByIdRequestWithEmptyParams(t *testing.T) {
	params := map[string]string{}
	c := generateRequestWithParamsOnly(params)

	_, err := ValidateGetChallengeByIdRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "invalid challenge id"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateVerifyAnswerRequest(t *testing.T) {
	params := map[string]string{"id": "1"}
	requestData := map[string]interface{}{
		"answer": validVerifyAnswerRequest.Answer,
	}
	c := generateRequestWithBodyAndParams(requestData, params)

	id, verifyAnswerRequest, err := ValidateVerifyAnswerRequest(c)
	if err != nil {
		t.Error("Expected no error, got", err)
		return
	}

	if id != 1 {
		t.Error("Expected id to be 1, got", id)
	}

	if verifyAnswerRequest.Answer != validVerifyAnswerRequest.Answer {
		t.Error("Expected answer to be", validVerifyAnswerRequest.Answer, "got", verifyAnswerRequest.Answer)
	}
}

func TestValidateVerifyAnswerRequestWithEmptyAnswer(t *testing.T) {
	params := map[string]string{"id": "1"}
	requestData := map[string]interface{}{
		"answer": "",
	}
	c := generateRequestWithBodyAndParams(requestData, params)

	_, _, err := ValidateVerifyAnswerRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "Key: 'VerifyAnswerRequest.Answer' Error:Field validation for 'Answer' failed on the 'required' tag"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateVerifyAnswerRequestWithEmptyParams(t *testing.T) {
	params := map[string]string{}
	requestData := map[string]interface{}{
		"answer": validVerifyAnswerRequest.Answer,
	}
	c := generateRequestWithBodyAndParams(requestData, params)

	_, _, err := ValidateVerifyAnswerRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "invalid challenge id"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateVerifyAnswerRequestWithInvalidId(t *testing.T) {
	params := map[string]string{"id": "invalid"}
	requestData := map[string]interface{}{
		"answer": validVerifyAnswerRequest.Answer,
	}
	c := generateRequestWithBodyAndParams(requestData, params)

	_, _, err := ValidateVerifyAnswerRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "invalid challenge id"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateVerifyAnswerRequestWithInvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBuffer([]byte("{"))
	c.Request = httptest.NewRequest("POST", "/challenges", body)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	_, _, err := ValidateVerifyAnswerRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "unexpected EOF"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateVerifyAnswerRequestWithAnswerAsNumber(t *testing.T) {
	params := map[string]string{"id": "1"}
	requestData := map[string]interface{}{
		"answer": 10,
	}
	c := generateRequestWithBodyAndParams(requestData, params)

	_, _, err := ValidateVerifyAnswerRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "json: cannot unmarshal number into Go struct field VerifyAnswerRequest.answer of type string"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateVerifyAnswerRequestWithMissingAnswer(t *testing.T) {
	params := map[string]string{"id": "1"}
	requestData := map[string]interface{}{}
	c := generateRequestWithBodyAndParams(requestData, params)

	_, _, err := ValidateVerifyAnswerRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "Key: 'VerifyAnswerRequest.Answer' Error:Field validation for 'Answer' failed on the 'required' tag"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateGetSubmissionsRequest(t *testing.T) {
	params := map[string]string{"id": "1"}
	c := generateRequestWithParamsOnly(params)

	id, err := ValidateGetSubmissionsRequest(c)
	if err != nil {
		t.Error("Expected no error, got", err)
		return
	}

	if id != 1 {
		t.Error("Expected id to be 1, got", id)
	}
}

func TestValidateGetSubmissionsRequestWithEmptyParams(t *testing.T) {
	params := map[string]string{}
	c := generateRequestWithParamsOnly(params)

	_, err := ValidateGetSubmissionsRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "invalid challenge id"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateGetSubmissionsRequestWithInvalidId(t *testing.T) {
	params := map[string]string{"id": "invalid"}
	c := generateRequestWithParamsOnly(params)

	_, err := ValidateGetSubmissionsRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "invalid challenge id"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateGetAnswerRequest(t *testing.T) {
	params := map[string]string{"id": "1"}
	c := generateRequestWithParamsOnly(params)

	id, err := ValidateGetAnswerRequest(c)
	if err != nil {
		t.Error("Expected no error, got", err)
		return
	}

	if id != 1 {
		t.Error("Expected id to be 1, got", id)
	}
}

func TestValidateGetAnswerRequestWithEmptyParams(t *testing.T) {
	params := map[string]string{}
	c := generateRequestWithParamsOnly(params)

	_, err := ValidateGetAnswerRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "invalid challenge id"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}

func TestValidateGetAnswerRequestWithInvalidId(t *testing.T) {
	params := map[string]string{"id": "invalid"}
	c := generateRequestWithParamsOnly(params)

	_, err := ValidateGetAnswerRequest(c)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	expectedError := "invalid challenge id"
	if err.Error() != expectedError {
		t.Error("Expected error message to be", expectedError, "got", err.Error())
	}
}
