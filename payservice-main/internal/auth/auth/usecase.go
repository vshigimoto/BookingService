package auth

import (
	"context"

	"payservice/internal/auth/entity"
)

type UseCase interface {
	GenerateToken(request GenerateTokenRequest, r context.Context) (*JwtUserToken, error)
	RenewToken(token string) (*JwtUserToken, error)
	GetToken(token string) (entity.UserToken, error)
	RegisterProcess(rq entity.UserRegister, r context.Context) error
	ConfirmUser(code string) error
	DeleteToken(token string) error
}
