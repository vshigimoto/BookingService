package auth

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"

	"service/internal/gallery/config"
)

var ErrExpiredToken = errors.New("expiration date validation error")

type ContextUserKey struct{}

type ContextUser struct {
	ID int64 `json:"user_id"`
}

type Service struct {
	jwtSecretKey string
}

func NewService(authConfig config.Auth) *Service {
	return &Service{
		jwtSecretKey: authConfig.JwtSecretKey,
	}
}

func (c *Service) GetJwtUser(jwtToken string) (*ContextUser, error) {
	token, err := jwt.Parse(
		jwtToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(c.jwtSecretKey), nil
		},
	)

	if err != nil {
		if validationErr, ok := err.(*jwt.ValidationError); ok {
			if validationErr.Errors&jwt.ValidationErrorExpired > 0 {
				return nil, ErrExpiredToken
			}
		}

		return nil, fmt.Errorf("failed to parse jwt err: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unexpected type %T", claims)
	}

	user, err := c.getUserFromJwt(claims)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from jwt err: %w", err)
	}

	return user, nil
}

func (c *Service) getUserFromJwt(claims jwt.MapClaims) (*ContextUser, error) {
	user := &ContextUser{}
	userId, ok := claims["user_id"]
	if !ok {
		return nil, fmt.Errorf("user is not exists in jwt")
	}

	userId = fmt.Sprintf("%.0f", userId)

	parsedUserId, err := strconv.ParseInt(userId.(string), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("unexpected in userID value: %T", userId)
	}

	user.ID = parsedUserId

	return user, nil
}

func (c *Service) GetContextUserKey() ContextUserKey {
	return ContextUserKey{}
}
