package models

import "github.com/golang-jwt/jwt/v5"

type JWTCustomClaims struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}
