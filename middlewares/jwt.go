package middlewares

import (
	"management-book/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(id uint, name string) string {
	claims := &models.JWTCustomClaims{
		Id:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return t
}
