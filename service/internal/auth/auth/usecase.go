package auth

import "context"

type UseCase interface {
	GenerateToken(ctx context.Context, request GenerateTokenRequest) (*JwtUserToken, error)
	RenewToken()
	SendCode()
	Register(ctx context.Context) error
}
