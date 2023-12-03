package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"service/internal/user/entity"
)

type db interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	WithTx(ctx context.Context, fn func(ctx context.Context) error) error
}

type Repository interface {
	UserRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, user entity.User) error
	ConfirmUser(ctx context.Context, userId int) error
	GetUserByLogin(ctx context.Context, login string) (*entity.User, error)
}

var (
	ErrNotFound = errors.New("not found")
)

type Repo struct {
	main    db
	replica db
}

func NewRepository(main db, replica db) *Repo {
	return &Repo{
		main:    main,
		replica: replica,
	}
}
