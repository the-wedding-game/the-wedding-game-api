package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"the-wedding-game-api/models"
)

func migrate() {
	log.Println("Migrating database")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),
	)

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to database: %v\n", err)
		log.Fatal("Could not connect database")
	}

	_ = db.AutoMigrate(&models.User{})
	_ = db.AutoMigrate(&models.Challenge{})
	_ = db.AutoMigrate(&models.AccessToken{})
	_ = db.AutoMigrate(&models.Answer{})
	_ = db.AutoMigrate(&models.Submission{})
}
