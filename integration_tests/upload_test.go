package integrationtests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
	"the-wedding-game-api/types"
	"the-wedding-game-api/utils"
)

func TestUploadImage(t *testing.T) {
	_, accessToken, err := createAdminAndGetAccessToken()
	if err != nil {
		t.Errorf("Error creating admin and getting access token")
		return
	}

	statusCode, body := makeRequestWithFile("POST", "/upload", "image", "assets/test_upload_image.jpg", accessToken.Token)
	if statusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", statusCode)
		return
	}

	fmt.Println(body)

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

	if strings.Contains(response.Url, "https://"+bucketName+".s3."+region+".amazonaws.com/"+folderName+"/") == false {
		t.Errorf("Invalid URL format")
		return
	}
}
