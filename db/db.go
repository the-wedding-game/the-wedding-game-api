package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var database *gorm.DB

func initializeDatabaseConnection() *gorm.DB {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),
	)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Println(err)
		log.Fatal("Could not connect database")
	}
	return db
}

func GetDbConnection() *gorm.DB {
	if database == nil {
		database = initializeDatabaseConnection()
	}
	return database
}
