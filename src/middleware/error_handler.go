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
		handleError(c, err.Err)
		return
	}
}

func handleError(c *gin.Context, err error) {
	if apperrors.IsAccessTokenNotFoundError(err) || apperrors.IsAuthorizationError(err) {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  "error",
			"message": "access denied",
		})
		c.Abort()
		return
	}

	if apperrors.IsAuthenticationError(err) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	if apperrors.IsValidationError(err) || strings.Contains(err.Error(), "Error:Field validation for") {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	if apperrors.IsNotFoundError(err) || apperrors.IsRecordNotFoundError(err) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": err.Error(),
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
