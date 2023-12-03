package auth

import (
	"booking/internal/auth/config"
	"booking/internal/auth/entity"
	"booking/internal/auth/repository"
	"booking/internal/auth/transport"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type Service struct {
	repo              repository.Repository
	jwtSecretKey      string
	passwordSecretKey string
	userTransport     *transport.UserTransport
}

func NewAuthService(repo repository.Repository, cfg config.Config, userTransport *transport.UserTransport) *Service {
	return &Service{
		repo:              repo,
		jwtSecretKey:      cfg.Auth.JwtSecretKey,
		passwordSecretKey: cfg.Auth.PasswordSecretKey,
		userTransport:     userTransport,
	}
}

func (s *Service) GenerateToken(ctx context.Context, request GenerateTokenRequest) (*JwtUserToken, error) {
	user, err := s.userTransport.GetUser(ctx, request.Login)
	if err != nil {
		return nil, fmt.Errorf("GetUser request err: %w", err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, fmt.Errorf("password is wrong")
	}
	type MyCustomClaims struct {
		UserId   string `json:"user_id"`
		UserRole string `json:"user_role"`
		jwt.RegisteredClaims
	}
	userRole, err := s.repo.GetUserRole(user.Id)
	if err != nil {
		return nil, err
	}
	claims := MyCustomClaims{
		strconv.Itoa(user.Id),
		userRole,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	secretKey := []byte(s.jwtSecretKey)
	claimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := claimToken.SignedString(secretKey)
	if err != nil {
		return nil, fmt.Errorf("SignedString err: %w", err)
	}

	rClaims := MyCustomClaims{
		strconv.Itoa(user.Id),
		userRole,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(40 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	rClaimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rClaims)

	refreshTokenString, err := rClaimToken.SignedString(secretKey)
	if err != nil {
		return nil, fmt.Errorf("SignedString err: %w", err)
	}

	userToken := entity.UserToken{
		Token:        tokenString,
		RefreshToken: refreshTokenString,
		UserId:       strconv.Itoa(user.Id),
	}

	err = s.repo.UpdateUserToken(userToken)
	if err != nil {
		return nil, fmt.Errorf("CreateUserToken err: %w", err)
	}

	jwtToken := &JwtUserToken{
		Token:        userToken.Token,
		RefreshToken: userToken.RefreshToken,
	}
	return jwtToken, nil
}
