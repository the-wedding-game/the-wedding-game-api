package validators

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
	"the-wedding-game-api/constants"
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
		return types.CreateChallengeRequest{}, apperrors.NewValidationError(constants.InvalidChallengeTypeError)
	}

	if createChallengeRequest.Type == types.AnswerQuestionChallenge && createChallengeRequest.Answer == "" {
		return types.CreateChallengeRequest{}, apperrors.NewValidationError(constants.AnswerRequiredError)
	}

	if !utils.IsURLStrict(createChallengeRequest.Image) {
		return types.CreateChallengeRequest{}, apperrors.NewValidationError(constants.InvalidImageURLError)
	}

	return createChallengeRequest, nil
}

func ValidateGetChallengeByIdRequest(c *gin.Context) (uint, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, apperrors.NewValidationError(constants.InvalidChallengeIDError)
	}

	return uint(id), nil
}

func ValidateUpdateChallengeRequest(c *gin.Context) (uint, types.UpdateChallengeRequest, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, types.UpdateChallengeRequest{}, apperrors.NewValidationError(constants.InvalidChallengeIDError)
	}

	var updateChallengeRequest types.UpdateChallengeRequest
	if err := c.BindJSON(&updateChallengeRequest); err != nil {
		return 0, types.UpdateChallengeRequest{}, apperrors.NewValidationError(err.Error())
	}

	err = validate.Struct(&updateChallengeRequest)
	if err != nil {
		return 0, types.UpdateChallengeRequest{}, apperrors.NewValidationError(err.Error())
	}

	if !utils.IsURLStrict(updateChallengeRequest.Image) {
		return 0, types.UpdateChallengeRequest{}, apperrors.NewValidationError(constants.InvalidImageURLError)
	}

	return uint(id), updateChallengeRequest, nil
}

func ValidateVerifyAnswerRequest(c *gin.Context) (uint, types.VerifyAnswerRequest, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, types.VerifyAnswerRequest{}, apperrors.NewValidationError(constants.InvalidChallengeIDError)
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

func ValidateGetSubmissionsRequest(c *gin.Context) (uint, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, apperrors.NewValidationError(constants.InvalidChallengeIDError)
	}

	return uint(id), nil
}

func ValidateGetAnswerRequest(c *gin.Context) (uint, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, apperrors.NewValidationError(constants.InvalidChallengeIDError)
	}

	return uint(id), nil
}
