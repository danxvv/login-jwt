package jwtHelper

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	secretKey := os.Getenv("SECRET_KEY")
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
}

func GenerateToken(userID string) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expira en 24 horas
	})
	return token.SignedString([]byte(secretKey))
}

func GetTokenClaims(token *jwt.Token) jwt.MapClaims {
	claims, _ := token.Claims.(jwt.MapClaims)
	return claims
}
