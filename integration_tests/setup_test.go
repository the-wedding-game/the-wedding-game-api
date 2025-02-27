package integrationtests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"
	"the-wedding-game-api/models"
	"the-wedding-game-api/routes"
	"time"
)

func Setup() {
	dockerComposeUp()
	setupTestDb()
	waitForS3()
}

func TearDown() {
	//dockerComposeDown()
}

func dockerComposeUp() {
	dockerComposeDown()
	log.Println("Starting Docker Compose...")

	cmdUp := exec.Command("docker-compose", "-f", "../_tests/docker-compose.yml", "up", "-d")
	if err := cmdUp.Run(); err != nil {
		log.Fatalf("Failed to start Docker Compose: %v", err)
	}
}

func dockerComposeDown() {
	log.Println("Stopping Docker Compose...")

	cmdDown := exec.Command("docker-compose", "-f", "../_tests/docker-compose.yml", "down")
	if err := cmdDown.Run(); err != nil {
		log.Fatalf("Failed to stop Docker Compose: %v", err)
	}
}

func setupTestDb() {
	_ = os.Setenv("DB_HOST", "localhost")
	_ = os.Setenv("DB_PORT", "5433")
	_ = os.Setenv("DB_USER", "the-wedding-game-api")
	_ = os.Setenv("DB_NAME", "the-wedding-game")
	_ = os.Setenv("DB_PASS", "abcd@123")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),
	)

	ready := false
	for i := 0; i < 10; i++ {
		db, err := gorm.Open(postgres.Open(dbURI))
		if err == nil {
			ready = true
			log.Println("Database is ready!")
			err := db.Migrator().DropTable(&models.User{}, &models.AccessToken{}, &models.Challenge{}, &models.Answer{}, &models.Submission{})
			if err != nil {
				panic(err)
			}

			log.Println("Migrating schema...")
			err = db.Debug().AutoMigrate(&models.User{}, &models.AccessToken{}, &models.Challenge{}, &models.Answer{}, &models.Submission{})
			if err != nil {
				panic(err)
				return
			}
			break
		}
		log.Printf("Waiting for database to be ready... (%d/10)", i+1)
		time.Sleep(1 * time.Second)
	}

	if !ready {
		panic("Database is not ready")
	}
}

func waitForS3() {
	// the following credentials are not real lol
	_ = os.Setenv("AWS_ACCESS_KEY_ID", "AKIAIOSFODNN7EXAMPLE")
	_ = os.Setenv("AWS_SECRET_ACCESS_KEY", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")
	_ = os.Setenv("AWS_REGION", "eu-west-1")
	_ = os.Setenv("AWS_BUCKET_NAME", "/test-bucket")
	_ = os.Setenv("AWS_FOLDER_NAME", "test-folder")
	_ = os.Setenv("AWS_BUCKET_ENDPOINT", "http://localhost:9444")

	ready := false
	for i := 0; i < 10; i++ {
		//send http request
		resp, err := http.Get(os.Getenv("AWS_BUCKET_ENDPOINT") + "/ui")
		if err == nil && resp.StatusCode == 200 {
			ready = true
			log.Println("S3 is ready!")
			break
		}
		log.Printf("Waiting for S3 to be ready... (%d/10)", i+1)
		time.Sleep(2 * time.Second)
	}

	if !ready {
		log.Fatalf("S3 is not ready")
	}
}

var router *gin.Engine

func TestMain(m *testing.M) {
	Setup()
	router = routes.GetRouter()

	code := m.Run()

	TearDown()
	os.Exit(code)
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
