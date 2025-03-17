package models

import (
	"errors"
	"testing"
	apperrors "the-wedding-game-api/errors"
	"the-wedding-game-api/types"
)

var (
	testChallenge123 = Challenge{ID: 123, Name: "name", Description: "description", Points: 10, Type: types.AnswerQuestionChallenge, Status: types.ActiveChallenge}
	testAnswer123    = Answer{ChallengeID: 123, Value: "answer"}
	testAnswer34324  = Answer{ChallengeID: 34324, Value: "another answer"}
)

func createTestAnswer(answer Answer) {
	database := GetConnection()
	database.Create(&answer)
}

func createTestChallenge(challenge Challenge) {
	database := GetConnection()
	database.Create(&challenge)
}

func TestNewAnswer(t *testing.T) {
	answer := NewAnswer(testAnswer123.ChallengeID, testAnswer123.Value)
	if answer.ChallengeID != testAnswer123.ChallengeID {
		t.Errorf("expected %d but got %d", testAnswer123.ChallengeID, answer.ChallengeID)
	}
	if answer.Value != testAnswer123.Value {
		t.Errorf("expected %s but got %s", testAnswer123.Value, answer.Value)
	}

	answer = NewAnswer(testAnswer34324.ChallengeID, testAnswer34324.Value)
	if answer.ChallengeID != testAnswer34324.ChallengeID {
		t.Errorf("expected %d but got %d", testAnswer34324.ChallengeID, answer.ChallengeID)
	}
	if answer.Value != testAnswer34324.Value {
		t.Errorf("expected %s but got %s", testAnswer34324.Value, answer.Value)
	}
}

func TestAnswerSave(t *testing.T) {
	SetupMockDb()

	answer := NewAnswer(testAnswer123.ChallengeID, testAnswer123.Value)
	savedAnswer, err := answer.Save()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if savedAnswer.ChallengeID != testAnswer123.ChallengeID {
		t.Errorf("expected %d but got %d", testAnswer123.ChallengeID, savedAnswer.ChallengeID)
	}
	if savedAnswer.Value != testAnswer123.Value {
		t.Errorf("expected %s but got %s", testAnswer123.Value, savedAnswer.Value)
	}

	mockDB := GetConnection()
	retrievedAnswer := Answer{}
	mockDB.First(&retrievedAnswer, testAnswer123.ID)
	if retrievedAnswer.ID != testAnswer123.ID {
		t.Errorf("expected %d but got %d", savedAnswer.ID, retrievedAnswer.ID)
	}
	if retrievedAnswer.ChallengeID != testAnswer123.ChallengeID {
		t.Errorf("expected %d but got %d", testAnswer123.ChallengeID, retrievedAnswer.ChallengeID)
	}
}

func TestAnswerSaveNegative(t *testing.T) {
	mockDB := SetupMockDb()
	mockDB.Error = errors.New("test_error")

	answer := NewAnswer(testAnswer123.ChallengeID, testAnswer123.Value)
	_, err := answer.Save()
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

func TestVerifyAnswer(t *testing.T) {
	SetupMockDb()

	createTestChallenge(testChallenge123)
	createTestAnswer(testAnswer123)

	isCorrect, err := VerifyAnswer(testAnswer123.ChallengeID, testAnswer123.Value)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if !isCorrect {
		t.Errorf("expected true but got false")
	}
}

func TestVerifyAnswerIncorrect(t *testing.T) {
	SetupMockDb()
	createTestChallenge(testChallenge123)
	createTestAnswer(testAnswer123)

	isCorrect, err := VerifyAnswer(testAnswer123.ChallengeID, "incorrect answer")
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if isCorrect {
		t.Errorf("expected false but got true")
	}
}

func TestVerifyAnswerNotFound(t *testing.T) {
	SetupMockDb()
	createTestChallenge(testChallenge123)

	_, err := VerifyAnswer(testAnswer123.ChallengeID, testAnswer123.Value)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if !apperrors.IsNotFoundError(err) {
		t.Errorf("expected not found error but got %s", err.Error())
	}
	if err.Error() != "Answer with Challenge with key 123 not found." {
		t.Errorf("expected Answer with Challenge with key 123 not found not found. but got %s", err.Error())
	}
}

func TestVerifyAnswerError(t *testing.T) {
	mockDB := SetupMockDb()
	mockDB.Error = errors.New("test_error")

	createTestChallenge(testChallenge123)
	createTestAnswer(testAnswer123)
	_, err := VerifyAnswer(testAnswer123.ChallengeID, testAnswer123.Value)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if !apperrors.IsDatabaseError(err) {
		t.Errorf("expected database error but got %s", err.Error())
	}
	if err.Error() != "test_error" {
		t.Errorf("expected test_error but got %s", err.Error())
	}
}
