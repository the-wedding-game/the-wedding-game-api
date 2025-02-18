package validators

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strconv"
	"the-wedding-game-api/types"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateCreateChallengeRequest(c *gin.Context) (types.CreateChallengeRequest, error) {
	var createChallengeRequest types.CreateChallengeRequest
	if err := c.BindJSON(&createChallengeRequest); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return types.CreateChallengeRequest{}, err
	}

	err := validate.Struct(&createChallengeRequest)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return types.CreateChallengeRequest{}, err
	}

	if createChallengeRequest.Type != types.UploadPhotoChallenge && createChallengeRequest.Type != types.AnswerQuestionChallenge {
		err := errors.New("invalid challenge type")
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return types.CreateChallengeRequest{}, err
	}

	if createChallengeRequest.Type == types.AnswerQuestionChallenge && createChallengeRequest.Answer == "" {
		err := errors.New("answer is required for answer question challenges")
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return types.CreateChallengeRequest{}, err
	}

	return createChallengeRequest, nil
}

func ValidateGetChallengeByIdRequest(c *gin.Context) (uint, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return 0, err
	}

	return uint(id), nil
}

func ValidateVerifyAnswerRequest(c *gin.Context) (uint, types.VerifyAnswerRequest, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return 0, types.VerifyAnswerRequest{}, err
	}

	var verifyAnswerRequest types.VerifyAnswerRequest
	if err := c.BindJSON(&verifyAnswerRequest); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return 0, types.VerifyAnswerRequest{}, err
	}

	if err := validate.Struct(&verifyAnswerRequest); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return 0, types.VerifyAnswerRequest{}, err
	}

	return uint(id), verifyAnswerRequest, nil
}
