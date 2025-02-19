package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
	"the-wedding-game-api/db"
)

func generateBasicRequest() *gin.Context {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("POST", "/challenges", nil)
	return c

}

func TestGetCurrentUser(t *testing.T) {
	mockDB := &db.MockDB{}
	db.GetConnection = func() db.DatabaseInterface {
		fmt.Println("Returning mock db")
		return mockDB
	}

	request := generateBasicRequest()
	request.Request.Header.Set("Authorization", "Bearer token")
	user, err := GetCurrentUser(request)
	if err != nil {
		t.Errorf("expected nil but got %s", err.Error())
	}
	fmt.Println(user)
}
