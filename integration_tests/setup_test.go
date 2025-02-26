package integrationtests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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

	for i := 0; i < 10; i++ {
		db, err := gorm.Open(postgres.Open(dbURI))
		if err == nil {
			log.Println("Database is ready!")
			log.Println("Migrating schema...")
			err := db.Debug().AutoMigrate(&models.User{}, &models.AccessToken{}, &models.Challenge{}, &models.Answer{}, &models.Submission{})
			if err != nil {
				log.Fatalf("Failed to migrate schema: %v", err)
				return
			}
			break
		}
		log.Printf("Waiting for database to be ready... (%d/10)", i+1)
		time.Sleep(1 * time.Second)
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
