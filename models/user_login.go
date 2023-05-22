package models

import "github.com/golang-jwt/jwt/v4"

type JwtCustomClaims struct {
	ID   uint   `json:"userId" form:"userId"`
	Name string `json:"name" form:"name"`
	jwt.RegisteredClaims
}
