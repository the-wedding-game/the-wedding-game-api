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

	fileExtension := ""
	if strings.Contains(originalFileName, ".") {
		fileExtension = originalFileName[strings.LastIndex(originalFileName, ".")-1:]
	}
	fileName := fmt.Sprintf("%s%s", fileNameUUID.String(), fileExtension)

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
		return "", fmt.Errorf("error while opening file: %w", err)
	}
	defer func(fileBytes multipart.File) {
		err := fileBytes.Close()
		if err != nil {
			return
		}
	}(fileBytes)

	reader, err := getReader(fileBytes)
	if err != nil {
		return "", fmt.Errorf("error while getting reader: %w", err)
	}

	storageService, err := storage.GetStorage()
	if err != nil {
		return "", fmt.Errorf("error while getting storage service: %w", err)
	}

	generatedFileName, err := generateRandomFileName(file.Filename)
	if err != nil {
		return "", fmt.Errorf("error while generating random file name: %w", err)
	}

	url, err := storageService.UploadFile(*reader, generatedFileName)
	if err != nil {
		return "", fmt.Errorf("error while uploading file: %w", err)
	}

	return url, nil
}
