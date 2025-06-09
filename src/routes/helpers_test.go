package routes

import (
	"bytes"
	"database/sql"
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
	database, err := getDatabaseConnection()
	defer closeDatabaseConnection(database)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	db, _ := database.DB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
	}(db)

	database.Exec(`DELETE FROM answers WHERE id > 0`)
	database.Exec(`DELETE FROM submissions WHERE id > 0`)
	database.Exec(`DELETE FROM challenges WHERE id > 0`)

	if database.Error != nil {
		log.Fatalf("Error deleting challenges: %v", database.Error)
	}
}

func dropSubmissionsTable() {
	database, err := getDatabaseConnection()
	defer closeDatabaseConnection(database)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	db, _ := database.DB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
	}(db)

	database.Exec(`DROP TABLE IF EXISTS submissions`)

	if database.Error != nil {
		log.Fatalf("Error dropping submissions table: %v", database.Error)
	}
}

func dropAnswersTable() {
	database, err := getDatabaseConnection()
	defer closeDatabaseConnection(database)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	db, _ := database.DB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
	}(db)

	database.Exec(`DROP TABLE IF EXISTS answers`)

	if database.Error != nil {
		log.Fatalf("Error dropping answers table: %v", database.Error)
	}
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

func createInactiveChallengeWithPoints(points uint) (models.Challenge, error) {
	challenge := models.Challenge{
		Name:        "test_challenge",
		Description: "test_description",
		Points:      points,
		Image:       "test_image",
		Type:        types.UploadPhotoChallenge,
		Status:      types.InactiveChallenge,
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

func getDatabaseConnection() (db *gorm.DB, err error) {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),
	)

	return gorm.Open(postgres.Open(dbURI))
}

func closeDatabaseConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting database connection: %v", err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("Error closing database connection: %v", err)
	}
}

func resetDatabase() error {
	database, err := getDatabaseConnection()
	if err != nil {
		return err
	}
	defer closeDatabaseConnection(database)

	database.Exec("TRUNCATE TABLE users, access_tokens, challenges, submissions RESTART IDENTITY CASCADE")

	return nil
}

func checkIfAnswerExists(answerId uint) (bool, error) {
	database, err := getDatabaseConnection()
	if err != nil {
		return false, err
	}
	defer closeDatabaseConnection(database)

	var count int64
	err = database.Model(&models.Answer{}).Where("id = ?", answerId).Count(&count).Error
	if err != nil {
		log.Fatalf("Error checking if answer exists: %v", err)
	}

	return count > 0, nil
}

func checkIfAnswerExistsForChallenge(challengeId uint) (bool, error) {
	database, err := getDatabaseConnection()
	if err != nil {
		return false, err
	}
	defer closeDatabaseConnection(database)

	var count int64
	err = database.Model(&models.Answer{}).Where("challenge_id = ?", challengeId).Count(&count).Error
	if err != nil {
		log.Fatalf("Error checking if answer exists for challenge: %v", err)
	}

	return count > 0, nil
}

func checkIfSubmissionExists(submissionId uint) (bool, error) {
	database, err := getDatabaseConnection()
	if err != nil {
		return false, err
	}
	defer closeDatabaseConnection(database)

	var count int64
	err = database.Model(&models.Submission{}).Where("id = ?", submissionId).Count(&count).Error
	if err != nil {
		log.Fatalf("Error checking if submission exists: %v", err)
	}

	return count > 0, nil
}

func checkIfSubmissionExistsForChallenge(challengeId uint) (bool, error) {
	database, err := getDatabaseConnection()
	if err != nil {
		return false, err
	}
	defer closeDatabaseConnection(database)

	var count int64
	err = database.Model(&models.Submission{}).Where("challenge_id = ?", challengeId).Count(&count).Error
	if err != nil {
		log.Fatalf("Error checking if submission exists for challenge: %v", err)
	}

	return count > 0, nil
}
