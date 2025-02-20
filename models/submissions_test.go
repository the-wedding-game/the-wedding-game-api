package models

import (
	"errors"
	"testing"
	test "the-wedding-game-api/_test"
	"the-wedding-game-api/db"
)

var (
	testSubmission1 = Submission{
		UserId:      132,
		ChallengeID: 3245,
		Answer:      "test_answer",
	}
	testSubmission2 = Submission{
		UserId:      235,
		ChallengeID: 9768,
		Answer:      "test_answer2",
	}
)

func TestNewSubmission(t *testing.T) {
	submission := NewSubmission(testSubmission1.UserId, testSubmission1.ChallengeID, testSubmission1.Answer)
	if submission.UserId != testSubmission1.UserId {
		t.Errorf("expected %d but got %d", testSubmission1.UserId, submission.UserId)
	}
	if submission.ChallengeID != testSubmission1.ChallengeID {
		t.Errorf("expected %d but got %d", testSubmission1.ChallengeID, submission.ChallengeID)
	}
	if submission.Answer != testSubmission1.Answer {
		t.Errorf("expected %s but got %s", testSubmission1.Answer, submission.Answer)
	}
}

func TestSubmissionSave(t *testing.T) {
	test.SetupMockDb()

	submission := &testSubmission1
	savedSubmission, err := submission.Save()
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}

	if savedSubmission.UserId != testSubmission1.UserId {
		t.Errorf("expected %d but got %d", testSubmission1.UserId, savedSubmission.UserId)
	}
	if savedSubmission.ChallengeID != testSubmission1.ChallengeID {
		t.Errorf("expected %d but got %d", testSubmission1.ChallengeID, savedSubmission.ChallengeID)
	}
	if savedSubmission.Answer != testSubmission1.Answer {
		t.Errorf("expected %s but got %s", testSubmission1.Answer, savedSubmission.Answer)
	}

	mockDb := db.GetConnection()
	var submissionFromDb Submission
	if err := mockDb.First(&submissionFromDb, savedSubmission.ID).GetError(); err != nil {
		t.Errorf("expected nil but got %v", err)
		return
	}
	if submissionFromDb.UserId != testSubmission1.UserId {
		t.Errorf("expected %d but got %d", testSubmission1.UserId, submissionFromDb.UserId)
	}
	if submissionFromDb.ChallengeID != testSubmission1.ChallengeID {
		t.Errorf("expected %d but got %d", testSubmission1.ChallengeID, submissionFromDb.ChallengeID)
	}
	if submissionFromDb.Answer != testSubmission1.Answer {
		t.Errorf("expected %s but got %s", testSubmission1.Answer, submissionFromDb.Answer)
	}
}

func TestSubmissionSaveError(t *testing.T) {
	mockDb := test.SetupMockDb()
	mockDb.Error = errors.New("test_error")

	_, err := testSubmission1.Save()
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}
	if err.Error() != "test_error" {
		t.Errorf("expected test_error but got %s", err.Error())
	}
}

func TestIsChallengeCompleted(t *testing.T) {
	test.SetupMockDb()

	submission := &testSubmission1
	_, err := submission.Save()
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}

	completed, err := IsChallengeCompleted(testSubmission1.UserId, testSubmission1.ChallengeID)
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}
	if !completed {
		t.Errorf("expected true but got false")
	}
}

func TestIsChallengeCompletedNotFound(t *testing.T) {
	test.SetupMockDb()

	_, err := IsChallengeCompleted(testSubmission1.UserId, testSubmission1.ChallengeID)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}
	if err.Error() != "record not found: *models.Submission" {
		t.Errorf("expected record not found: *models.Submission but got %s", err.Error())
	}
}

func TestIsChallengeCompletedError(t *testing.T) {
	mockDb := test.SetupMockDb()

	_, err := testSubmission1.Save()
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}

	mockDb.Error = errors.New("test_error")

	_, err = IsChallengeCompleted(testSubmission1.UserId, testSubmission1.ChallengeID)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}
	if err.Error() != "test_error" {
		t.Errorf("expected test_error but got %s", err.Error())
	}
}

func TestGetCompletedChallenges(t *testing.T) {
	test.SetupMockDb()

	_, err := testSubmission1.Save()
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}

	_, err = testSubmission2.Save()
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}

	submissions, err := GetCompletedChallenges(testSubmission1.UserId)
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}
	if len(submissions) != 2 {
		t.Errorf("expected 1 but got %d", len(submissions))
	}
	if submissions[0].UserId != testSubmission1.UserId {
		t.Errorf("expected %d but got %d", testSubmission1.UserId, submissions[0].UserId)
	}
	if submissions[0].ChallengeID != testSubmission1.ChallengeID {
		t.Errorf("expected %d but got %d", testSubmission1.ChallengeID, submissions[0].ChallengeID)
	}
	if submissions[0].Answer != testSubmission1.Answer {
		t.Errorf("expected %s but got %s", testSubmission1.Answer, submissions[0].Answer)
	}
	if submissions[1].UserId != testSubmission2.UserId {
		t.Errorf("expected %d but got %d", testSubmission2.UserId, submissions[1].UserId)
	}
	if submissions[1].ChallengeID != testSubmission2.ChallengeID {
		t.Errorf("expected %d but got %d", testSubmission2.ChallengeID, submissions[1].ChallengeID)
	}
	if submissions[1].Answer != testSubmission2.Answer {
		t.Errorf("expected %s but got %s", testSubmission2.Answer, submissions[1].Answer)
	}
}

func TestGetCompletedChallengesError(t *testing.T) {
	mockDb := test.SetupMockDb()

	_, err := testSubmission1.Save()
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}

	_, err = testSubmission2.Save()
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}

	mockDb.Error = errors.New("test_error")

	_, err = GetCompletedChallenges(testSubmission1.UserId)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}
	if err.Error() != "test_error" {
		t.Errorf("expected test_error but got %s", err.Error())
	}
}

func TestGetCompletedChallengesNotFound(t *testing.T) {
	test.SetupMockDb()

	submissions, err := GetCompletedChallenges(testSubmission1.UserId)
	if err != nil {
		t.Errorf("expected error but got nil")
		return
	}

	if len(submissions) != 0 {
		t.Errorf("expected 0 but got %d", len(submissions))
	}
}
