package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"the-wedding-game-api/models"
)

func GetCurrentUser(c *gin.Context) (models.User, error) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		return models.User{}, errors.New("access token not provided")
	}

	if len(accessToken) < 7 || accessToken[:7] != "Bearer " {
		return models.User{}, errors.New("invalid access token")
	}

	accessToken = accessToken[7:]
	return models.GetUserByAccessToken(accessToken)
}

func CheckIsLoggedIn(c *gin.Context) error {
	_, err := GetCurrentUser(c)
	if err != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Access denied"})
		return errors.New("access denied")
	}
	return nil
}

func CheckIsAdmin(c *gin.Context) error {
	user, err := GetCurrentUser(c)
	if err != nil || user.Role != models.Admin {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Access denied"})
		return errors.New("access denied")
	}
	return nil
}
