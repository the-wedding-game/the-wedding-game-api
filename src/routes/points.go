package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"the-wedding-game-api/middleware"
	"the-wedding-game-api/models"
	"the-wedding-game-api/types"
)

func GetCurrentUserPoints(c *gin.Context) {
	user, err := middleware.GetCurrentUser(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	points, err := user.GetPoints()
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, types.CurrentUserPointsResponse{
		Points: points,
	})
	return
}

func GetLeaderboard(c *gin.Context) {
	_, err := middleware.GetCurrentUser(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	leaderboard, err := models.GetLeaderboard()
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, types.GetLeaderboardResponse{
		Leaderboard: leaderboard,
	})
	return
}
