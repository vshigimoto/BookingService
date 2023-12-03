package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"service/internal/auth/entity"
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
	UserTokenRepository
}

type UserTokenRepository interface {
	CreateUserToken(ctx context.Context, userToken entity.UserToken) error
	UpdateUserToken(ctx context.Context, userToken entity.UserToken) error
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
