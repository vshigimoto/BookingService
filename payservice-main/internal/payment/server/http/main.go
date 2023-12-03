package http

import "time"

// @title Payment Service API
// @version 1.0
// @description Payment service API in Go using Gin Framework
// @host localhost:8082
// @BasePath /v2

type Card struct {
	ID        int
	Requisite string `json:"requisite"`
	Exp       string `json:"exp"`
	CVC       int    `json:"cvc"`
	FullName  string `json:"fullName"`
}

type Account struct {
	Id       int
	UserID   int
	Currency string `json:"currency"`
	Balance  int    `json:"balance"`
}

type AccountTransaction struct {
	ID       int
	FromID   int    `json:"from_id"`
	ToID     int    `json:"to_id"`
	Currency string `json:"currency"`
	Money    int    `json:"money"`
	Message  string `json:"message"`
	Time     time.Time
}
