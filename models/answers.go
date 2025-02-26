package models

import (
	"gorm.io/gorm"
	"strconv"
	"the-wedding-game-api/db"
	apperrors "the-wedding-game-api/errors"
)

type Answer struct {
	gorm.Model
	ChallengeID uint   `gorm:"not null;unique"`
	Value       string `gorm:"not null"`
	Challenge   Challenge
}

func NewAnswer(challengeId uint, value string) Answer {
	answer := Answer{
		ChallengeID: challengeId,
		Value:       value,
	}
	return answer
}

func (answer Answer) Save() (Answer, error) {
	conn := db.GetConnection()
	if err := conn.Create(&answer).GetError(); err != nil {
		return Answer{}, err
	}
	return answer, nil
}

func VerifyAnswer(challengeId uint, answer string) (bool, error) {
	conn := db.GetConnection()
	var answerModel Answer

	err := conn.Where("challenge_id = ?", challengeId).First(&answerModel).GetError()
	if err != nil {
		if apperrors.IsRecordNotFoundError(err) {
			return false, apperrors.NewNotFoundError("Answer with Challenge", strconv.Itoa(int(challengeId)))
		}
		return false, err
	}

	return answerModel.Value == answer, nil
}
