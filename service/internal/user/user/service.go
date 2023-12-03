package user

import (
	"context"
	"fmt"

	"service/internal/user/entity"
	"service/internal/user/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) UseCase {
	return &Service{repo: repo}
}

func (s *Service) GetUserByLogin(ctx context.Context, login string) (*entity.User, error) {
	user, err := s.repo.GetUserByLogin(ctx, login)
	if err != nil {
		return nil, fmt.Errorf("GetUserByLogin err: %w", err)
	}

	return user, nil
}
