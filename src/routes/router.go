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

	router.GET("/challenges/:id", GetChallengeById)
	router.POST("/challenges", CreateChallenge)
	router.GET("/challenges", GetAllChallenges)
	router.POST("/challenges/:id/verify", VerifyAnswer)

	router.POST("/auth/login", Login)
	router.GET("/auth/current-user", GetCurrentUser)

	router.GET("/points/me", GetCurrentUserPoints)

	router.POST("/upload", HandleImageUpload)

	return router
}
