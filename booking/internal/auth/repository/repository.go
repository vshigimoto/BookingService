package repository

import (
	"booking/internal/auth/entity"
	"database/sql"
)

type UserTokenRepository interface {
	CreateUserToken(userToken entity.UserToken) error
	UpdateUserToken(userToken entity.UserToken) error
}

type Repository interface {
	UserTokenRepository
}

type Repo struct {
	main    sql.DB
	replica sql.DB
}

func NewRepository(main *sql.DB, replica *sql.DB) *Repo {
	return &Repo{
		main:    *main,
		replica: *replica,
	}
}
