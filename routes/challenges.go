package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"the-wedding-game-api/middleware"
	"the-wedding-game-api/middleware/validators"
	"the-wedding-game-api/models"
	"the-wedding-game-api/types"
)

func GetChallengeById(c *gin.Context) {
	id, err := validators.ValidateGetChallengeByIdRequest(c)
	if err != nil {
		return
	}

	challenge, err := models.GetChallengeByID(id)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	response := types.ChallengeCreatedResponse{
		Id:          challenge.ID,
		Name:        challenge.Name,
		Description: challenge.Description,
		Points:      challenge.Points,
		Image:       challenge.Image,
		Status:      challenge.Status,
		Type:        challenge.Type,
	}
	c.IndentedJSON(http.StatusOK, response)
	return
}

func CreateChallenge(c *gin.Context) {
	if middleware.CheckIsAdmin(c) != nil {
		return
	}

	challengeRequest, err := validators.ValidateCreateChallengeRequest(c)
	if err != nil {
		return
	}

	createdChallenge, err := models.CreateNewChallenge(challengeRequest)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
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
	challengesArr, err := models.GetAllChallenges()
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var response types.GetChallengesResponse
	for _, challenge := range challengesArr {
		response.Challenges = append(response.Challenges, types.ChallengeCreatedResponse{
			Id:          challenge.ID,
			Name:        challenge.Name,
			Description: challenge.Description,
			Points:      challenge.Points,
			Image:       challenge.Image,
			Status:      challenge.Status,
			Type:        challenge.Type,
		})
	}

	c.IndentedJSON(http.StatusOK, response)
	return
}
