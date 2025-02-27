package validators

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

func TestValidateUploadImageRequest(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	file, err := os.Open("../../_tests/assets/test_upload_image.jpg")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	part, err := writer.CreateFormFile("image", file.Name())
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}

	err = writer.Close()
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "/upload", body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)
	c.Request = req
	fileHeader, err := ValidateUploadImageRequest(c)
	if err != nil {
		t.Errorf("Expected error to be nil, got %v", err)
		return
	}

	if fileHeader.Filename != "test_upload_image.jpg" {
		t.Errorf("Expected filename to be test_upload_image.jpg, got %v", fileHeader.Filename)
	}
}

func TestValidateUploadImageRequestWithoutFile(t *testing.T) {
	req, err := http.NewRequest("POST", "/upload", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "multipart/form-data")

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)
	c.Request = req

	_, err = ValidateUploadImageRequest(c)
	if err == nil {
		t.Errorf("Expected error, got %v", err)
		return
	}

	expectedError := "image is required"
	if err.Error() != expectedError {
		t.Errorf("Expected error to be %v, got %v", expectedError, err)
	}
}

func TestValidateUploadImageRequestWithWrongKey(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	file, err := os.Open("../../_tests/assets/test_upload_image.jpg")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	part, err := writer.CreateFormFile("image2", file.Name())
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}

	err = writer.Close()
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "/upload", body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)
	c.Request = req

	_, err = ValidateUploadImageRequest(c)
	if err == nil {
		t.Errorf("Expected error, got %v", err)
		return
	}

	expectedError := "image is required"
	if err.Error() != expectedError {
		t.Errorf("Expected error to be %v, got %v", expectedError, err)
	}
}

func TestValidateUploadImageRequestWithInvalidFile(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	file, err := os.Open("../../_tests/assets/test_text.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	part, err := writer.CreateFormFile("image", file.Name())
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}

	err = writer.Close()
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "/upload", body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)
	c.Request = req

	_, err = ValidateUploadImageRequest(c)
	if err == nil {
		t.Errorf("Expected error, got %v", err)
		return
	}

	expectedError := "file must be an image"
	if err.Error() != expectedError {
		t.Errorf("Expected error to be %v, got %v", expectedError, err)
	}
}

func TestValidateUploadImageRequestWithEmptyFile(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	file, err := os.Open("../../_tests/assets/test_empty_image.jpg")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	part, err := writer.CreateFormFile("image", file.Name())
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}

	err = writer.Close()
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "/upload", body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)
	c.Request = req

	_, err = ValidateUploadImageRequest(c)
	if err == nil {
		t.Errorf("Expected error, got %v", err)
		return
	}

	expectedError := "file is empty"
	if err.Error() != expectedError {
		t.Errorf("Expected error to be %v, got %v", expectedError, err)
	}
}

func TestValidateUploadImageRequestWithLargeFile(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	file, err := os.Open("../../_tests/assets/test_large_image.jpg")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	part, err := writer.CreateFormFile("image", file.Name())
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}

	err = writer.Close()
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "/upload", body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(nil)
	c.Request = req

	_, err = ValidateUploadImageRequest(c)
	if err == nil {
		t.Errorf("Expected error, got %v", err)
		return
	}

	expectedError := "maximum file size is 1048576 bytes"
	if err.Error() != expectedError {
		t.Errorf("Expected error to be %v, got %v", expectedError, err)
	}
}
