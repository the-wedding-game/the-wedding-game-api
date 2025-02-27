package test

import (
	"bytes"
	apperrors "the-wedding-game-api/errors"
)

type MockStorage struct {
	err error
}

func (m *MockStorage) UploadFile(_ bytes.Reader, fileName string) (string, error) {
	if m.err != nil {
		err := m.err
		m.err = nil
		return "", err
	}
	return "https://example.com/" + fileName, nil
}

func (m *MockStorage) SetError(err string) {
	m.err = apperrors.NewStorageError(err)
}
