package entity

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type SecretCode struct {
	Id        int
	Code      string
	UserId    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MyJWTClaims struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
	*jwt.RegisteredClaims
}
