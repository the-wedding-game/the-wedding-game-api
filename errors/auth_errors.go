package apperrors

import "errors"

type AuthenticationError struct {
	code    string
	Message string
}

func (e AuthenticationError) Error() string {
	return e.Message
}

func NewAuthenticationError(message string) AuthenticationError {
	return AuthenticationError{
		code:    "AuthenticationError",
		Message: message,
	}
}

func IsAuthenticationError(err error) bool {
	var authenticationError AuthenticationError
	ok := errors.As(err, &authenticationError)
	return ok
}

type AuthorizationError struct {
	code    string
	Message string
}

func (e AuthorizationError) Error() string {
	return e.Message
}

func NewAuthorizationError() AuthorizationError {
	return AuthorizationError{
		code:    "AuthorizationError",
		Message: "access denied",
	}
}

func IsAuthorizationError(err error) bool {
	var authorizationError AuthorizationError
	ok := errors.As(err, &authorizationError)
	return ok
}

type AccessTokenNotFoundError struct {
	code    string
	Message string
}

func (e AccessTokenNotFoundError) Error() string {
	return e.Message
}

func NewAccessTokenNotFoundError() AccessTokenNotFoundError {
	return AccessTokenNotFoundError{
		code:    "AccessTokenNotFoundError",
		Message: "access token not found",
	}
}

func IsAccessTokenNotFoundError(err error) bool {
	var accessTokenNotFoundError AccessTokenNotFoundError
	ok := errors.As(err, &accessTokenNotFoundError)
	return ok
}
