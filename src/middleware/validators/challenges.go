package validators

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
	apperrors "the-wedding-game-api/errors"
	"the-wedding-game-api/types"
	"the-wedding-game-api/utils"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateCreateChallengeRequest(c *gin.Context) (types.CreateChallengeRequest, error) {
	var createChallengeRequest types.CreateChallengeRequest
	if err := c.BindJSON(&createChallengeRequest); err != nil {
		return types.CreateChallengeRequest{}, apperrors.NewValidationError(err.Error())
	}

	err := validate.Struct(&createChallengeRequest)
	if err != nil {
		return types.CreateChallengeRequest{}, apperrors.NewValidationError(err.Error())
	}

	if createChallengeRequest.Type != types.UploadPhotoChallenge && createChallengeRequest.Type != types.AnswerQuestionChallenge {
		return types.CreateChallengeRequest{}, apperrors.NewValidationError("invalid challenge type")
	}

	if createChallengeRequest.Type == types.AnswerQuestionChallenge && createChallengeRequest.Answer == "" {
		return types.CreateChallengeRequest{}, apperrors.NewValidationError("answer is required for answer question challenges")
	}

	if !utils.IsURLStrict(createChallengeRequest.Image) {
		return types.CreateChallengeRequest{}, apperrors.NewValidationError("invalid image url")
	}

	return createChallengeRequest, nil
}

func ValidateGetChallengeByIdRequest(c *gin.Context) (uint, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, apperrors.NewValidationError("invalid challenge id")
	}

	return uint(id), nil
}

func ValidateVerifyAnswerRequest(c *gin.Context) (uint, types.VerifyAnswerRequest, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, types.VerifyAnswerRequest{}, apperrors.NewValidationError("invalid challenge id")
	}

	var verifyAnswerRequest types.VerifyAnswerRequest
	if err := c.BindJSON(&verifyAnswerRequest); err != nil {
		return 0, types.VerifyAnswerRequest{}, apperrors.NewValidationError(err.Error())
	}

	if err := validate.Struct(&verifyAnswerRequest); err != nil {
		return 0, types.VerifyAnswerRequest{}, apperrors.NewValidationError(err.Error())
	}

	return uint(id), verifyAnswerRequest, nil
}
