package integrationtests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"the-wedding-game-api/models"
	"the-wedding-game-api/types"
)

var counter = 0

func createUserAndGetAccessToken() (models.User, models.AccessToken, error) {
	counter++
	user := models.User{
		Username: "test_user_for_challenges_" + strconv.Itoa(counter),
		Role:     types.Player,
	}
	user, err := user.Save()
	if err != nil {
		return models.User{}, models.AccessToken{}, err
	}

	accessToken, err := models.LinkAccessTokenToUser(user.ID)
	if err != nil {
		return models.User{}, models.AccessToken{}, err
	}

	return user, accessToken, nil
}

func createAdminAndGetAccessToken() (models.User, models.AccessToken, error) {
	counter++
	user := models.User{
		Username: "test_user_for_challenges_" + strconv.Itoa(counter),
		Role:     types.Admin,
	}
	user, err := user.Save()
	if err != nil {
		return models.User{}, models.AccessToken{}, err
	}

	accessToken, err := models.LinkAccessTokenToUser(user.ID)
	if err != nil {
		return models.User{}, models.AccessToken{}, err
	}

	return user, accessToken, nil
}

func deleteAllChallenges() {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),
	)

	database, err := gorm.Open(postgres.Open(dbURI))
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	database.Delete(&models.Challenge{}, "1 = 1")
}

func makeRequest(method string, path string, bodyObj interface{}) (int, string) {
	body, err := json.Marshal(bodyObj)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	return resp.Code, resp.Body.String()
}

func makeRequestWithFile(method string, path string, fileKey string, filePath string, accessToken string) (int, string) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	part, err := writer.CreateFormFile(fileKey, file.Name())
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

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+accessToken)
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	return resp.Code, resp.Body.String()
}

func makeRequestWithoutFile(method string, path string, accessToken string) (int, string) {
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "multipart/form-data")
	if accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+accessToken)
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	return resp.Code, resp.Body.String()
}

func makeRequestWithToken(method string, path string, bodyObj interface{}, accessToken string) (int, string) {
	body, err := json.Marshal(bodyObj)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	if accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+accessToken)
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	return resp.Code, resp.Body.String()
}

func createChallenge() (models.Challenge, error) {
	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      100,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}

	challenge, err := challenge.Save()
	if err != nil {
		return models.Challenge{}, err
	}

	return challenge, nil
}

func createAnswerQuestionChallenge() (models.Challenge, error) {
	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      100,
		Image:       "test_image",
		Type:        types.AnswerQuestionChallenge,
		Status:      types.ActiveChallenge,
	}

	challenge, err := challenge.Save()
	if err != nil {
		return models.Challenge{}, err
	}

	return challenge, nil
}

func createChallengeWithPoints(points uint) (models.Challenge, error) {
	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      points,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.ActiveChallenge,
	}

	challenge, err := challenge.Save()
	if err != nil {
		return models.Challenge{}, err
	}

	return challenge, nil
}

func completeChallenge(challengeID uint, userID uint) error {
	submission := models.Submission{
		ChallengeID: challengeID,
		UserID:      userID,
	}
	_, err := submission.Save()
	if err != nil {
		return err
	}

	return nil
}

func createSubmission(challengeID uint, userID uint, answer string) error {
	submission := models.Submission{
		ChallengeID: challengeID,
		UserID:      userID,
		Answer:      answer,
	}
	_, err := submission.Save()
	if err != nil {
		return err
	}

	return nil
}

func resetDatabase() error {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),
	)

	database, err := gorm.Open(postgres.Open(dbURI))
	if err != nil {
		return err
	}

	database.Exec("TRUNCATE TABLE users, access_tokens, challenges, submissions RESTART IDENTITY CASCADE")

	return nil
}
