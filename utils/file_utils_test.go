package utils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"testing"
	test "the-wedding-game-api/_tests"
)

func getFileHeader(filepath string) (*multipart.FileHeader, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error while opening file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("error while closing file: %w", err))
		}
	}(file)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	formFile, err := writer.CreateFormFile("image", file.Name())
	if err != nil {
		return nil, fmt.Errorf("error while creating form file: %w", err)
	}

	_, err = io.Copy(formFile, file)
	if err != nil {
		return nil, fmt.Errorf("error while copying file: %w", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("error while closing writer: %w", err)
	}

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	_, fileHeader, err := req.FormFile("image")
	if err != nil {
		return nil, fmt.Errorf("error while getting form file: %w", err)
	}

	return fileHeader, nil
}

func TestGenerateRandomFileName1(t *testing.T) {
	generatedFileName, err := generateRandomFileName("test.txt")
	if err != nil {
		t.Errorf("Error while generating random file name")
		return
	}

	fmt.Println(generatedFileName)

	if !test.IsUUID(generatedFileName[:36]) {
		t.Errorf("Incorrect UUID")
		return
	}

	if test.GetFileExtension(generatedFileName) != ".txt" {
		t.Errorf("Incorrect file extension")
		return
	}
}

func TestGenerateRandomFileName2(t *testing.T) {
	generatedFileName, err := generateRandomFileName("random.jpg")
	if err != nil {
		t.Errorf("Error while generating random file name")
		return
	}

	fmt.Println(generatedFileName)

	if !test.IsUUID(generatedFileName[:36]) {
		t.Errorf("Incorrect UUID")
		return
	}

	if test.GetFileExtension(generatedFileName) != ".jpg" {
		t.Errorf("Incorrect file extension")
		return
	}
}

func TestGenerateRandomFileName3(t *testing.T) {
	generatedFileName, err := generateRandomFileName("abcd.efgh")
	if err != nil {
		t.Errorf("Error while generating random file name")
		return
	}

	fmt.Println(generatedFileName)

	if !test.IsUUID(generatedFileName[:36]) {
		t.Errorf("Incorrect UUID")
		return
	}

	if test.GetFileExtension(generatedFileName) != ".efgh" {
		t.Errorf("Incorrect file extension")
		return
	}
}

func TestGenerateRandomFileName4(t *testing.T) {
	generatedFileName, err := generateRandomFileName("abcd.ef")
	if err != nil {
		t.Errorf("Error while generating random file name")
		return
	}

	fmt.Println(generatedFileName)

	if !test.IsUUID(generatedFileName[:36]) {
		t.Errorf("Incorrect UUID")
		return
	}

	if test.GetFileExtension(generatedFileName) != ".ef" {
		t.Errorf("Incorrect file extension")
		return
	}
}

func TestGenerateRandomFileName5(t *testing.T) {
	generatedFileName, err := generateRandomFileName("abcd.def.ghi")
	if err != nil {
		t.Errorf("Error while generating random file name")
		return
	}

	fmt.Println(generatedFileName)

	if !test.IsUUID(generatedFileName[:36]) {
		t.Errorf("Incorrect UUID")
		return
	}

	if test.GetFileExtension(generatedFileName) != ".ghi" {
		t.Errorf("Incorrect file extension")
		return
	}
}

func TestGenerateRandomFileNameWithNoExtension1(t *testing.T) {
	generatedFileName, err := generateRandomFileName("abcd")
	if err != nil {
		t.Errorf("Error while generating random file name")
		return
	}

	fmt.Println(generatedFileName)

	if !test.IsUUID(generatedFileName[:36]) {
		t.Errorf("Incorrect UUID")
		return
	}

	if test.GetFileExtension(generatedFileName) != "" {
		t.Errorf("Incorrect file extension")
		return
	}
}

func TestGenerateRandomFileNameWithNoExtension2(t *testing.T) {
	generatedFileName, err := generateRandomFileName("abcd.")
	if err != nil {
		t.Errorf("Error while generating random file name")
		return
	}

	if !test.IsUUID(generatedFileName[:36]) {
		t.Errorf("Incorrect UUID")
		return
	}

	if test.GetFileExtension(generatedFileName) != "." {
		t.Errorf("Incorrect file extension")
		return
	}
}

func TestGetReader(t *testing.T) {
	file, err := os.Open("../_tests/assets/test.txt")
	if err != nil {
		t.Errorf("Error while opening file")
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("error while closing file: %w", err))
		}
	}(file)

	reader, err := getReader(file)
	if err != nil {
		t.Errorf("Error while getting reader")
		return
	}

	if reader == nil {
		t.Errorf("Reader is nil")
		return
	}

	content := make([]byte, 13)
	_, err = reader.Read(content)
	if err != nil {
		t.Errorf("Error while reading content")
		return
	}

	fmt.Println(string(content))

	expected := "Hello, World!"
	if string(content) != expected {
		t.Errorf("Incorrect content")
		return
	}
}

func TestUploadFile(t *testing.T) {
	test.SetupMockStorage()

	fileHeader, err := getFileHeader("../_tests/assets/test.txt")
	if err != nil {
		t.Errorf(fmt.Errorf("error while getting file header: %w", err).Error())
		return
	}

	url, err := UploadFile(fileHeader)
	if err != nil {
		t.Errorf(fmt.Errorf("error while uploading file: %w", err).Error())
		return
	}

	fmt.Println(url)
	fmt.Println(url[:20])

	if !IsURLStrict(url) {
		t.Errorf("invalid URL")
		return
	}

	if url[:20] != "https://example.com/" {
		t.Errorf("Incorrect prefix")
		return
	}

	if !test.IsUUID(url[20:56]) {
		t.Errorf("incorrect UUID")
		return
	}

	if test.GetFileExtension(url) != ".txt" {
		t.Errorf("Incorrect file extension")
	}
}

func TestUploadFileStorageError(t *testing.T) {
	mockStorage := test.SetupMockStorage()
	mockStorage.SetError("mocked error while uploading file")

	fileHeader, err := getFileHeader("../_tests/assets/test.txt")
	if err != nil {
		t.Errorf(fmt.Errorf("error while getting file header: %w", err).Error())
		return
	}

	_, err = UploadFile(fileHeader)
	if err == nil {
		t.Errorf("Expected error")
		return
	}
}
