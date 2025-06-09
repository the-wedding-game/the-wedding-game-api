package models

import (
	"fmt"
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

func NewChallenge(name string, description string, points uint, image string, challengeType types.ChallengeType,
	status types.ChallengeStatus) Challenge {
	challenge := Challenge{
		Name:        name,
		Description: description,
		Points:      points,
		Image:       image,
		Type:        challengeType,
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
	if err := challenge.checkForInvalidUpdateFields(updateChallengeRequest); err != nil {
		return Challenge{}, err
	}

	conn := GetConnection()
	updatedChallenge, err := conn.UpdateChallenge(challenge, updateChallengeRequest)
	if err != nil {
		return Challenge{}, err
	}

	if err := updatedChallenge.updateUnderlyingAnswer(challenge.Type, updateChallengeRequest.Answer); err != nil {
		return Challenge{}, err
	}

	return updatedChallenge, nil
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

func (challenge Challenge) hasSubmissions() (bool, error) {
	conn := GetConnection()
	return conn.HasSubmissions(challenge.ID)
}

func (challenge Challenge) checkForInvalidUpdateFields(updateChallengeRequest types.UpdateChallengeRequest) error {
	hasSubmission, err := challenge.hasSubmissions()
	if err != nil {
		return fmt.Errorf("error while retrieving submissions: %w", err)
	}

	if hasSubmission {
		// Cannot update challenge type if submissions exist
		if updateChallengeRequest.Type != challenge.Type {
			return apperrors.NewValidationError("Cannot update challenge type if submissions exist")
		}

		// Cannot update answer if submissions exist
		if updateChallengeRequest.Type == types.AnswerQuestionChallenge && updateChallengeRequest.Answer != "" {
			sameAnswer, err := verifyAnswerForQuestion(challenge.ID, updateChallengeRequest.Answer)
			if err != nil {
				return fmt.Errorf("error verifying answer: %w", err)
			}
			if !sameAnswer {
				return apperrors.NewValidationError("Cannot update answer if submissions exist")
			}
		}
	}

	if challenge.Type == types.UploadPhotoChallenge && updateChallengeRequest.Type == types.AnswerQuestionChallenge && updateChallengeRequest.Answer == "" {
		return apperrors.NewValidationError("Answer cannot be empty when changing to AnswerQuestion challenge type")
	}

	return nil
}

func (challenge Challenge) updateUnderlyingAnswer(oldType types.ChallengeType, answer string) error {
	if challenge.Type == types.AnswerQuestionChallenge {
		if err := challenge.createOrUpdateAnswer(oldType, answer); err != nil {
			return fmt.Errorf("error while creating or updating answer: %w", err)
		}
	}

	// If challenge was previously an AnswerQuestionChallenge, delete the answer
	if challenge.Type == types.UploadPhotoChallenge && oldType == types.AnswerQuestionChallenge {
		if err := DeleteAnswer(challenge.ID); err != nil {
			return err
		}
	}

	return nil
}

func (challenge Challenge) createOrUpdateAnswer(oldType types.ChallengeType, answer string) error {
	// If challenge was previously an answer question challenge, update the answer
	if oldType == types.AnswerQuestionChallenge && answer != "" {
		answer := NewAnswer(challenge.ID, answer)
		_, err := answer.Update()
		if err != nil {
			return fmt.Errorf("error while updating answer for challenge: %w", err)
		}
	}

	if oldType == types.UploadPhotoChallenge {
		// If challenge was previously an UploadPhotoChallenge, create a new answer
		answer := NewAnswer(challenge.ID, answer)
		_, err := answer.Save()
		if err != nil {
			return fmt.Errorf("error while creating answer for challenge: %w", err)
		}
	}

	return nil
}

func (challenge Challenge) Delete() error {
	conn := GetConnection()

	if err := conn.DeleteAnswerForChallenge(challenge.ID); err != nil {
		return fmt.Errorf("error deleting answer for challenge: %w", err)
	}

	if err := conn.DeleteSubmissionsForChallenge(challenge.ID); err != nil {
		return fmt.Errorf("error deleting submissions for challenge: %w", err)
	}

	if err := conn.DeleteChallenge(challenge.ID); err != nil {
		return fmt.Errorf("error deleting challenge: %w", err)
	}

	return nil
}
