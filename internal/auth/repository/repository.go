package repository

import (
	"database/sql"
)

type UserTokenRepository interface {
	GetUserRole(userId int) (string, error)
}

type Repository interface {
	UserTokenRepository
}

type Repo struct {
	main    *sql.DB
	replica *sql.DB
}

func New(main *sql.DB, replica *sql.DB) *Repo {
	return &Repo{
		main:    main,
		replica: replica,
	}
}
