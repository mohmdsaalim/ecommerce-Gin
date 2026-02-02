package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mohmdsaalim/ecommerce-Gin/config"
)

// var jwtsecret = []byte()
func GenerateToken(UserID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":UserID,
		"role":role,
		"exp":time.Now().Add(time.Hour).Unix(),
		"iat":time.Now().Unix(),
	}
	secret := []byte(config.AppConfig.JWT.Secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
	
}