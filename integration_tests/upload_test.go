package integrationtests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
	"the-wedding-game-api/storage"
	"the-wedding-game-api/types"
	"the-wedding-game-api/utils"
)

func TestUploadImage(t *testing.T) {
	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating admin and getting access token")
		return
	}

	statusCode, body := makeRequestWithFile("POST", "/upload", "image", "../_tests/assets/test_upload_image.jpg", accessToken.Token)
	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
		return
	}

	var response types.UploadResponse
	decoder := json.NewDecoder(bytes.NewReader([]byte(body)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&response); err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
		return
	}

	if response.Url == "" {
		t.Errorf("Expected URL to not be empty")
		return
	}

	if !utils.IsURLStrict(response.Url) {
		t.Errorf("Expected URL to be valid")
		return
	}

	region := os.Getenv("AWS_REGION")
	bucketName := os.Getenv("AWS_BUCKET_NAME")
	folderName := os.Getenv("AWS_FOLDER_NAME")

	if strings.Contains(response.Url, "https://"+storage.RemoveLeadingSlash(bucketName)+".s3."+region+".amazonaws.com/"+folderName+"/") == false {
		t.Errorf("Invalid URL format:" + response.Url)
		return
	}
}

func TestUploadImageNotAdmin(t *testing.T) {
	_, accessToken, err := createUserAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating user and getting access token")
		return
	}

	statusCode, body := makeRequestWithFile("POST", "/upload", "image", "../_tests/assets/test_upload_image.jpg", accessToken.Token)
	if statusCode != http.StatusForbidden {
		t.Errorf("Expected status code 403, got %v", statusCode)
		return
	}

	expectedBody := `{"message":"access denied","status":"error"}`
	if body != expectedBody {
		t.Errorf("Expected body to be %v, got %v", expectedBody, body)
	}
}

func TestUploadImageNoToken(t *testing.T) {
	statusCode, body := makeRequestWithFile("POST", "/upload", "image", "../_tests/assets/test_upload_image.jpg", "")
	if statusCode != http.StatusForbidden {
		t.Errorf("Expected status code 403, got %v", statusCode)
		return
	}

	expectedBody := `{"message":"access denied","status":"error"}`
	if body != expectedBody {
		t.Errorf("Expected body to be %v, got %v", expectedBody, body)
	}
}

func TestUploadImageWithoutFormFile(t *testing.T) {
	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating admin and getting access token")
		return
	}

	statusCode, body := makeRequestWithoutFile("POST", "/upload", accessToken.Token)
	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	expectedBody := `{"message":"image is required","status":"error"}`
	if body != expectedBody {
		t.Errorf("Expected body to be %v, got %v", expectedBody, body)
	}

	fmt.Println(body)
}

func TestUploadImageWithInvalidFormFile(t *testing.T) {
	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating admin and getting access token")
		return
	}

	statusCode, body := makeRequestWithFile("POST", "/upload", "invalid", "../_tests/assets/test_upload_image.jpg", accessToken.Token)
	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	expectedBody := `{"message":"image is required","status":"error"}`
	if body != expectedBody {
		t.Errorf("Expected body to be %v, got %v", expectedBody, body)
	}
}

func TestUploadImageWithEmptyFile(t *testing.T) {
	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating admin and getting access token")
		return
	}

	statusCode, body := makeRequestWithFile("POST", "/upload", "image", "../_tests/assets/test_empty_image.jpg", accessToken.Token)
	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	expectedBody := `{"message":"file is empty","status":"error"}`
	if body != expectedBody {
		t.Errorf("Expected body to be %v, got %v", expectedBody, body)
	}
}

func TestUploadWithNonImageFile(t *testing.T) {
	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating admin and getting access token")
		return
	}

	statusCode, body := makeRequestWithFile("POST", "/upload", "image", "../_tests/assets/test_text.txt", accessToken.Token)
	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	expectedBody := `{"message":"file must be an image","status":"error"}`
	if body != expectedBody {
		t.Errorf("Expected body to be %v, got %v", expectedBody, body)
	}
}

func TestUploadImageWithInvalidToken(t *testing.T) {
	statusCode, body := makeRequestWithFile("POST", "/upload", "image", "../_tests/assets/test_upload_image.jpg", "invalid")

	fmt.Println(body)
	if statusCode != http.StatusForbidden {
		t.Errorf("Expected status code 403, got %v", statusCode)
		return
	}

	expectedBody := `{"message":"access denied","status":"error"}`
	if body != expectedBody {
		t.Errorf("Expected body to be %v, got %v", expectedBody, body)
	}
}

func TestUploadImageWithTooLargeFile(t *testing.T) {
	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating admin and getting access token")
		return
	}

	statusCode, body := makeRequestWithFile("POST", "/upload", "image", "../_tests/assets/test_large_image.jpg", accessToken.Token)
	if statusCode != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %v", statusCode)
	}

	expectedBody := `{"message":"maximum file size is 1048576 bytes","status":"error"}`
	if body != expectedBody {
		t.Errorf("Expected body to be %v, got %v", expectedBody, body)
	}
}
