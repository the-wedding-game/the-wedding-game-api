package utils

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"strings"
	"the-wedding-game-api/storage"
)

func generateRandomFileName(originalFileName string) (string, error) {
	fileNameUUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	fileExtension := strings.Split(originalFileName, ".")[1]
	fileName := fmt.Sprintf("%s.%s", fileNameUUID.String(), fileExtension)

	return fileName, nil
}

func getReader(fileBytes multipart.File) (*bytes.Reader, error) {
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, fileBytes); err != nil {
		return nil, err
	}

	return bytes.NewReader(buf.Bytes()), nil
}

func UploadFile(file *multipart.FileHeader) (string, error) {
	fileBytes, err := file.Open()
	if err != nil {
		return "", err
	}
	defer func(fileBytes multipart.File) {
		err := fileBytes.Close()
		if err != nil {
			return
		}
	}(fileBytes)

	reader, err := getReader(fileBytes)
	if err != nil {
		return "", err
	}

	storageService, err := storage.GetStorage()
	if err != nil {
		return "", err
	}

	generatedFileName, err := generateRandomFileName(file.Filename)
	if err != nil {
		return "", err
	}

	url, err := storageService.UploadFile(*reader, generatedFileName)
	if err != nil {
		return "", err
	}

	return url, nil
}
