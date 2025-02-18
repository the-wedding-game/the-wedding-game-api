package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"the-wedding-game-api/db"
	apperrors "the-wedding-game-api/errors"
)

type Answer struct {
	gorm.Model
	ChallengeID uint      `gorm:"not null"`
	Value       string    `gorm:"not null"`
	Challenge   Challenge `gorm:"foreignKey:ChallengeID"`
}

func NewAnswer(challengeId uint, value string) Answer {
	answer := Answer{
		ChallengeID: challengeId,
		Value:       value,
	}
	return answer
}

func (answer Answer) Save() (Answer, error) {
	conn := db.GetDB()
	if err := conn.Create(&answer).Error; err != nil {
		return Answer{}, err
	}
	return answer, nil
}

func VerifyAnswer(challengeId uint, answer string) (bool, error) {
	conn := db.GetDB()
	var answerModel Answer

	err := conn.Where("challenge_id = ?", challengeId).First(&answerModel).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, apperrors.NewNotFoundError("Challenge", strconv.Itoa(int(challengeId)))
		}
		return false, err
	}

	return answerModel.Value == answer, nil
}
