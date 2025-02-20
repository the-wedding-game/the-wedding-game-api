package apperrors

import "errors"

type DatabaseError struct {
	code    string
	Message string
}

func (e DatabaseError) Error() string {
	return e.Message
}

func NewDatabaseError(message string) DatabaseError {
	return DatabaseError{
		code:    "DatabaseError",
		Message: message,
	}
}

func IsDatabaseError(err error) bool {
	var databaseError DatabaseError
	ok := errors.As(err, &databaseError)
	return ok
}
