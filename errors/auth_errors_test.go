package apperrors

import "testing"

func TestCreateAuthenticationError(t *testing.T) {
	authenticationError := NewAuthenticationError("authentication failed")
	if authenticationError.Message != "authentication failed" {
		t.Errorf("expected authentication failed but got %s", authenticationError.Message)
	}
	if authenticationError.code != "AuthenticationError" {
		t.Errorf("expected AuthenticationError but got %s", authenticationError.code)
	}
	if !IsAuthenticationError(authenticationError) {
		t.Errorf("expected true but got false")
	}

	authenticationError = NewAuthenticationError("hello there")
	if authenticationError.Message != "hello there" {
		t.Errorf("expected hello there but got %s", authenticationError.Message)
	}

	if authenticationError.code != "AuthenticationError" {
		t.Errorf("expected AuthenticationError but got %s", authenticationError.code)
	}

	if !IsAuthenticationError(authenticationError) {
		t.Errorf("expected true but got false")
	}

}

func TestAuthenticationErrorMessage(t *testing.T) {
	authenticationError := NewAuthenticationError("hello there")
	if authenticationError.Error() != "hello there" {
		t.Errorf("expected hello there but got %s", authenticationError.Error())
	}

	authenticationError = NewAuthenticationError("random message")
	if authenticationError.Error() != "random message" {
		t.Errorf("expected authentication failed but got %s", authenticationError.Error())
	}
}

func TestIsAuthenticationError(t *testing.T) {
	authenticationError := NewAuthenticationError("hello there")
	if !IsAuthenticationError(authenticationError) {
		t.Errorf("expected true but got false")
	}

	authenticationError = NewAuthenticationError("random message")
	if !IsAuthenticationError(authenticationError) {
		t.Errorf("expected true but got false")
	}
}

func TestIsAuthenticationErrorNegative(t *testing.T) {
	authorizationError := NewAuthorizationError()
	if IsAuthenticationError(authorizationError) {
		t.Errorf("expected false but got true")
	}

	accessTokenNotFoundError := NewAccessTokenNotFoundError()
	if IsAuthenticationError(accessTokenNotFoundError) {
		t.Errorf("expected false but got true")
	}
}

func TestCreateAuthorizationError(t *testing.T) {
	authorizationError := NewAuthorizationError()
	if authorizationError.Message != "access denied" {
		t.Errorf("expected access denied but got %s", authorizationError.Message)
	}
	if authorizationError.code != "AuthorizationError" {
		t.Errorf("expected AuthorizationError but got %s", authorizationError.code)
	}
	if !IsAuthorizationError(authorizationError) {
		t.Errorf("expected true but got false")
	}
}

func TestAuthorizationErrorMessage(t *testing.T) {
	authorizationError := NewAuthorizationError()
	if authorizationError.Error() != "access denied" {
		t.Errorf("expected access denied but got %s", authorizationError.Error())
	}
}

func TestIsAuthorizationError(t *testing.T) {
	authorizationError := NewAuthorizationError()
	if !IsAuthorizationError(authorizationError) {
		t.Errorf("expected true but got false")
	}
}

func TestIsAuthorizationErrorNegative(t *testing.T) {
	authenticationError := NewAuthenticationError("hello there")
	if IsAuthorizationError(authenticationError) {
		t.Errorf("expected false but got true")
	}

	accessTokenNotFoundError := NewAccessTokenNotFoundError()
	if IsAuthorizationError(accessTokenNotFoundError) {
		t.Errorf("expected false but got true")
	}
}

func TestCreateAccessTokenNotFoundError(t *testing.T) {
	accessTokenNotFoundError := NewAccessTokenNotFoundError()
	if accessTokenNotFoundError.Message != "access token not found" {
		t.Errorf("expected access token not found but got %s", accessTokenNotFoundError.Message)
	}
	if accessTokenNotFoundError.code != "AccessTokenNotFoundError" {
		t.Errorf("expected AccessTokenNotFoundError but got %s", accessTokenNotFoundError.code)
	}
	if !IsAccessTokenNotFoundError(accessTokenNotFoundError) {
		t.Errorf("expected true but got false")
	}
}

func TestAccessTokenNotFoundErrorMessage(t *testing.T) {
	accessTokenNotFoundError := NewAccessTokenNotFoundError()
	if accessTokenNotFoundError.Error() != "access token not found" {
		t.Errorf("expected access token not found but got %s", accessTokenNotFoundError.Error())
	}
}

func TestIsAccessTokenNotFoundError(t *testing.T) {
	accessTokenNotFoundError := NewAccessTokenNotFoundError()
	if !IsAccessTokenNotFoundError(accessTokenNotFoundError) {
		t.Errorf("expected true but got false")
	}
}

func TestIsAccessTokenNotFoundErrorNegative(t *testing.T) {
	authenticationError := NewAuthenticationError("hello there")
	if IsAccessTokenNotFoundError(authenticationError) {
		t.Errorf("expected false but got true")
	}

	authorizationError := NewAuthorizationError()
	if IsAccessTokenNotFoundError(authorizationError) {
		t.Errorf("expected false but got true")
	}
}
