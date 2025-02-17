package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	models2 "the-wedding-game-api/models"
)

func GetCurrentUser(c *gin.Context) (models2.User, error) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		return models2.User{}, errors.New("access token not provided")
	}

	if len(accessToken) < 7 || accessToken[:7] != "Bearer " {
		return models2.User{}, errors.New("invalid access token")
	}

	accessToken = accessToken[7:]
	return models2.GetUserByAccessToken(accessToken)
}

func CheckIsAdmin(c *gin.Context) error {
	user, err := GetCurrentUser(c)
	if err != nil || user.Role != models2.Admin {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Access denied"})
		return errors.New("access denied")
	}
	return nil
}
