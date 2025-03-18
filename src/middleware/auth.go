package middleware

import (
	"github.com/gin-gonic/gin"
	apperrors "the-wedding-game-api/errors"
	"the-wedding-game-api/models"
	"the-wedding-game-api/types"
)

func GetCurrentUser(c *gin.Context) (models.User, error) {
	user, exists := c.Get("user")

	if exists {
		return user.(models.User), nil
	}

	user, err := parseAuthorizationForUser(c)
	if err != nil {
		return models.User{}, err
	}

	return user.(models.User), nil
}

func CheckIsAdmin(c *gin.Context) error {
	user, err := GetCurrentUser(c)
	if err != nil {
		return err
	}

	if user.Role != types.Admin {
		return apperrors.NewAuthorizationError()
	}
	return nil
}

func IsAdmin(c *gin.Context) {
	if err := CheckIsAdmin(c); err != nil {
		handleError(c, err)
		return
	}
	c.Next()
}

func IsLoggedIn(c *gin.Context) {
	_, err := GetCurrentUser(c)
	if err != nil {
		handleError(c, err)
	}

	c.Next()
}

func parseAuthorizationForUser(c *gin.Context) (models.User, error) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		return models.User{}, apperrors.NewAuthenticationError("access token is not provided")
	}

	if len(accessToken) < 7 || accessToken[:7] != "Bearer " {
		return models.User{}, apperrors.NewAuthenticationError("invalid access token format")
	}

	accessToken = accessToken[7:]

	user, err := models.GetUserByAccessToken(accessToken)
	if err != nil {
		return models.User{}, err
	}

	c.Set("user", user)
	return user, nil
}
