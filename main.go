package main

import (
	"github.com/gin-gonic/gin"
	"the-wedding-game-api/middleware"
	"the-wedding-game-api/routes"
)

func main() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.ErrorHandler)

	router.GET("/challenges/:id", routes.GetChallengeById)
	router.POST("/challenges", routes.CreateChallenge)
	router.GET("/challenges", routes.GetAllChallenges)
	router.POST("/challenges/:id/verify", routes.VerifyAnswer)

	router.POST("/auth/login", routes.Login)
	router.GET("/auth/current-user", routes.GetCurrentUser)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
