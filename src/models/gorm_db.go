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
			ORDER BY ID 
		`).Scan(&challenges)
	} else {
		p.db = p.db.Raw(`
			SELECT ID, Name, Description, Points, Image, Type, Status
			FROM challenges
			WHERE status = ?
			ORDER BY ID 
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
		WHERE submissions.user_id = ? AND challenges.status = ?
		GROUP BY submissions.user_id
		`, userId, types.ActiveChallenge).Scan(&points)

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
		WHERE challenges.status = ?
		GROUP BY users.username
		ORDER BY points DESC
		`, types.ActiveChallenge).Scan(&leaderboard)

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
		WHERE challenges.type = ? AND challenges.status = ?
		ORDER BY submissions.created_at DESC
	`, types.UploadPhotoChallenge, types.ActiveChallenge).Scan(&gallery)

	if tx.Error != nil {
		return nil, apperrors.NewDatabaseError(tx.Error.Error())
	}

	return gallery, nil
}

func (p *database) HasSubmissions(challengeId uint) (bool, error) {
	var count int64
	tx := p.db.Raw(`
		SELECT COUNT(*) AS count
		FROM submissions
		WHERE challenge_id = ?
	`, challengeId).Scan(&count)
	if tx.Error != nil {
		return false, apperrors.NewDatabaseError(tx.Error.Error())
	}

	return count > 0, nil
}

func (p *database) GetChallengeById(challengeId uint) (Challenge, error) {
	var challenge Challenge
	tx := p.db.Raw(`
		SELECT *
		FROM challenges
		WHERE id = ?
	`, challengeId).Scan(&challenge)

	if tx.Error != nil {
		return Challenge{}, apperrors.NewDatabaseError(tx.Error.Error())
	}

	return challenge, nil
}

func (p *database) UpdateChallenge(existingChallenge Challenge, updateChallengeRequest types.UpdateChallengeRequest) (Challenge, error) {
	if updateChallengeRequest.Name == "" {
		updateChallengeRequest.Name = existingChallenge.Name
	}
	if updateChallengeRequest.Description == "" {
		updateChallengeRequest.Description = existingChallenge.Description
	}
	if updateChallengeRequest.Points == 0 {
		updateChallengeRequest.Points = existingChallenge.Points
	}
	if updateChallengeRequest.Image == "" {
		updateChallengeRequest.Image = existingChallenge.Image
	}
	if updateChallengeRequest.Status == "" {
		updateChallengeRequest.Status = existingChallenge.Status
	}
	if updateChallengeRequest.Type == "" {
		updateChallengeRequest.Type = existingChallenge.Type
	}

	var updatedChallenge Challenge
	tx := p.db.Raw(`
		UPDATE challenges
		SET name = ?, description = ?, points = ?, image = ?, status = ?, type = ?
		WHERE id = ?
		RETURNING *`,
		updateChallengeRequest.Name, updateChallengeRequest.Description, updateChallengeRequest.Points, updateChallengeRequest.Image, updateChallengeRequest.Status, updateChallengeRequest.Type, existingChallenge.ID,
	).Scan(&updatedChallenge)

	if tx.Error != nil {
		return Challenge{}, apperrors.NewDatabaseError(tx.Error.Error())
	}

	return updatedChallenge, nil
}

func (p *database) UpdateAnswer(challengeId uint, answer string) (Answer, error) {
	var updatedAnswer Answer
	tx := p.db.Raw(`
		UPDATE answers
		SET value = ?
		WHERE challenge_id = ?
		RETURNING *
	`, answer, challengeId).Scan(&updatedAnswer)

	if tx.Error != nil {
		return Answer{}, apperrors.NewDatabaseError(tx.Error.Error())
	}

	if tx.RowsAffected == 0 {
		return Answer{}, apperrors.NewRecordNotFoundError(fmt.Sprintf("Answer with Challenge ID %d not found", challengeId))
	}

	return updatedAnswer, nil
}

func (p *database) DeleteAnswer(challengeId uint) error {
	tx := p.db.Exec(`
		DELETE FROM answers
		WHERE challenge_id = ?
	`, challengeId)

	if tx.Error != nil {
		return apperrors.NewDatabaseError(tx.Error.Error())
	}

	if tx.RowsAffected == 0 {
		return apperrors.NewRecordNotFoundError(fmt.Sprintf("Answer with Challenge ID %d not found", challengeId))
	}

	return nil
}

func (p *database) GetSubmissionsForChallenge(challengeId uint) ([]types.SubmissionForChallenge, error) {
	var submissions = make([]types.SubmissionForChallenge, 0)
	tx := p.db.Raw(`
		SELECT
		    submissions.id,
		    submissions.answer,
		    submissions.challenge_id AS "ChallengeId",
		    challenges.name AS "ChallengeName",
		    submissions.user_id AS "UserId",
		    users.username
		FROM submissions
		INNER JOIN users ON submissions.user_id = users.id
		INNER JOIN challenges ON submissions.challenge_id = challenges.id
		WHERE challenge_id = ?
	`, challengeId).Scan(&submissions)

	if tx.Error != nil {
		return nil, apperrors.NewDatabaseError(tx.Error.Error())
	}

	return submissions, nil
}

func (p *database) GetAnswerForChallenge(challengeId uint) (string, error) {
	var answer string
	tx := p.db.Raw(`
		SELECT value
		FROM answers
		WHERE challenge_id = ?
	`, challengeId).Scan(&answer)

	if tx.Error != nil {
		return "", apperrors.NewDatabaseError(tx.Error.Error())
	}

	if tx.RowsAffected == 0 {
		return "", apperrors.NewRecordNotFoundError(fmt.Sprintf("Answer with Challenge ID %d not found", challengeId))
	}

	return answer, nil
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
