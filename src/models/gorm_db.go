package models

import (
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	apperrors "the-wedding-game-api/errors"
	"the-wedding-game-api/types"
	"time"
)

type database struct {
	db                *gorm.DB
	initialConnection *gorm.DB
}

func (p *database) GetSession() DatabaseInterface {
	return &database{
		db:                p.db.Session(&gorm.Session{}),
		initialConnection: p.db.Session(&gorm.Session{}),
	}
}

func (p *database) Where(query interface{}, args ...interface{}) DatabaseInterface {
	p.db = p.db.Where(query, args...)
	return p
}

func (p *database) First(dest interface{}, where ...interface{}) DatabaseInterface {
	p.db = p.db.Limit(1).First(dest, where...)
	p.initialConnection.Error = p.db.Error
	p.db = p.initialConnection
	return p
}

func (p *database) Create(value interface{}) DatabaseInterface {
	p.db = p.db.Create(value)
	p.initialConnection.Error = p.db.Error
	p.db = p.initialConnection
	return p
}

func (p *database) Find(dest interface{}, where ...interface{}) DatabaseInterface {
	p.db = p.db.Find(dest, where...)
	p.initialConnection.Error = p.db.Error
	p.db = p.initialConnection
	return p
}

func (p *database) GetAllChallenges(showInactive bool) ([]Challenge, error) {
	var challenges []Challenge
	if showInactive {
		p.db = p.db.Raw(`
			SELECT ID, Name, Description, Points, Image, Type, Status
			FROM challenges
		`).Scan(&challenges)
	} else {
		p.db = p.db.Raw(`
			SELECT ID, Name, Description, Points, Image, Type, Status
			FROM challenges
			WHERE status = ?
		`, types.ActiveChallenge).Scan(&challenges)
	}

	return challenges, nil
}

func (p *database) GetPointsForUser(userId uint) (uint, error) {
	var points uint
	tx := p.db.Raw(`
		SELECT SUM(challenges.points) AS points
		FROM submissions
		INNER JOIN challenges ON submissions.challenge_id = challenges.id
		WHERE submissions.user_id = ?
		GROUP BY submissions.user_id
		`, userId).Scan(&points)

	if tx.Error != nil {
		return 0, apperrors.NewDatabaseError(tx.Error.Error())
	}

	return points, nil
}

func (p *database) GetLeaderboard() ([]types.LeaderboardEntry, error) {
	var leaderboard []types.LeaderboardEntry
	tx := p.db.Raw(`
		SELECT users.username, SUM(challenges.points) AS points
		FROM submissions
		INNER JOIN users ON submissions.user_id = users.id
		INNER JOIN challenges ON submissions.challenge_id = challenges.id
		GROUP BY users.username
		ORDER BY points DESC
		`).Scan(&leaderboard)

	if tx.Error != nil {
		return nil, apperrors.NewDatabaseError(tx.Error.Error())
	}

	return leaderboard, nil
}

func (p *database) GetGallery() ([]types.GalleryItem, error) {
	var gallery []types.GalleryItem
	tx := p.db.Raw(`
		SELECT 
		    submissions.answer AS Url,
		    users.username AS "SubmittedBy"
		FROM submissions
		INNER JOIN users ON submissions.user_id = users.id
		INNER JOIN challenges ON submissions.challenge_id = challenges.id
		WHERE challenges.type = ?
		ORDER BY submissions.created_at DESC
	`, types.UploadPhotoChallenge).Scan(&gallery)

	if tx.Error != nil {
		return nil, apperrors.NewDatabaseError(tx.Error.Error())
	}

	return gallery, nil
}

func (p *database) GetError() error {
	err := p.db.Error
	if err == nil {
		return nil
	}
	p.db.Error = nil

	if errors.Is(err, gorm.ErrRecordNotFound) {
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

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to database: %v\n", err)
		log.Fatal("Could not connect database")
	}

	conn, _ := db.DB()
	conn.SetConnMaxIdleTime(time.Minute * 5)
	conn.SetConnMaxLifetime(time.Minute * 15)
	conn.SetMaxOpenConns(100)

	return &database{
		db:                db,
		initialConnection: db,
	}
}
