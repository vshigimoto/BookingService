package user

import (
	"context"

	"service/internal/user/entity"
)

type UseCase interface {
	GetUserByLogin(ctx context.Context, login string) (*entity.User, error)
}
