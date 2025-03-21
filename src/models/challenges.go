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

func (challenge Challenge) Update(updateChallengeRequest types.UpdateChallengeRequest) (Challenge, error) {
	hasSubmission, err := challenge.hasSubmissions()
	if err != nil {
		return Challenge{}, err
	}

	if hasSubmission {
		// Cannot update challenge type if submissions exist
		if updateChallengeRequest.Type != challenge.Type {
			return Challenge{}, apperrors.NewValidationError("Cannot update challenge type if submissions exist")
		}

		// Cannot update answer if submissions exist
		if updateChallengeRequest.Type == types.AnswerQuestionChallenge && updateChallengeRequest.Answer != "" {
			sameAnswer, err := verifyAnswerForQuestion(challenge.ID, updateChallengeRequest.Answer)
			if err != nil {
				return Challenge{}, err
			}
			if !sameAnswer {
				return Challenge{}, apperrors.NewValidationError("Cannot update answer if submissions exist")
			}
		}
	}

	conn := GetConnection()
	updatedChallenge, err := conn.UpdateChallenge(challenge, updateChallengeRequest)
	if err != nil {
		return Challenge{}, err
	}

	if updatedChallenge.Type == types.AnswerQuestionChallenge && updateChallengeRequest.Answer != "" {
		answer := NewAnswer(challenge.ID, updateChallengeRequest.Answer)

		// If challenge was previously an AnswerQuestionChallenge, update the answer
		if challenge.Type == types.AnswerQuestionChallenge {
			_, err := answer.Update()
			if err != nil {
				return Challenge{}, err
			}
		} else {
			// If challenge was previously an UploadPhotoChallenge, create a new answer
			_, err := answer.Save()
			if err != nil {
				return Challenge{}, err
			}
		}
	}

	// If challenge was previously an AnswerQuestionChallenge, delete the answer
	if updatedChallenge.Type == types.UploadPhotoChallenge && challenge.Type == types.AnswerQuestionChallenge {
		if err := DeleteAnswer(challenge.ID); err != nil {
			return Challenge{}, err
		}
	}

	return updatedChallenge, nil
}

func (challenge Challenge) hasSubmissions() (bool, error) {
	conn := GetConnection()
	return conn.HasSubmissions(challenge.ID)
}

func GetAllChallenges(showInactive bool) ([]Challenge, error) {
	conn := GetConnection()
	return conn.GetAllChallenges(showInactive)
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
