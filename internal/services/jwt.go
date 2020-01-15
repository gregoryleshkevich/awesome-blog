package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type JWTService struct{}

// GetToken by username
func (jwts JWTService) GetToken(email string) (string, time.Time) {

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(Config.JWTKey)

	return tokenString, expirationTime
}

// Verify chek token
func (jwts JWTService) Verify(token string) (error) {

	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	
	return err
}
