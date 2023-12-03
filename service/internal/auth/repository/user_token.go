package repository

import (
	"context"
	"fmt"
	"service/internal/auth/metrics"

	"github.com/jmoiron/sqlx"

	"service/internal/auth/entity"
)

func (r *Repo) CreateUserToken(ctx context.Context, userToken entity.UserToken) error {
	ok, fail := metrics.DatabaseQueryTime("CreateUserToken")
	defer fail()

	q := `
INSERT INTO user_token (token, refresh_token, user_id)
VALUES (?, ?, ?)
ON DUPLICATE KEY UPDATE
	user_id = VALUES(user_id),
	token = VALUES(token),
	refresh_token = VALUES(refresh_token);
`
	query, args, err := sqlx.In(
		q,
		userToken.Token,
		userToken.RefreshToken,
		userToken.UserId,
	)

	if err != nil {
		return fmt.Errorf("query bake failed: %w", err)
	}

	_, err = r.main.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("db exec query failed: %w", err)
	}

	ok()

	return nil
}

func (r *Repo) UpdateUserToken(ctx context.Context, userToken entity.UserToken) error {
	q := `
UPDATE user_token SET token = ?, refresh_token = ? WHERE user_id = ?;
`
	query, args, err := sqlx.In(
		q,
		userToken.Token,
		userToken.RefreshToken,
		userToken.UserId,
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
