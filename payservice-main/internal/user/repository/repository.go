package repository

import (
	"database/sql"

	"payservice/internal/user/entity"
)

type Repository interface {
	UserRepository
}

type UserRepository interface {
	CreateUser(user entity.User) (id int, err error)
	GetAllUsers() []entity.User
	UpdateUser(id string, user entity.User) error
	DeleteUser(id string) error
	GetUserByID(id string) entity.User
	GetUserByLogin(login string) entity.User
}

type Repo struct {
	main    *sql.DB
	replica *sql.DB
}

func NewUserRepository(mainDB *sql.DB, replicaDB *sql.DB) *Repo {
	return &Repo{
		main:    mainDB,
		replica: replicaDB,
	}
}
