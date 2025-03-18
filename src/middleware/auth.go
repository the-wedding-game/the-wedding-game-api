package middleware

import (
	"github.com/gin-gonic/gin"
	apperrors "the-wedding-game-api/errors"
	"the-wedding-game-api/models"
	"the-wedding-game-api/types"
)

func GetCurrentUser(c *gin.Context) (models.User, error) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		return models.User{}, apperrors.NewAuthenticationError("access token is not provided")
	}

	if len(accessToken) < 7 || accessToken[:7] != "Bearer " {
		return models.User{}, apperrors.NewAuthenticationError("invalid access token format")
	}

	accessToken = accessToken[7:]
	return models.GetUserByAccessToken(accessToken)
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
