package apperrors

type StorageError struct {
	code    string
	Message string
}

func (e StorageError) Error() string {
	return e.Message
}

func NewStorageError(message string) StorageError {
	return StorageError{
		code:    "StorageError",
		Message: message,
	}
}
