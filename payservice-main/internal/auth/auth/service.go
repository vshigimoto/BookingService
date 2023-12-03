package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"payservice/internal/auth/config"
	"payservice/internal/auth/controller/consumer/dto"
	"payservice/internal/auth/entity"
	"payservice/internal/auth/repository"
	"payservice/internal/auth/transport"
	"payservice/internal/kafka"
)

type Service struct {
	repo                     repository.Repo
	jwtSecretKey             string
	passwordSecretKey        string
	userGrpcTransport        *transport.UserGrpcTransport
	userVerificationProducer *kafka.Producer
	logger                   *zap.SugaredLogger
}

type Claims struct {
	UserID int
	jwt.RegisteredClaims
}

func CreateToken(UserID int, t int) Claims {
	return Claims{
		UserID: UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(t) * time.Minute)),
		},
	}
}

func NewService(repo repository.Repo, cfg config.Auth, userGrpcTransport *transport.UserGrpcTransport, userVerificationProducer *kafka.Producer, l *zap.SugaredLogger) UseCase {
	return &Service{
		repo:                     repo,
		jwtSecretKey:             cfg.JwtSecretKey,
		passwordSecretKey:        cfg.PasswordSecretKey,
		userGrpcTransport:        userGrpcTransport,
		userVerificationProducer: userVerificationProducer,
		logger:                   l,
	}
}

func (s *Service) GenerateToken(request GenerateTokenRequest, r context.Context) (*JwtUserToken, error) {
	secretKey := []byte(s.jwtSecretKey)

	user, err := s.userGrpcTransport.GetUserByLogin(r, request.Login)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, err
	}

	aClaims := CreateToken(int(user.Id), 15)
	rClaims := CreateToken(int(user.Id), 60)

	aClaimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, aClaims)
	accessTokenString, err := aClaimToken.SignedString(secretKey)
	if err != nil {
		s.logger.Errorf("Error with create AccessToken")
		return nil, err
	}

	rClaimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rClaims)
	refreshTokenString, err := rClaimToken.SignedString(secretKey)
	if err != nil {
		s.logger.Errorf("Error with create RefreshToken")
		return nil, err
	}

	userToken := entity.UserToken{
		UserID:       int(user.Id),
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}

	err = s.repo.CreateToken(userToken)
	if err != nil {
		s.logger.Errorf("Error with create CreateToken")
		return nil, err
	}

	result := &JwtUserToken{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}

	return result, nil
}

func (s *Service) RenewToken(token string) (*JwtUserToken, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecretKey), nil
	})
	if err != nil {
		return &JwtUserToken{}, nil
	}

	aClaims := CreateToken(claims.UserID, 15)

	rClaims := CreateToken(claims.UserID, 60)

	secretKey := []byte(s.jwtSecretKey)

	aClaimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, aClaims)
	accessTokenString, err := aClaimToken.SignedString(secretKey)
	if err != nil {
		s.logger.Errorf("Error with create AccessToken")
		return nil, err
	}

	rClaimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rClaims)
	refreshTokenString, err := rClaimToken.SignedString(secretKey)
	if err != nil {
		s.logger.Errorf("Error with create RefreshToken")
		return nil, err
	}

	userToken := entity.UserToken{
		UserID:       int(claims.UserID),
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}

	err = s.repo.CreateToken(userToken)

	result := &JwtUserToken{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}
	if err != nil {
		return &JwtUserToken{}, err
	}

	return result, nil
}

func (s *Service) GetToken(token string) (u entity.UserToken, err error) {
	u, err = s.repo.GetToken(token)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (s *Service) RegisterProcess(rq entity.UserRegister, r context.Context) error {
	randNum1 := rand.Intn(999-100) + 100
	randNum2 := rand.Intn(999-100) + 100

	userCode := dto.UserCode{Code: fmt.Sprintf("%d%d", randNum1, randNum2)}
	b, err := json.Marshal(&userCode)
	if err != nil {
		return fmt.Errorf("failed to marshall UserCode err: %w", err)
	}

	id, err := s.userGrpcTransport.RegisterUser(r, rq)
	if err != nil {
		return err
	}

	s.userVerificationProducer.ProduceMessage(b)
	err = s.repo.CreateUserCode(id, userCode.Code)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ConfirmUser(code string) error {
	err := s.repo.ConfirmUserCode(code)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteToken(token string) error {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecretKey), nil
	})
	if err != nil {
		return err
	}

	err = s.repo.DeleteToken(claims.UserID)
	if err != nil {
		return err
	}

	return nil
}
