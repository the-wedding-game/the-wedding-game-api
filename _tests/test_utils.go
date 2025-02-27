package test

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"strings"
	"the-wedding-game-api/db"
)

func SetupMockDb() *MockDB {
	mockDB := &MockDB{}
	db.GetConnection = func() db.DatabaseInterface {
		return mockDB
	}
	return mockDB
}

func GenerateBasicRequest() *gin.Context {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("POST", "/challenges", nil)
	return c
}

type BodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w BodyLogWriter) GetBody() string {
	return w.body.String()
}

func AttachBodyLogWriter(c *gin.Context) *BodyLogWriter {
	blw := &BodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	return blw
}

func IsUUID(uuid string) bool {
	if len(uuid) != 36 {
		return false
	}

	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		return false
	}

	return true
}

func GetFileExtension(fileName string) string {
	if !strings.Contains(fileName, ".") {
		return ""
	}
	return fileName[strings.LastIndex(fileName, "."):]
}
