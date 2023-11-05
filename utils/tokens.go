package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
	"vote-app/config"
)

func GenerateToken(userId int64) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "manage-budget",
		ID:        strconv.Itoa(int(userId)),
	})

	token, _ := claims.SignedString([]byte(config.GetEnv("JWT_SECRET_KEY")))
	return token
}

func VerifyToken(token string) (*jwt.Token, error) {
	jwtKey := []byte(config.GetEnv("JWT_SECRET_KEY"))
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})
}

func DecodeToken(token string) (string, error) {
	jwtKey := []byte(config.GetEnv("JWT_SECRET_KEY"))
	verifyToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := verifyToken.Claims.(*jwt.RegisteredClaims)
	if !ok && !verifyToken.Valid {
		return "", errors.New("invalid token")
	}

	return claims.ID, nil
}
