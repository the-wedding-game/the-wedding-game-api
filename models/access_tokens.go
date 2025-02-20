package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"strconv"
	"the-wedding-game-api/db"
	apperrors "the-wedding-game-api/errors"
	"time"
)

type AccessToken struct {
	gorm.Model
	Token     string `gorm:"unique"`
	UserID    uint   `gorm:"not null"`
	ExpiresOn int64  `gorm:"not null"`
}

func generateAccessToken() string {
	return uuid.New().String()
}

func LinkAccessTokenToUser(userId uint) (AccessToken, error) {
	conn := db.GetConnection()
	token := generateAccessToken()
	expiresOn := time.Now().Add(24 * time.Hour).Unix()
	accessToken := AccessToken{Token: token, UserID: userId, ExpiresOn: expiresOn}
	if err := conn.Create(&accessToken).GetError(); err != nil {
		return AccessToken{}, err
	}
	return accessToken, nil
}

func GetUserByAccessToken(token string) (User, error) {
	conn := db.GetConnection()
	var accessToken AccessToken
	if err := conn.Where("token = ?", token).First(&accessToken).GetError(); err != nil {
		if apperrors.IsRecordNotFoundError(err) {
			return User{}, apperrors.NewAccessTokenNotFoundError()
		}
		return User{}, err
	}

	var user User
	if err := conn.Where("id = ?", accessToken.UserID).First(&user).GetError(); err != nil {
		if apperrors.IsRecordNotFoundError(err) {
			return User{}, apperrors.NewNotFoundError("User", strconv.Itoa(int(accessToken.UserID)))
		}
		return User{}, err
	}

	return user, nil
}
