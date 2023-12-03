package http

// @title User Service API
// @version 1.0
// @description User service API in Go using Gin Framework
// @host localhost:8080
// @BasePath /v1

type User struct {
	Id          int    `db:"id"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	Phone       string `db:"phone"`
	Login       string `db:"login"`
	Password    string `db:"password"`
	IsConfirmed bool   `db:"isconfirmed"`
}
