package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mohmdsaalim/ecommerce-Gin/config"
)

// var jwtsecret = []byte()
// TokenDetails keeps track of token properties
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
}

// GenerateToken generates an access token
func GenerateToken(UserID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": UserID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 1).Unix(), // 1 hour access token
		"iat":     time.Now().Unix(),
	}
	secret := []byte(config.AppConfig.JWT.Secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// GenerateRefreshToken generates a refresh token
func GenerateRefreshToken(UserID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": UserID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days refresh token
		"iat":     time.Now().Unix(),
	}
	secret := []byte(config.AppConfig.JWT.Secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// ValidateToken validates a JWT token and returns the claims
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(config.AppConfig.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
