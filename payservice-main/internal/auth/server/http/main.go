package http

import (
	_ "payservice/internal/auth/entity"
)

// @title Auth Service API
// @version 1.0
// @description Auth service API in Go using Gin Framework
// @host localhost:8081
// @BasePath /auth

type ReqLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RespLogin struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserRegister struct {
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Phone       string `json:"Phone"`
	Login       string `json:"Login"`
	Password    string `json:"Password"`
	IsConfirmed bool   `json:"isConfirmed"`
}
