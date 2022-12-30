package auth

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("API_SECRET"))

type ClaimJWT struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// Function to generate JWT token
func GenerateJWT(email string, username string) (tokenString string, err error) {
	expiredTime := time.Now().Add(1 * time.Hour)
	claims := &ClaimJWT{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

// Function to validate JWT token
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&ClaimJWT{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*ClaimJWT)
	if !ok {
		err = errors.New("Couldn't parse claims token")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("Token has expired")
		return
	}
	return
}

// Function to take email data from user based on JWT token
func GetEmail(signedToken string) (email string, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&ClaimJWT{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*ClaimJWT)
	if !ok {
		err = errors.New("Couldn't parse claims token")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("Token has expired")
		return
	}

	return claims.Email, nil
}
