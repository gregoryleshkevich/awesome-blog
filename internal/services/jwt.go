package services

import (
	"time"
	"awesome-blog/internal/services"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string
	jwt.StandardClaims
}

type JWTService struct{}

// GetToken by username
func (jwts JWTService) GetToken(username string) (string, time.Time) {

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(services.Config.JWTKey)

	return tokenString, expirationTime
}
