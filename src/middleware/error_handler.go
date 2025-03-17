package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	apperrors "the-wedding-game-api/errors"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	err := c.Errors.Last()

	if err != nil {
		if apperrors.IsAccessTokenNotFoundError(err.Err) || apperrors.IsAuthorizationError(err.Err) {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  "error",
				"message": "access denied",
			})
			c.Abort()
			return
		}

		if apperrors.IsAuthenticationError(err.Err) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": err.Err.Error(),
			})
			c.Abort()
			return
		}

		if apperrors.IsValidationError(err.Err) || strings.Contains(err.Err.Error(), "Error:Field validation for") {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": err.Err.Error(),
			})
			c.Abort()
			return
		}

		if apperrors.IsNotFoundError(err.Err) {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": err.Err.Error(),
			})
			c.Abort()
			return
		}

		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "An unexpected error occurred.",
		})
		c.Abort()
	}
}
