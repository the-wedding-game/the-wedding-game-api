package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"the-wedding-game-api/middleware"
	"the-wedding-game-api/models"
	"the-wedding-game-api/types"
)

func Login(c *gin.Context) {
	var loginRequest types.LoginRequest
	if err := c.BindJSON(&loginRequest); err != nil {
		_ = c.Error(err)
		return
	}

	exists, user, err := models.DoesUserExist(loginRequest.Username)
	if err != nil {
		_ = c.Error(err)
		return
	}

	if !exists {
		user = models.NewUser(loginRequest.Username)
		user, err = user.Save()
		if err != nil {
			_ = c.Error(err)
			return
		}
	}

	if user.Role == types.Admin {
		err := models.ValidatePassword(loginRequest.Password)
		if err != nil {
			_ = c.Error(err)
			return
		}
	}

	accessToken, err := models.LinkAccessTokenToUser(user.ID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, types.LoginResponse{
		User: types.UserResponse{
			Username: user.Username,
			Role:     user.Role,
		},
		AccessToken: accessToken.Token,
	})
	return
}

func GetCurrentUser(c *gin.Context) {
	user, err := middleware.GetCurrentUser(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, types.UserResponse{
		Username: user.Username,
		Role:     user.Role,
	})
	return
}
