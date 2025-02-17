package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"the-wedding-game-api/middleware"
	"the-wedding-game-api/models"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User        models.User `json:"user"`
	AccessToken string      `json:"access_token"`
}

func Login(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.BindJSON(&loginRequest); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	exists, user, err := models.DoesUserExist(loginRequest.Username)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if exists {
		if user.Role == models.Admin {
			c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Admins are not allowed to login"})
			return
		}

		accessToken, err := models.LinkAccessTokenToUser(user.ID)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, LoginResponse{
			User:        user,
			AccessToken: accessToken.Token,
		})
		return
	} else {
		user := models.NewUser(loginRequest.Username)
		createdUser, err := user.Save()
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		accessToken, err := models.LinkAccessTokenToUser(user.ID)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, LoginResponse{
			User:        createdUser,
			AccessToken: accessToken.Token,
		})
		return
	}
}

func GetCurrentUser(c *gin.Context) {
	user, err := middleware.GetCurrentUser(c)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
	return
}
