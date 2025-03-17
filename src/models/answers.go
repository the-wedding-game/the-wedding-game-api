package models

import (
	"gorm.io/gorm"
	"strconv"
	apperrors "the-wedding-game-api/errors"
	"the-wedding-game-api/types"
	"the-wedding-game-api/utils"
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
	conn := GetConnection()
	if err := conn.Create(&answer).GetError(); err != nil {
		return Answer{}, err
	}
	return answer, nil
}

func VerifyAnswer(challengeId uint, answer string) (bool, error) {
	conn := GetConnection()

	var challengeModel Challenge
	err := conn.Where("id = ?", challengeId).First(&challengeModel).GetError()
	if err != nil {
		if apperrors.IsRecordNotFoundError(err) {
			return false, apperrors.NewNotFoundError("Challenge", strconv.Itoa(int(challengeId)))
		}
		return false, err
	}

	if challengeModel.Type == types.AnswerQuestionChallenge {
		return verifyAnswerForQuestion(challengeId, answer)
	}

	if challengeModel.Type == types.UploadPhotoChallenge {
		return verifyAnswerForPhoto(answer)
	}

	return false, nil
}

func verifyAnswerForQuestion(challengeId uint, answer string) (bool, error) {
	conn := GetConnection()

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

func verifyAnswerForPhoto(answer string) (bool, error) {
	if !utils.IsURLStrict(answer) {
		return false, apperrors.NewValidationError("invalid image url")
	}

	return true, nil
}
