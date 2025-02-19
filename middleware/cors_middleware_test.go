package middleware

import (
	"testing"
	test "the-wedding-game-api/_test"
)

func TestCORSMiddleware(t *testing.T) {
	request := test.GenerateBasicRequest()
	corsMiddleware := CORSMiddleware()
	corsMiddleware(request)

	if request.Writer.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Errorf("expected * but got %s", request.Writer.Header().Get("Access-Control-Allow-Origin"))
	}

	if request.Writer.Header().Get("Access-Control-Allow-Credentials") != "true" {
		t.Errorf("expected true but got %s", request.Writer.Header().Get("Access-Control-Allow-Credentials"))
	}

	if request.Writer.Header().Get("Access-Control-Allow-Headers") != "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With" {
		t.Errorf("expected Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With but got %s", request.Writer.Header().Get("Access-Control-Allow-Headers"))
	}

	if request.Writer.Header().Get("Access-Control-Allow-Methods") != "POST, OPTIONS, GET, PUT" {
		t.Errorf("expected POST, OPTIONS, GET, PUT but got %s", request.Writer.Header().Get("Access-Control-Allow-Methods"))
	}
}

func TestCORSMiddlewareAsOptions(t *testing.T) {
	request := test.GenerateBasicRequest()
	request.Request.Method = "OPTIONS"
	corsMiddleware := CORSMiddleware()
	corsMiddleware(request)

	if request.Writer.Status() != 204 {
		t.Errorf("expected 204 but got %d", request.Writer.Status())
	}
}
