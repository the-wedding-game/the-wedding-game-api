package models

import (
	"errors"
	"reflect"
	"testing"
	apperrors "the-wedding-game-api/errors"
	"the-wedding-game-api/types"
)

var (
	testSubmission1 = Submission{
		UserID:      132,
		ChallengeID: 3245,
		Answer:      "test_answer",
	}
	testSubmission2 = Submission{
		UserID:      235,
		ChallengeID: 9768,
		Answer:      "test_answer2",
	}
)

func TestNewSubmission(t *testing.T) {
	submission := NewSubmission(testSubmission1.UserID, testSubmission1.ChallengeID, testSubmission1.Answer)
	if submission.UserID != testSubmission1.UserID {
		t.Errorf("expected %d but got %d", testSubmission1.UserID, submission.UserID)
	}
	if submission.ChallengeID != testSubmission1.ChallengeID {
		t.Errorf("expected %d but got %d", testSubmission1.ChallengeID, submission.ChallengeID)
	}
	if submission.Answer != testSubmission1.Answer {
		t.Errorf("expected %s but got %s", testSubmission1.Answer, submission.Answer)
	}
}

func TestSubmissionSave(t *testing.T) {
	SetupMockDb()

	submission := &testSubmission1
	savedSubmission, err := submission.Save()
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}

	if savedSubmission.UserID != testSubmission1.UserID {
		t.Errorf("expected %d but got %d", testSubmission1.UserID, savedSubmission.UserID)
	}
	if savedSubmission.ChallengeID != testSubmission1.ChallengeID {
		t.Errorf("expected %d but got %d", testSubmission1.ChallengeID, savedSubmission.ChallengeID)
	}
	if savedSubmission.Answer != testSubmission1.Answer {
		t.Errorf("expected %s but got %s", testSubmission1.Answer, savedSubmission.Answer)
	}

	mockDb := GetConnection()
	var submissionFromDb Submission
	if err := mockDb.First(&submissionFromDb, savedSubmission.ID).GetError(); err != nil {
		t.Errorf("expected nil but got %v", err)
		return
	}
	if submissionFromDb.UserID != testSubmission1.UserID {
		t.Errorf("expected %d but got %d", testSubmission1.UserID, submissionFromDb.UserID)
	}
	if submissionFromDb.ChallengeID != testSubmission1.ChallengeID {
		t.Errorf("expected %d but got %d", testSubmission1.ChallengeID, submissionFromDb.ChallengeID)
	}
	if submissionFromDb.Answer != testSubmission1.Answer {
		t.Errorf("expected %s but got %s", testSubmission1.Answer, submissionFromDb.Answer)
	}
}

func TestSubmissionSaveError(t *testing.T) {
	mockDb := SetupMockDb()
	mockDb.Error = errors.New("test_error")

	_, err := testSubmission1.Save()
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}
	if !apperrors.IsDatabaseError(err) {
		t.Errorf("expected true but got false")
	}
	if err.Error() != "test_error" {
		t.Errorf("expected test_error but got %s", err.Error())
	}
}

func TestIsChallengeCompleted(t *testing.T) {
	SetupMockDb()

	submission := &testSubmission1
	_, err := submission.Save()
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}

	completed, err := IsChallengeCompleted(testSubmission1.UserID, testSubmission1.ChallengeID)
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}
	if !completed {
		t.Errorf("expected true but got false")
	}
}

func TestIsChallengeCompletedNotFound(t *testing.T) {
	SetupMockDb()

	isCompleted, err := IsChallengeCompleted(testSubmission1.UserID, testSubmission1.ChallengeID)
	if err != nil {
		t.Errorf("expected nil but got %v", err)
		return
	}
	if isCompleted {
		t.Errorf("expected false but got true")
	}
}

func TestIsChallengeCompletedError(t *testing.T) {
	mockDb := SetupMockDb()

	_, err := testSubmission1.Save()
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}

	mockDb.Error = errors.New("test_error")

	_, err = IsChallengeCompleted(testSubmission1.UserID, testSubmission1.ChallengeID)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}
	if !apperrors.IsDatabaseError(err) {
		t.Errorf("expected true but got false")
	}
	if err.Error() != "test_error" {
		t.Errorf("expected test_error but got %s", err.Error())
	}
}

func TestGetCompletedChallenges(t *testing.T) {
	SetupMockDb()

	_, err := testSubmission1.Save()
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}

	_, err = testSubmission2.Save()
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}

	submissions, err := GetCompletedChallenges(testSubmission1.UserID)
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}
	if len(submissions) != 2 {
		t.Errorf("expected 1 but got %d", len(submissions))
	}
	if submissions[0].UserID != testSubmission1.UserID {
		t.Errorf("expected %d but got %d", testSubmission1.UserID, submissions[0].UserID)
	}
	if submissions[0].ChallengeID != testSubmission1.ChallengeID {
		t.Errorf("expected %d but got %d", testSubmission1.ChallengeID, submissions[0].ChallengeID)
	}
	if submissions[0].Answer != testSubmission1.Answer {
		t.Errorf("expected %s but got %s", testSubmission1.Answer, submissions[0].Answer)
	}
	if submissions[1].UserID != testSubmission2.UserID {
		t.Errorf("expected %d but got %d", testSubmission2.UserID, submissions[1].UserID)
	}
	if submissions[1].ChallengeID != testSubmission2.ChallengeID {
		t.Errorf("expected %d but got %d", testSubmission2.ChallengeID, submissions[1].ChallengeID)
	}
	if submissions[1].Answer != testSubmission2.Answer {
		t.Errorf("expected %s but got %s", testSubmission2.Answer, submissions[1].Answer)
	}
}

func TestGetCompletedChallengesError(t *testing.T) {
	mockDb := SetupMockDb()

	_, err := testSubmission1.Save()
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}

	_, err = testSubmission2.Save()
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}

	mockDb.Error = errors.New("test_error")

	_, err = GetCompletedChallenges(testSubmission1.UserID)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}
	if err.Error() != "test_error" {
		t.Errorf("expected test_error but got %s", err.Error())
	}
}

func TestGetCompletedChallengesNotFound(t *testing.T) {
	SetupMockDb()

	submissions, err := GetCompletedChallenges(testSubmission1.UserID)
	if err != nil {
		t.Errorf("expected error but got nil")
		return
	}

	if len(submissions) != 0 {
		t.Errorf("expected 0 but got %d", len(submissions))
	}
}

func TestGetLeaderboard(t *testing.T) {
	SetupMockDb()

	leaderboard, err := GetLeaderboard()
	if err != nil {
		t.Errorf("expected nil but got %v", err)
		return
	}

	expectedLeaderboard := []types.LeaderboardEntry{
		{Username: "user1", Points: 100},
		{Username: "user2", Points: 200},
		{Username: "user3", Points: 300},
	}

	if !reflect.DeepEqual(leaderboard, expectedLeaderboard) {
		t.Errorf("expected %v but got %v", expectedLeaderboard, leaderboard)
	}
}

func TestGetLeaderboardError(t *testing.T) {
	mockDb := SetupMockDb()
	mockDb.Error = errors.New("test_error")

	_, err := GetLeaderboard()
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if !apperrors.IsDatabaseError(err) {
		t.Errorf("expected true but got false")
	}
	if err.Error() != "test_error" {
		t.Errorf("expected test_error but got %s", err.Error())
	}
}
