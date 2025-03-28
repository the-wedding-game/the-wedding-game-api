package routes

import (
	"github.com/gin-gonic/gin"
	"the-wedding-game-api/middleware"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.ErrorHandler)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello!",
		})
	})

	router.GET("/challenges/:id", middleware.IsLoggedIn, GetChallengeById)
	router.POST("/challenges", middleware.IsAdmin, CreateChallenge)
	router.GET("/challenges", middleware.IsLoggedIn, GetAllChallenges)
	router.POST("/challenges/:id/verify", middleware.IsLoggedIn, VerifyAnswer)
	router.GET("/challenges/:id/submissions", middleware.IsLoggedIn, GetSubmissions)
	router.PUT("/challenges/:id", middleware.IsAdmin, UpdateChallenge)
	router.GET("/challenges/:id/answer", middleware.IsAdmin, GetAnswer)

	router.POST("/auth/login", Login)
	router.GET("/auth/current-user", GetCurrentUser)

	router.GET("/points/me", middleware.IsLoggedIn, GetCurrentUserPoints)
	router.GET("/leaderboard", middleware.IsLoggedIn, GetLeaderboard)

	router.GET("/gallery", middleware.IsLoggedIn, GetGallery)

	router.POST("/upload", middleware.IsLoggedIn, HandleImageUpload)

	router.GET("/admin/challenges", middleware.IsAdmin, GetAllChallengesAdmin)

	return router
}
