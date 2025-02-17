package models

import (
	"github.com/jinzhu/gorm"
	"the-wedding-game-api/db"
)

type UserRole string

const (
	Admin  UserRole = "ADMIN"
	Player UserRole = "PLAYER"
)

type User struct {
	gorm.Model
	Username string   `json:"username" validate:"required" gorm:"not null,unique"`
	Role     UserRole `json:"role" validate:"required" gorm:"not null"`
}

func NewUser(username string) User {
	return User{
		Username: username,
		Role:     Player,
	}
}

func DoesUserExist(username string) (bool, User, error) {
	var user User
	conn := db.GetDB()
	if err := conn.Where("username = ?", username).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, User{}, nil
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
