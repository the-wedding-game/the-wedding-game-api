package models

import (
	"the-wedding-game-api/types"
)

var databaseConnection DatabaseInterface

type DatabaseInterface interface {
	GetSession() DatabaseInterface
	Where(query interface{}, args ...interface{}) DatabaseInterface
	First(dest interface{}, where ...interface{}) DatabaseInterface
	Create(value interface{}) DatabaseInterface
	Find(dest interface{}, where ...interface{}) DatabaseInterface
	GetAllChallenges(showInactive bool) ([]Challenge, error)
	GetPointsForUser(userId uint) (uint, error)
	GetLeaderboard() ([]types.LeaderboardEntry, error)
	GetGallery() ([]types.GalleryItem, error)
	GetError() error
}

func getConnection() DatabaseInterface {
	if databaseConnection != nil {
		return databaseConnection.GetSession()
	}
	databaseConnection = newDatabase()
	return databaseConnection.GetSession()
}

func ResetConnection() {
	databaseConnection = nil
}

var GetConnection = getConnection
