package integrationtests

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"os/exec"
	"the-wedding-game-api/models"
	"time"
)

func Setup() {
	dockerComposeUp()
	setupTestDb()
}

func TearDown() {
	dockerComposeDown()
}

func dockerComposeUp() {
	dockerComposeDown()
	log.Println("Starting Docker Compose...")

	cmdUp := exec.Command("docker-compose", "-f", "docker-compose.yml", "up", "-d")
	if err := cmdUp.Run(); err != nil {
		log.Fatalf("Failed to start Docker Compose: %v", err)
	}
}

func dockerComposeDown() {
	log.Println("Stopping Docker Compose...")

	cmdDown := exec.Command("docker-compose", "-f", "docker-compose.yml", "down")
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
		db, err := gorm.Open("postgres", dbURI)
		if err == nil {
			log.Println("Database is ready!")
			log.Println("Migrating schema...")
			db.AutoMigrate(&models.User{}, &models.AccessToken{}, &models.Challenge{}, &models.Answer{}, &models.Submission{})
			break
		}
		log.Printf("Waiting for database to be ready... (%d/10)", i+1)
		time.Sleep(1 * time.Second)
	}
}
