package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	"service/internal/user/entity"
)

func (r *Repo) CreateUser(ctx context.Context, user entity.User) error {
	q := `
INSERT INTO user (first_name, last_name, phone, login ,password)
VALUES (?, ?, ?, ?, ?);
`
	query, args, err := sqlx.In(
		q,
		user.FirstName,
		user.LastName,
		user.Phone,
		user.Login,
		user.Password,
	)

	if err != nil {
		return fmt.Errorf("query bake failed: %w", err)
	}

	_, err = r.main.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("db exec query failed: %w", err)
	}

	return nil
}

func (r *Repo) ConfirmUser(ctx context.Context, userId int) error {
	return nil
}

func (r *Repo) GetUserByLogin(ctx context.Context, login string) (*entity.User, error) {
	q := "SELECT id, first_name, last_name, phone, login, password, is_confirmed, is_deleted, created_at, updated_at FROM user WHERE login = ?"

	query, args, err := sqlx.In(q, login)
	if err != nil {
		return nil, fmt.Errorf("query bake failed: %w", err)
	}

	var row entity.User

	if err := r.replica.GetContext(ctx, &row, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("db query failed: %w", err)
	}

	return &row, nil
}
