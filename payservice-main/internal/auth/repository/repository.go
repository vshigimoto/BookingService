package repository

import (
	"database/sql"

	"payservice/internal/auth/entity"
)

type Repository interface {
	AuthRepository
}

type AuthRepository interface {
	CreateToken(u entity.UserToken) error
	UpdateToken() (err error)
	DeleteToken(id int) (err error)
	GetToken(token string) (entity.UserToken, error)
	CreateUserCode(Code string) error
	ConfirmUserCode() error
}

type Repo struct {
	main    *sql.DB
	replica *sql.DB
}

func NewRepository(mainDB *sql.DB, replicaDB *sql.DB) *Repo {
	return &Repo{
		main:    mainDB,
		replica: replicaDB,
	}
}
