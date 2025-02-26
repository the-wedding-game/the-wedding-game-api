package middleware

import (
	"errors"
	"net/http"
	"testing"
	test "the-wedding-game-api/_tests"
	apperrors "the-wedding-game-api/errors"
)

func TestErrorHandlerWithNoError(t *testing.T) {
	request := test.GenerateBasicRequest()
	ErrorHandler(request)
}

func TestErrorHandlerWithAccessTokenNotFoundError(t *testing.T) {
	request := test.GenerateBasicRequest()
	blw := test.AttachBodyLogWriter(request)
	_ = request.Error(apperrors.NewAccessTokenNotFoundError())
	ErrorHandler(request)

	if request.Writer.Status() != http.StatusForbidden {
		t.Errorf("expected 403 but got %d", request.Writer.Status())
	}

	expectedBody := "{\"message\":\"access denied\",\"status\":\"error\"}"
	if blw.GetBody() != expectedBody {
		t.Errorf("expected %s but got %s", expectedBody, blw.GetBody())
	}
}

func TestErrorHandlerWithAuthenticationError(t *testing.T) {
	request := test.GenerateBasicRequest()
	blw := test.AttachBodyLogWriter(request)
	_ = request.Error(apperrors.NewAuthenticationError("test_error"))
	ErrorHandler(request)

	if request.Writer.Status() != http.StatusForbidden {
		t.Errorf("expected 403 but got %d", request.Writer.Status())
	}

	expectedBody := "{\"message\":\"access denied\",\"status\":\"error\"}"
	if blw.GetBody() != expectedBody {
		t.Errorf("expected %s but got %s", expectedBody, blw.GetBody())
	}
}

func TestErrorHandlerWithAuthorizationError(t *testing.T) {
	request := test.GenerateBasicRequest()
	blw := test.AttachBodyLogWriter(request)
	_ = request.Error(apperrors.NewAuthorizationError())
	ErrorHandler(request)

	if request.Writer.Status() != http.StatusForbidden {
		t.Errorf("expected 403 but got %d", request.Writer.Status())
	}

	expectedBody := "{\"message\":\"access denied\",\"status\":\"error\"}"
	if blw.GetBody() != expectedBody {
		t.Errorf("expected %s but got %s", expectedBody, blw.GetBody())
	}
}

func TestErrorHandlerWithValidationError(t *testing.T) {
	request := test.GenerateBasicRequest()
	blw := test.AttachBodyLogWriter(request)
	_ = request.Error(apperrors.NewValidationError("test_error"))
	ErrorHandler(request)

	if request.Writer.Status() != http.StatusBadRequest {
		t.Errorf("expected 400 but got %d", request.Writer.Status())
	}

	expectedBody := "{\"message\":\"test_error\",\"status\":\"error\"}"
	if blw.GetBody() != expectedBody {
		t.Errorf("expected %s but got %s", expectedBody, blw.GetBody())
	}
}

func TestErrorHandlerWithNotFoundError(t *testing.T) {
	request := test.GenerateBasicRequest()
	blw := test.AttachBodyLogWriter(request)
	_ = request.Error(apperrors.NewNotFoundError("test_entity", "test_key"))
	ErrorHandler(request)

	if request.Writer.Status() != http.StatusNotFound {
		t.Errorf("expected 404 but got %d", request.Writer.Status())
	}

	expectedBody := "{\"message\":\"test_entity with key test_key not found.\",\"status\":\"error\"}"
	if blw.GetBody() != expectedBody {
		t.Errorf("expected %s but got %s", expectedBody, blw.GetBody())
	}
}

func TestErrorHandlerWithUnexpectedError(t *testing.T) {
	request := test.GenerateBasicRequest()
	blw := test.AttachBodyLogWriter(request)
	_ = request.Error(errors.New("unexpected error"))
	ErrorHandler(request)

	if request.Writer.Status() != http.StatusInternalServerError {
		t.Errorf("expected 500 but got %d", request.Writer.Status())
	}

	expectedBody := "{\"message\":\"An unexpected error occurred.\",\"status\":\"error\"}"
	if blw.GetBody() != expectedBody {
		t.Errorf("expected %s but got %s", expectedBody, blw.GetBody())
	}
}
