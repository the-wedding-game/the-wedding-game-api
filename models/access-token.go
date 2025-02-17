package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"the-wedding-game-api/db"
	"time"
)

type AccessToken struct {
	gorm.Model
	Token     string `gorm:"unique" json:"token"`
	UserID    uint   `json:"user_id"`
	ExpiresOn int64  `json:"expires_on"`
}

func generateAccessToken() string {
	return uuid.New().String()
}

func LinkAccessTokenToUser(userId uint) (AccessToken, error) {
	conn := db.GetDB()
	token := generateAccessToken()
	expiresOn := time.Now().Add(24 * time.Hour).Unix()
	accessToken := AccessToken{Token: token, UserID: userId, ExpiresOn: expiresOn}
	if err := conn.Create(&accessToken).Error; err != nil {
		return AccessToken{}, err
	}
	return accessToken, nil
}

func GetUserByAccessToken(token string) (User, error) {
	conn := db.GetDB()
	var accessToken AccessToken
	if err := conn.Where("token = ?", token).First(&accessToken).Error; err != nil {
		return User{}, err
	}
	var user User
	if err := conn.Where("id = ?", accessToken.UserID).First(&user).Error; err != nil {
		return User{}, err
	}
	fmt.Printf("User: %v\n", user)
	return user, nil
}
