package test

import (
	"bytes"
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
