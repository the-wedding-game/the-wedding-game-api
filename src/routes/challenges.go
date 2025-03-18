package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"the-wedding-game-api/middleware"
	"the-wedding-game-api/middleware/validators"
	"the-wedding-game-api/models"
	"the-wedding-game-api/types"
)

func GetChallengeById(c *gin.Context) {
	user, err := middleware.GetCurrentUser(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	id, err := validators.ValidateGetChallengeByIdRequest(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	challenge, err := models.GetChallengeByID(id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	completed, err := models.IsChallengeCompleted(user.ID, challenge.ID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response := types.GetChallengeResponse{
		Id:          challenge.ID,
		Name:        challenge.Name,
		Description: challenge.Description,
		Points:      challenge.Points,
		Image:       challenge.Image,
		Status:      challenge.Status,
		Type:        challenge.Type,
		Completed:   completed,
	}

	c.IndentedJSON(http.StatusOK, response)
	return
}

func CreateChallenge(c *gin.Context) {
	if err := middleware.CheckIsAdmin(c); err != nil {
		_ = c.Error(err)
		return
	}

	challengeRequest, err := validators.ValidateCreateChallengeRequest(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	createdChallenge, err := models.CreateNewChallenge(challengeRequest)
	if err != nil {
		_ = c.Error(err)
		return
	}

	response := types.ChallengeCreatedResponse{
		Id:          createdChallenge.ID,
		Name:        createdChallenge.Name,
		Description: createdChallenge.Description,
		Points:      createdChallenge.Points,
		Image:       createdChallenge.Image,
		Status:      createdChallenge.Status,
		Type:        createdChallenge.Type,
	}
	c.IndentedJSON(http.StatusCreated, response)
	return
}

func GetAllChallenges(c *gin.Context) {
	user, err := middleware.GetCurrentUser(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	challengesArr, err := models.GetAllChallenges(false)
	if err != nil {
		_ = c.Error(err)
		return
	}

	submissions, err := models.GetCompletedChallenges(user.ID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var response types.GetChallengesResponse
	response.Challenges = make([]types.GetChallengeResponse, 0)
	for _, challenge := range challengesArr {
		isCompleted := models.IsChallengeInSubmissions(challenge.ID, submissions)

		response.Challenges = append(response.Challenges, types.GetChallengeResponse{
			Id:          challenge.ID,
			Name:        challenge.Name,
			Description: challenge.Description,
			Points:      challenge.Points,
			Image:       challenge.Image,
			Status:      challenge.Status,
			Type:        challenge.Type,
			Completed:   isCompleted,
		})
	}

	c.IndentedJSON(http.StatusOK, response)
	return
}

func GetAllChallengesAdmin(c *gin.Context) {
	challengesArr, err := models.GetAllChallenges(true)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var response types.GetChallengesAdminResponse
	response.Challenges = make([]types.GetChallengeAdminResponse, len(challengesArr))
	for i, challenge := range challengesArr {
		response.Challenges[i] = types.GetChallengeAdminResponse{
			Id:          challenge.ID,
			Name:        challenge.Name,
			Description: challenge.Description,
			Points:      challenge.Points,
			Image:       challenge.Image,
			Status:      challenge.Status,
			Type:        challenge.Type,
		}
	}

	c.IndentedJSON(http.StatusOK, response)
	return
}

func VerifyAnswer(c *gin.Context) {
	user, err := middleware.GetCurrentUser(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	challengeId, verifyAnswerRequest, err := validators.ValidateVerifyAnswerRequest(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	correct, err := models.VerifyAnswer(challengeId, verifyAnswerRequest.Answer)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if !correct {
		response := types.VerifyAnswerResponse{Correct: false}
		c.IndentedJSON(http.StatusOK, response)
		return
	}

	isAlreadyCompleted, err := models.IsChallengeCompleted(user.ID, challengeId)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if !isAlreadyCompleted {
		submission := models.NewSubmission(user.ID, challengeId, verifyAnswerRequest.Answer)
		_, err = submission.Save()
		if err != nil {
			_ = c.Error(err)
			return
		}
	}

	response := types.VerifyAnswerResponse{Correct: true}
	c.IndentedJSON(http.StatusOK, response)
	return
}
