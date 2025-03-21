package models

import (
	"errors"
	"testing"
	"the-wedding-game-api/types"
)

var (
	testChallenge1 = Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      10,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}
	testChallenge2 = Challenge{
		Name:        "test_challenge2",
		Description: "test_description2",
		Points:      20,
		Image:       "test_image2",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.InactiveChallenge,
	}
)

func TestNewChallenge(t *testing.T) {
	challenge := NewChallenge(testChallenge1.Name, testChallenge1.Description, testChallenge1.Points, testChallenge1.Image,
		testChallenge1.Type, testChallenge1.Status)
	if challenge.Name != "test_challenge" {
		t.Errorf("expected test_challenge but got %s", challenge.Name)
	}
	if challenge.Description != "test_description" {
		t.Errorf("expected test_description but got %s", challenge.Description)
	}
	if challenge.Points != 10 {
		t.Errorf("expected 10 but got %d", challenge.Points)
	}
	if challenge.Image != "test_image" {
		t.Errorf("expected test_image but got %s", challenge.Image)
	}
	if challenge.Type != types.UploadPhotoChallenge {
		t.Errorf("expected UPLOAD_PHOTO_CHALLENGE but got %s", challenge.Type)
	}
	if challenge.Status != types.ActiveChallenge {
		t.Errorf("expected ACTIVE_CHALLENGE but got %s", challenge.Status)
	}
}

func TestCreateNewChallengeTypeUpload(t *testing.T) {
	SetupMockDb()

	createChallengeRequest := types.CreateChallengeRequest{
		Name:        testChallenge1.Name,
		Description: testChallenge1.Description,
		Points:      testChallenge1.Points,
		Image:       testChallenge1.Image,
		Type:        testChallenge1.Type,
	}

	challenge, err := CreateNewChallenge(createChallengeRequest)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if challenge.Name != testChallenge1.Name {
		t.Errorf("expected %s but got %s", testChallenge1.Name, challenge.Name)
	}
	if challenge.Description != testChallenge1.Description {
		t.Errorf("expected %s but got %s", testChallenge1.Description, challenge.Description)
	}
	if challenge.Points != testChallenge1.Points {
		t.Errorf("expected %d but got %d", testChallenge1.Points, challenge.Points)
	}
	if challenge.Image != testChallenge1.Image {
		t.Errorf("expected %s but got %s", testChallenge1.Image, challenge.Image)
	}
	if challenge.Type != testChallenge1.Type {
		t.Errorf("expected %s but got %s", testChallenge1.Type, challenge.Type)
	}

	challengeInDb, err := GetChallengeByID(challenge.ID)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}
	if challengeInDb.Name != testChallenge1.Name {
		t.Errorf("expected %s but got %s", testChallenge1.Name, challengeInDb.Name)
	}
	if challengeInDb.Description != testChallenge1.Description {
		t.Errorf("expected %s but got %s", testChallenge1.Description, challengeInDb.Description)
	}
	if challengeInDb.Points != testChallenge1.Points {
		t.Errorf("expected %d but got %d", testChallenge1.Points, challengeInDb.Points)
	}
	if challengeInDb.Image != testChallenge1.Image {
		t.Errorf("expected %s but got %s", testChallenge1.Image, challengeInDb.Image)
	}
	if challengeInDb.Type != testChallenge1.Type {
		t.Errorf("expected %s but got %s", testChallenge1.Type, challengeInDb.Type)
	}

	//Ensure an answer is not created for an upload photo challenge
	mockDb := GetConnection()
	var answer Answer
	err = mockDb.First(&answer, "challenge_id = ?", challenge.ID).GetError()
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}
	if err.Error() != "record not found: *models.Answer" {
		t.Errorf("expected record not found: *models.Answer but got %s", err.Error())
	}
}

func TestCreateNewChallengeTypeAnswer(t *testing.T) {
	SetupMockDb()

	createChallengeRequest := types.CreateChallengeRequest{
		Name:        testChallenge1.Name,
		Description: testChallenge1.Description,
		Points:      testChallenge1.Points,
		Image:       testChallenge1.Image,
		Type:        types.AnswerQuestionChallenge,
		Answer:      "test_answer",
	}

	challenge, err := CreateNewChallenge(createChallengeRequest)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if challenge.Name != testChallenge1.Name {
		t.Errorf("expected %s but got %s", testChallenge1.Name, challenge.Name)
	}
	if challenge.Description != testChallenge1.Description {
		t.Errorf("expected %s but got %s", testChallenge1.Description, challenge.Description)
	}
	if challenge.Points != testChallenge1.Points {
		t.Errorf("expected %d but got %d", testChallenge1.Points, challenge.Points)
	}
	if challenge.Image != testChallenge1.Image {
		t.Errorf("expected %s but got %s", testChallenge1.Image, challenge.Image)
	}
	if challenge.Type != types.AnswerQuestionChallenge {
		t.Errorf("expected %s but got %s", types.AnswerQuestionChallenge, challenge.Type)
	}

	challengeInDb, err := GetChallengeByID(challenge.ID)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}
	if challengeInDb.Name != testChallenge1.Name {
		t.Errorf("expected %s but got %s", testChallenge1.Name, challengeInDb.Name)
	}
	if challengeInDb.Description != testChallenge1.Description {
		t.Errorf("expected %s but got %s", testChallenge1.Description, challengeInDb.Description)
	}
	if challengeInDb.Points != testChallenge1.Points {
		t.Errorf("expected %d but got %d", testChallenge1.Points, challengeInDb.Points)
	}
	if challengeInDb.Image != testChallenge1.Image {
		t.Errorf("expected %s but got %s", testChallenge1.Image, challengeInDb.Image)
	}
	if challengeInDb.Type != types.AnswerQuestionChallenge {
		t.Errorf("expected %s but got %s", types.AnswerQuestionChallenge, challengeInDb.Type)
	}

	//Ensure answer is created in database
	mockDb := GetConnection()
	var answer Answer
	err = mockDb.First(&answer).GetError()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}
	if answer.Value != "test_answer" {
		t.Errorf("expected test_answer but got %s", answer.Value)
	}
}

func TestCreateNewChallengeError(t *testing.T) {
	mockDb := SetupMockDb()

	createChallengeRequest := types.CreateChallengeRequest{
		Name:        testChallenge1.Name,
		Description: testChallenge1.Description,
		Points:      testChallenge1.Points,
		Image:       testChallenge1.Image,
		Type:        testChallenge1.Type,
	}

	mockDb.Error = errors.New("test_error")
	_, err := CreateNewChallenge(createChallengeRequest)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if err.Error() != "test_error" {
		t.Errorf("expected test_error but got %s", err.Error())
	}
}

func TestChallengeSave(t *testing.T) {
	SetupMockDb()

	challenge := NewChallenge(testChallenge1.Name, testChallenge1.Description, testChallenge1.Points, testChallenge1.Image,
		testChallenge1.Type, testChallenge1.Status)
	savedChallenge, err := challenge.Save()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if savedChallenge.Name != testChallenge1.Name {
		t.Errorf("expected %s but got %s", testChallenge1.Name, savedChallenge.Name)
	}
	if savedChallenge.Description != testChallenge1.Description {
		t.Errorf("expected %s but got %s", testChallenge1.Description, savedChallenge.Description)
	}
	if savedChallenge.Points != testChallenge1.Points {
		t.Errorf("expected %d but got %d", testChallenge1.Points, savedChallenge.Points)
	}
	if savedChallenge.Image != testChallenge1.Image {
		t.Errorf("expected %s but got %s", testChallenge1.Image, savedChallenge.Image)
	}
	if savedChallenge.Type != testChallenge1.Type {
		t.Errorf("expected %s but got %s", testChallenge1.Type, savedChallenge.Type)
	}
	if savedChallenge.Status != testChallenge1.Status {
		t.Errorf("expected %s but got %s", testChallenge1.Status, savedChallenge.Status)
	}

	challengeInDb, err := GetChallengeByID(savedChallenge.ID)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}
	if challengeInDb.Name != testChallenge1.Name {
		t.Errorf("expected %s but got %s", testChallenge1.Name, challengeInDb.Name)
	}
	if challengeInDb.Description != testChallenge1.Description {
		t.Errorf("expected %s but got %s", testChallenge1.Description, challengeInDb.Description)
	}
	if challengeInDb.Points != testChallenge1.Points {
		t.Errorf("expected %d but got %d", testChallenge1.Points, challengeInDb.Points)
	}
	if challengeInDb.Image != testChallenge1.Image {
		t.Errorf("expected %s but got %s", testChallenge1.Image, challengeInDb.Image)
	}
	if challengeInDb.Type != testChallenge1.Type {
		t.Errorf("expected %s but got %s", testChallenge1.Type, challengeInDb.Type)
	}
	if challengeInDb.Status != testChallenge1.Status {
		t.Errorf("expected %s but got %s", testChallenge1.Status, challengeInDb.Status)
	}
}

func TestChallengeSaveError(t *testing.T) {
	mockDb := SetupMockDb()

	challenge := NewChallenge(testChallenge1.Name, testChallenge1.Description, testChallenge1.Points, testChallenge1.Image,
		testChallenge1.Type, testChallenge1.Status)
	mockDb.Error = errors.New("test_error")
	_, err := challenge.Save()
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if err.Error() != "test_error" {
		t.Errorf("expected test_error but got %s", err.Error())
	}
}

func TestGetAllChallenges(t *testing.T) {
	SetupMockDb()

	_, err := testChallenge1.Save()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	_, err = testChallenge2.Save()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	challenges, err := GetAllChallenges(false)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if len(challenges) != 2 {
		t.Errorf("expected 2 but got %d", len(challenges))
		return
	}
	if challenges[0].Name != testChallenge1.Name {
		t.Errorf("expected test_challenge1 but got %s", challenges[0].Name)
	}
	if challenges[1].Name != testChallenge2.Name {
		t.Errorf("expected test_challenge2 but got %s", challenges[1].Name)
	}
}

func TestGetAllChallengesError(t *testing.T) {
	mockDb := SetupMockDb()

	mockDb.Error = errors.New("test_error")
	_, err := GetAllChallenges(false)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if err.Error() != "test_error" {
		t.Errorf("expected test_error but got %s", err.Error())
	}
}

func TestGetAllChallengesEmpty(t *testing.T) {
	SetupMockDb()

	challenges, err := GetAllChallenges(false)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if len(challenges) != 0 {
		t.Errorf("expected 0 but got %d", len(challenges))
	}
}

func TestGetChallengeByIDTypeUpload(t *testing.T) {
	SetupMockDb()

	_, err := testChallenge1.Save()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	challenge, err := GetChallengeByID(testChallenge1.ID)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if challenge.Name != testChallenge1.Name {
		t.Errorf("expected test_challenge1 but got %s", challenge.Name)
	}
	if challenge.Description != testChallenge1.Description {
		t.Errorf("expected test_description but got %s", challenge.Description)
	}
	if challenge.Points != testChallenge1.Points {
		t.Errorf("expected 10 but got %d", challenge.Points)
	}
	if challenge.Image != testChallenge1.Image {
		t.Errorf("expected test_image but got %s", challenge.Image)
	}
	if challenge.Type != types.UploadPhotoChallenge {
		t.Errorf("expected UPLOAD_PHOTO_CHALLENGE but got %s", challenge.Type)
	}
	if challenge.Status != types.ActiveChallenge {
		t.Errorf("expected ACTIVE_CHALLENGE but got %s", challenge.Status)
	}
}

func TestGetChallengeByIDTypeAnswer(t *testing.T) {
	SetupMockDb()

	_, err := testChallenge2.Save()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	challenge, err := GetChallengeByID(testChallenge2.ID)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if challenge.Name != testChallenge2.Name {
		t.Errorf("expected test_challenge2 but got %s", challenge.Name)
	}
	if challenge.Description != testChallenge2.Description {
		t.Errorf("expected test_description2 but got %s", challenge.Description)
	}
	if challenge.Points != testChallenge2.Points {
		t.Errorf("expected 20 but got %d", challenge.Points)
	}
	if challenge.Image != testChallenge2.Image {
		t.Errorf("expected test_image2 but got %s", challenge.Image)
	}
	if challenge.Type != types.AnswerQuestionChallenge {
		t.Errorf("expected ANSWER_QUESTION_CHALLENGE but got %s", challenge.Type)
	}
	if challenge.Status != types.InactiveChallenge {
		t.Errorf("expected INACTIVE_CHALLENGE but got %s", challenge.Status)
	}
}

func TestGetChallengeByIDError(t *testing.T) {
	mockDb := SetupMockDb()
	_, err := testChallenge1.Save()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	mockDb.Error = errors.New("test_error")
	_, err = GetChallengeByID(1)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if err.Error() != "test_error" {
		t.Errorf("expected test_error but got %s", err.Error())
	}
}

func TestGetChallengeByIDNotFound(t *testing.T) {
	SetupMockDb()

	_, err := GetChallengeByID(1)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if err.Error() != "Challenge with key 1 not found." {
		t.Errorf("expected Challenge with key 1 not found. but got %s", err.Error())
	}
}

func TestChallengeHasSubmissionsTrue(t *testing.T) {
	mockDb := SetupMockDb()

	challenge, err := testChallenge1.Save()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	submission := NewSubmission(1, testChallenge1.ID, "test_answer")
	_, err = mockDb.AddSubmission(submission)
	if err != nil {
		t.Errorf("Error adding submission %s", err.Error())
		return
	}

	hasSubmissions, err := challenge.hasSubmissions()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if !hasSubmissions {
		t.Errorf("expected true but got false")
	}
}

func TestChallengeHasSubmissionsFalse(t *testing.T) {
	SetupMockDb()

	challenge, err := testChallenge1.Save()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	hasSubmissions, err := challenge.hasSubmissions()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if hasSubmissions {
		t.Errorf("expected false but got true")
	}
}

func TestChallengeHasSubmissionsError(t *testing.T) {
	mockDb := SetupMockDb()

	challenge, err := testChallenge1.Save()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	mockDb.Error = errors.New("test_error")
	_, err = challenge.hasSubmissions()
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if err.Error() != "test_error" {
		t.Errorf("expected test_error but got %s", err.Error())
	}
}

func TestChallengeUpdateHasSubmissions(t *testing.T) {
	mockDb := SetupMockDb()

	challenge, err := testChallenge1.Save()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	submission := NewSubmission(1, challenge.ID, "test_answer")
	_, err = mockDb.AddSubmission(submission)
	if err != nil {
		t.Errorf("error adding submission %v", err.Error())
		return
	}

	updateChallengeRequest := types.UpdateChallengeRequest{
		Name:        "test_challenge_updated",
		Description: "test_description_updated",
		Points:      100,
		Image:       "test_image_updated",
		Type:        types.AnswerQuestionChallenge,
		Answer:      "test_answer_updated",
	}

	_, err = challenge.Update(updateChallengeRequest)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if err.Error() != "Cannot update challenge type if submissions exist" {
		t.Errorf("expected Cannot update challenge type if submissions exist but got %s", err.Error())
	}
}

func TestChallengeUpdateHasNoSubmissions(t *testing.T) {
	SetupMockDb()

	challenge, err := testChallenge1.Save()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	updateChallengeRequest := types.UpdateChallengeRequest{
		Name:        "test_challenge_updated",
		Description: "test_description_updated",
		Points:      100,
		Image:       "test_image_updated",
		Type:        types.UploadPhotoChallenge,
		Status:      types.InactiveChallenge,
	}

	updatedChallenge, err := challenge.Update(updateChallengeRequest)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if updatedChallenge.Name != "test_challenge_updated" {
		t.Errorf("expected test_challenge_updated but got %s", updatedChallenge.Name)
	}
	if updatedChallenge.Description != "test_description_updated" {
		t.Errorf("expected test_description_updated but got %s", updatedChallenge.Description)
	}
	if updatedChallenge.Points != 100 {
		t.Errorf("expected 100 but got %d", updatedChallenge.Points)
	}
	if updatedChallenge.Image != "test_image_updated" {
		t.Errorf("expected test_image_updated but got %s", updatedChallenge.Image)
	}
	if updatedChallenge.Type != types.UploadPhotoChallenge {
		t.Errorf("expected ANSWER_QUESTION_CHALLENGE but got %s", updatedChallenge.Type)
	}
	if updatedChallenge.Status != types.InactiveChallenge {
		t.Errorf("expected ACTIVE_CHALLENGE but got %s", updatedChallenge.Status)
	}
}

func TestUpdateChallengeWithSubmissionsButNoBreakingChanges(t *testing.T) {
	mockDb := SetupMockDb()

	challenge, err := testChallenge1.Save()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	submission := NewSubmission(1, challenge.ID, "test_answer")
	_, err = mockDb.AddSubmission(submission)
	if err != nil {
		t.Errorf("error adding submission %v", err.Error())
		return
	}

	updateChallengeRequest := types.UpdateChallengeRequest{
		Name:        "test_challenge_updated",
		Description: "test_description_updated",
		Points:      100,
		Image:       "test_image_updated",
		Type:        types.UploadPhotoChallenge,
		Status:      types.InactiveChallenge,
	}

	updatedChallenge, err := challenge.Update(updateChallengeRequest)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if updatedChallenge.Name != "test_challenge_updated" {
		t.Errorf("expected test_challenge_updated but got %s", updatedChallenge.Name)
	}
	if updatedChallenge.Description != "test_description_updated" {
		t.Errorf("expected test_description_updated but got %s", updatedChallenge.Description)
	}
	if updatedChallenge.Points != 100 {
		t.Errorf("expected 100 but got %d", updatedChallenge.Points)
	}
	if updatedChallenge.Image != "test_image_updated" {
		t.Errorf("expected test_image_updated but got %s", updatedChallenge.Image)
	}
	if updatedChallenge.Type != types.UploadPhotoChallenge {
		t.Errorf("expected UPLOAD_PHOTO_CHALLENGE but got %s", updatedChallenge.Type)
	}
	if updatedChallenge.Status != types.InactiveChallenge {
		t.Errorf("expected INACTIVE_CHALLENGE but got %s", updatedChallenge.Status)
	}
}

func TestUpdateChallengeWithSubmissionsButSameAnswer(t *testing.T) {
	SetupMockDb()

	createChallengeRequest := types.CreateChallengeRequest{
		Name:        testChallenge1.Name,
		Description: testChallenge1.Description,
		Points:      testChallenge1.Points,
		Image:       testChallenge1.Image,
		Type:        types.AnswerQuestionChallenge,
		Answer:      "test_answer",
	}

	challenge, err := CreateNewChallenge(createChallengeRequest)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	submission := NewSubmission(1, challenge.ID, "test_answer")
	_, err = submission.Save()
	if err != nil {
		t.Errorf("error adding submission %v", err.Error())
		return
	}

	updateChallengeRequest := types.UpdateChallengeRequest{
		Name:        "test_challenge_updated",
		Description: "test_description_updated",
		Points:      100,
		Image:       "test_image_updated",
		Type:        types.AnswerQuestionChallenge,
		Answer:      "test_answer",
		Status:      types.InactiveChallenge,
	}

	updatedChallenge, err := challenge.Update(updateChallengeRequest)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if updatedChallenge.Name != "test_challenge_updated" {
		t.Errorf("expected test_challenge_updated but got %s", updatedChallenge.Name)
	}

	if updatedChallenge.Description != "test_description_updated" {
		t.Errorf("expected test_description_updated but got %s", updatedChallenge.Description)
	}

	if updatedChallenge.Points != 100 {
		t.Errorf("expected 100 but got %d", updatedChallenge.Points)
	}

	if updatedChallenge.Image != "test_image_updated" {
		t.Errorf("expected test_image_updated but got %s", updatedChallenge.Image)
	}

	if updatedChallenge.Type != types.AnswerQuestionChallenge {
		t.Errorf("expected ANSWER_QUESTION_CHALLENGE but got %s", updatedChallenge.Type)
	}

	if updatedChallenge.Status != types.InactiveChallenge {
		t.Errorf("expected ACTIVE_CHALLENGE but got %s", updatedChallenge.Status)
	}
}
