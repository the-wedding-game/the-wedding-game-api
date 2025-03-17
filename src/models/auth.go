package models

import (
	"gorm.io/gorm"
	"os"
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
	conn := GetConnection()
	if err := conn.Where("username = ?", username).First(&user).GetError(); err != nil {
		if apperrors.IsRecordNotFoundError(err) {
			return false, User{}, nil
		}
		return false, User{}, err
	}
	return true, user, nil
}

func ValidatePassword(password string) error {
	expectedPassword := os.Getenv("ADMIN_PASSWORD")
	if password != expectedPassword {
		return apperrors.NewAuthenticationError("invalid password")
	}
	return nil
}

func (user User) Save() (User, error) {
	conn := GetConnection()
	if err := conn.Create(&user).GetError(); err != nil {
		return User{}, err
	}
	return user, nil
}

func (user User) GetPoints() (uint, error) {
	conn := GetConnection()
	points, err := conn.GetPointsForUser(user.ID)
	if err != nil {
		return 0, err
	}
	return points, nil
}
