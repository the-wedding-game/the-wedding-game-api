package models

import (
	"github.com/jinzhu/gorm"
	"the-wedding-game-api/db"
)

type Answer struct {
	gorm.Model
	ChallengeID uint   `gorm:"not null"`
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
	conn := db.GetDB()
	if err := conn.Create(&answer).Error; err != nil {
		return Answer{}, err
	}
	return answer, nil
}
