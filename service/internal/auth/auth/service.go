package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"service/internal/auth/config"
	"service/internal/auth/controller/consumer/dto"
	"service/internal/auth/entity"
	"service/internal/auth/repository"
	"service/internal/auth/transport"
	"service/internal/kafka"
)

type Service struct {
	repo                     repository.Repository
	jwtSecretKey             string
	passwordSecretKey        string
	userTransport            *transport.UserTransport
	userGrpcTransport        *transport.UserGrpcTransport
	userVerificationProducer *kafka.Producer
}

func NewAuthService(
	repo repository.Repository,
	authConfig config.Auth,
	userTransport *transport.UserTransport,
	userGrpcTransport *transport.UserGrpcTransport,
	userVerificationProducer *kafka.Producer,
) UseCase {
	return &Service{
		repo:                     repo,
		jwtSecretKey:             authConfig.JwtSecretKey,
		passwordSecretKey:        authConfig.PasswordSecretKey,
		userTransport:            userTransport,
		userGrpcTransport:        userGrpcTransport,
		userVerificationProducer: userVerificationProducer,
	}
}

func (a *Service) GenerateToken(ctx context.Context, request GenerateTokenRequest) (*JwtUserToken, error) {
	//user, err := a.userTransport.GetUser(ctx, request.Login)
	//if err != nil {
	//	return nil, fmt.Errorf("GetUser request err: %w", err)
	//}

	user, err := a.userGrpcTransport.GetUserByLoin(ctx, request.Login)
	if err != nil {
		return nil, fmt.Errorf("GetUser request err: %w", err)
	}

	generatedPassword := a.generatePassword(request.Password)
	if user.Password != generatedPassword {
		return nil, fmt.Errorf("password is wrong")
	}

	type MyCustomClaims struct {
		UserId int `json:"user_id"`
		jwt.RegisteredClaims
	}

	claims := MyCustomClaims{
		int(user.Id),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}

	secretKey := []byte(a.jwtSecretKey)
	claimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := claimToken.SignedString(secretKey)
	if err != nil {
		return nil, fmt.Errorf("SignedString err: %w", err)
	}

	rClaims := MyCustomClaims{
		int(user.Id),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(40 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
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
		UserId:       int(user.Id),
	}

	err = a.repo.CreateUserToken(ctx, userToken)
	if err != nil {
		return nil, fmt.Errorf("CreateUserToken err: %w", err)
	}

	jwtToken := &JwtUserToken{
		Token:        userToken.Token,
		RefreshToken: userToken.RefreshToken,
	}

	return jwtToken, nil
}

func (a *Service) RenewToken() {
	//TODO implement me
	panic("implement me")
}

func (a *Service) SendCode() {
	//TODO implement me
	panic("implement me")
}

func (a *Service) Register(ctx context.Context) error {
	// user created

	randNum1 := rand.Intn(10)
	randNum2 := rand.Intn(10)

	msg := dto.UserCode{Code: fmt.Sprintf("%d%d", randNum1, randNum2)}

	b, err := json.Marshal(&msg)
	if err != nil {
		return fmt.Errorf("failed to marshall UserCode err: %w", err)
	}

	a.userVerificationProducer.ProduceMessage(b)

	return nil
}

func (a *Service) generatePassword(password string) string {
	hash := hmac.New(sha256.New, []byte(a.passwordSecretKey))
	_, _ = hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}
