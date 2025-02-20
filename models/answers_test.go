package models

import (
	"errors"
	"testing"
	test "the-wedding-game-api/_tests"
	"the-wedding-game-api/db"
	apperrors "the-wedding-game-api/errors"
)

var (
	testAnswer1 = Answer{ChallengeID: 123, Value: "answer"}
	testAnswer2 = Answer{ChallengeID: 34324, Value: "another answer"}
)

func createTestAnswer(answer Answer) {
	database := db.GetConnection()
	database.Create(&answer)
}

func TestNewAnswer(t *testing.T) {
	answer := NewAnswer(testAnswer1.ChallengeID, testAnswer1.Value)
	if answer.ChallengeID != testAnswer1.ChallengeID {
		t.Errorf("expected %d but got %d", testAnswer1.ChallengeID, answer.ChallengeID)
	}
	if answer.Value != testAnswer1.Value {
		t.Errorf("expected %s but got %s", testAnswer1.Value, answer.Value)
	}

	answer = NewAnswer(testAnswer2.ChallengeID, testAnswer2.Value)
	if answer.ChallengeID != testAnswer2.ChallengeID {
		t.Errorf("expected %d but got %d", testAnswer2.ChallengeID, answer.ChallengeID)
	}
	if answer.Value != testAnswer2.Value {
		t.Errorf("expected %s but got %s", testAnswer2.Value, answer.Value)
	}
}

func TestAnswerSave(t *testing.T) {
	test.SetupMockDb()

	answer := NewAnswer(testAnswer1.ChallengeID, testAnswer1.Value)
	savedAnswer, err := answer.Save()
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if savedAnswer.ChallengeID != testAnswer1.ChallengeID {
		t.Errorf("expected %d but got %d", testAnswer1.ChallengeID, savedAnswer.ChallengeID)
	}
	if savedAnswer.Value != testAnswer1.Value {
		t.Errorf("expected %s but got %s", testAnswer1.Value, savedAnswer.Value)
	}

	mockDB := db.GetConnection()
	retrievedAnswer := Answer{}
	mockDB.First(&retrievedAnswer, testAnswer1.ID)
	if retrievedAnswer.ID != testAnswer1.ID {
		t.Errorf("expected %d but got %d", savedAnswer.ID, retrievedAnswer.ID)
	}
	if retrievedAnswer.ChallengeID != testAnswer1.ChallengeID {
		t.Errorf("expected %d but got %d", testAnswer1.ChallengeID, retrievedAnswer.ChallengeID)
	}
}

func TestAnswerSaveNegative(t *testing.T) {
	mockDB := test.SetupMockDb()
	mockDB.Error = errors.New("test_error")

	answer := NewAnswer(testAnswer1.ChallengeID, testAnswer1.Value)
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
	test.SetupMockDb()
	createTestAnswer(testAnswer1)

	isCorrect, err := VerifyAnswer(testAnswer1.ChallengeID, testAnswer1.Value)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if !isCorrect {
		t.Errorf("expected true but got false")
	}
}

func TestVerifyAnswerIncorrect(t *testing.T) {
	test.SetupMockDb()
	createTestAnswer(testAnswer1)

	isCorrect, err := VerifyAnswer(testAnswer1.ChallengeID, "incorrect answer")
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
		return
	}

	if isCorrect {
		t.Errorf("expected false but got true")
	}
}

func TestVerifyAnswerNotFound(t *testing.T) {
	test.SetupMockDb()

	_, err := VerifyAnswer(testAnswer1.ChallengeID, testAnswer1.Value)
	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	if !apperrors.IsNotFoundError(err) {
		t.Errorf("expected not found error but got %s", err.Error())
	}
	if err.Error() != "Challenge with key 123 not found." {
		t.Errorf("expected Challenge with key 123 not found. but got %s", err.Error())
	}
}

func TestVerifyAnswerError(t *testing.T) {
	mockDB := test.SetupMockDb()
	mockDB.Error = errors.New("test_error")

	createTestAnswer(testAnswer1)
	_, err := VerifyAnswer(testAnswer1.ChallengeID, testAnswer1.Value)
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
