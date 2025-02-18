package models

import (
	"the-wedding-game-api/db"
)

func init() {
	conn := db.GetDB()
	conn.AutoMigrate(&User{})
	conn.AutoMigrate(&AccessToken{})
	conn.AutoMigrate(&Challenge{})
	conn.AutoMigrate(&Answer{})
	conn.AutoMigrate(&Submission{})
}
