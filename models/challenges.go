package models

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"the-wedding-game-api/db"
	apperrors "the-wedding-game-api/errors"
	"the-wedding-game-api/types"
)

type Challenge struct {
	gorm.Model
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
	conn := db.GetConnection()
	if err := conn.Create(&challenge).GetError(); err != nil {
		return Challenge{}, err
	}
	return challenge, nil
}

func GetAllChallenges() ([]Challenge, error) {
	conn := db.GetConnection()
	var challenges []Challenge
	if err := conn.Where("status = ?", types.ActiveChallenge).Find(&challenges).GetError(); err != nil {
		return nil, err
	}
	return challenges, nil
}

func GetChallengeByID(id uint) (Challenge, error) {
	conn := db.GetConnection()
	var challenge Challenge
	if err := conn.First(&challenge, id).GetError(); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return Challenge{}, apperrors.NewNotFoundError("Challenge", strconv.Itoa(int(id)))
		}
		return Challenge{}, err
	}
	return challenge, nil
}
