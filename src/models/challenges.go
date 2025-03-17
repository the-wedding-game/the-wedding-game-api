package models

import (
	"gorm.io/gorm"
	"strconv"
	apperrors "the-wedding-game-api/errors"
	"the-wedding-game-api/types"
)

type Challenge struct {
	gorm.Model
	ID          uint                  `gorm:"primarykey"`
	Name        string                `gorm:"not null"`
	Description string                `gorm:"not null"`
	Points      uint                  `gorm:"not null"`
	Image       string                `gorm:"not null"`
	Type        types.ChallengeType   `gorm:"not null"`
	Status      types.ChallengeStatus `gorm:"default:'ACTIVE'"`
}

func NewChallenge(name string, description string, points uint, image string, _type types.ChallengeType,
	status types.ChallengeStatus) Challenge {
	challenge := Challenge{
		Name:        name,
		Description: description,
		Points:      points,
		Image:       image,
		Type:        _type,
		Status:      status,
	}
	return challenge
}

func CreateNewChallenge(createChallengeRequest types.CreateChallengeRequest) (Challenge, error) {
	challenge := NewChallenge(
		createChallengeRequest.Name,
		createChallengeRequest.Description,
		createChallengeRequest.Points,
		createChallengeRequest.Image,
		createChallengeRequest.Type,
		types.ActiveChallenge,
	)

	createdChallenge, err := challenge.Save()
	if err != nil {
		return Challenge{}, err
	}

	if createdChallenge.Type == types.AnswerQuestionChallenge {
		answer := NewAnswer(
			createdChallenge.ID,
			createChallengeRequest.Answer,
		)
		_, err := answer.Save()
		if err != nil {
			return Challenge{}, err
		}
	}

	return createdChallenge, nil
}

func (challenge Challenge) Save() (Challenge, error) {
	conn := GetConnection()
	if err := conn.Create(&challenge).GetError(); err != nil {
		return Challenge{}, err
	}
	return challenge, nil
}

func GetAllChallenges(showInactive bool) ([]Challenge, error) {
	conn := GetConnection()
	challenges, err := conn.GetAllChallenges(showInactive)
	if err != nil {
		return []Challenge{}, err
	}

	return challenges, nil
}

func GetChallengeByID(id uint) (Challenge, error) {
	conn := GetConnection()
	var challenge Challenge
	if err := conn.First(&challenge, id).GetError(); err != nil {
		if apperrors.IsRecordNotFoundError(err) {
			return Challenge{}, apperrors.NewNotFoundError("Challenge", strconv.Itoa(int(id)))
		}
		return Challenge{}, err
	}
	return challenge, nil
}
