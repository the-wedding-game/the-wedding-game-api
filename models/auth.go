package models

import (
	"github.com/jinzhu/gorm"
	"the-wedding-game-api/db"
	apperrors "the-wedding-game-api/errors"
	"the-wedding-game-api/types"
)

type User struct {
	gorm.Model
	Username string         `gorm:"unique;not null"`
	Role     types.UserRole `gorm:"default:'PLAYER'"`
}

func NewUser(username string) User {
	return User{
		Username: username,
		Role:     types.Player,
	}
}

func DoesUserExist(username string) (bool, User, error) {
	var user User
	conn := db.GetDB()
	if err := conn.Where("username = ?", username).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, User{}, apperrors.NewNotFoundError("User", username)
		}
		return false, User{}, err
	}
	return true, user, nil
}

func (user User) Save() (User, error) {
	conn := db.GetDB()
	if err := conn.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
