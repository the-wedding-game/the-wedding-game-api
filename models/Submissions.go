package models

import (
	"github.com/jinzhu/gorm"
	"the-wedding-game-api/db"
)

type Submission struct {
	gorm.Model
	UserId      uint      `gorm:"not null;uniqueIndex:idx_user_challenge"`
	ChallengeID uint      `gorm:"not null;uniqueIndex:idx_user_challenge"`
	Answer      string    `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserId"`
	Challenge   Challenge `gorm:"foreignKey:ChallengeID"`
}

func NewSubmission(userId uint, challengeId uint, answer string) Submission {
	submission := Submission{
		UserId:      userId,
		ChallengeID: challengeId,
		Answer:      answer,
	}
	return submission
}

func (s *Submission) Save() (*Submission, error) {
	conn := db.GetDB()
	if err := conn.Create(s).Error; err != nil {
		return nil, err
	}
	return s, nil
}

func IsChallengeCompleted(userId uint, challengeId uint) (bool, error) {
	conn := db.GetDB()
	var submission Submission
	err := conn.Where("user_id = ? AND challenge_id = ?", userId, challengeId).First(&submission).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func GetCompletedChallenges(userId uint) ([]Submission, error) {
	conn := db.GetDB()
	var submissions []Submission
	if err := conn.Where("user_id = ?", userId).Find(&submissions).Error; err != nil {
		return nil, err
	}
	return submissions, nil
}
