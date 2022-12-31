// Package jwtauth provides jwtauth
package jwtauth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type userClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

func getEnv() string {
	var secret string
	val, ok := os.LookupEnv("SECRET_SIGN_KEY")
	if !ok {
		// env not set
		secret = "supersecretkey"
	} else {
		// env is set
		secret = val
	}

	return secret
}

var signKey = getEnv()

// var signKey = []byte("supersecretkey")

// GenerateJWT function
func GenerateJWT(userID uint) (string, error) {
	claims := &userClaims{
		ID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			// set expiry to one week
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
			Issuer:    "BTH-app",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(signKey))
	if err != nil {
		log.Println(err)
		return "", err
	}

	return signedToken, nil
}

// ValidateJWT get user ID from JWT
func ValidateJWT(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*userClaims); ok && token.Valid {
		// valid token
		return claims.ID, nil
	}

	// if claims not OK
	return 0, err
}
