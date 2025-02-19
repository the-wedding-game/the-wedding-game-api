package test

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
)

func GenerateBasicRequest() *gin.Context {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("POST", "/challenges", nil)
	return c
}
