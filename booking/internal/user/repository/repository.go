package repository

import (
	"booking/internal/user/entity"
	"context"
	"database/sql"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (id int, err error)
	GetByLogin(ctx context.Context, login string) (*entity.User, error)
	UpdateUser(ctx context.Context, id int, user *entity.User) error
	DeleteUser(ctx context.Context, id int) error
	GetUsers(ctx context.Context) ([]entity.User, error)
}

type Repository interface {
	UserRepository
}

type Repo struct {
	main    *sql.DB
	replica *sql.DB
}

func NewRepository(main *sql.DB, replica *sql.DB) *Repo {
	return &Repo{
		main:    main,
		replica: replica,
	}
}
