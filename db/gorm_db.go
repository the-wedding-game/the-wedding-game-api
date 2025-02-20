package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"os"
	apperrors "the-wedding-game-api/errors"
)

type database struct {
	db *gorm.DB
}

func (p *database) Where(query interface{}, args ...interface{}) DatabaseInterface {
	p.db = p.db.Where(query, args...)
	return p
}

func (p *database) First(dest interface{}, where ...interface{}) DatabaseInterface {
	p.db = p.db.Limit(1).First(dest, where...)
	return p
}

func (p *database) Create(value interface{}) DatabaseInterface {
	p.db = p.db.Create(value)
	return p
}

func (p *database) Find(dest interface{}, where ...interface{}) DatabaseInterface {
	p.db = p.db.Find(dest, where...)
	return p
}

func (p *database) GetError() error {
	err := p.db.Error
	if err == nil {
		return nil
	}
	p.db.Error = nil

	if gorm.IsRecordNotFoundError(err) {
		return apperrors.NewRecordNotFoundError(err.Error())
	}

	return apperrors.NewDatabaseError(err.Error())
}

func newDatabase() DatabaseInterface {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),
	)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Printf("Error connecting to database: %v\n", err)
		log.Fatal("Could not connect database")
	}

	return &database{db: db}
}
