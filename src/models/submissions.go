package models

import (
	"gorm.io/gorm"
	apperrors "the-wedding-game-api/errors"
	"the-wedding-game-api/types"
)

type Submission struct {
	gorm.Model
	UserID      uint   `gorm:"not null;uniqueIndex:idx_user_challenge"`
	ChallengeID uint   `gorm:"not null;uniqueIndex:idx_user_challenge"`
	Answer      string `gorm:"not null"`
	User        User
	Challenge   Challenge
}

func NewSubmission(userId uint, challengeId uint, answer string) Submission {
	submission := Submission{
		UserID:      userId,
		ChallengeID: challengeId,
		Answer:      answer,
	}
	return submission
}

func (s *Submission) Save() (*Submission, error) {
	conn := GetConnection()
	if err := conn.Create(s).GetError(); err != nil {
		return nil, err
	}
	return s, nil
}

func IsChallengeCompleted(userId uint, challengeId uint) (bool, error) {
	conn := GetConnection()
	var submission Submission
	err := conn.Where("user_id = ? AND challenge_id = ?", userId, challengeId).First(&submission).GetError()
	if err != nil {
		if apperrors.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func GetCompletedChallenges(userId uint) ([]Submission, error) {
	conn := GetConnection()
	var submissions []Submission
	if err := conn.Where("user_id = ?", userId).Find(&submissions).GetError(); err != nil {
		return nil, err
	}
	return submissions, nil
}

func GetLeaderboard() ([]types.LeaderboardEntry, error) {
	conn := GetConnection()
	leaderboard, err := conn.GetLeaderboard()
	if err != nil {
		return nil, err
	}
	return leaderboard, nil
}
