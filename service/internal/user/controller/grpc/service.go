package grpc

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"service/internal/user/repository"
	pb "service/pkg/userservice/gw"
)

type Service struct {
	pb.UnimplementedUserServiceServer
	logger *zap.SugaredLogger
	repo   repository.Repository
}

func NewService(logger *zap.SugaredLogger, repo repository.Repository) *Service {
	return &Service{
		logger: logger,
		repo:   repo,
	}
}

func (s *Service) GetUserByLogin(ctx context.Context, request *pb.GetUserByLoginRequest) (*pb.GetUserByLoginResponse, error) {
	user, err := s.repo.GetUserByLogin(ctx, request.Login)
	if err != nil {
		s.logger.Errorf("failed to GetUserByLogin err: %v", err)
		return nil, fmt.Errorf("GetUserByLogin err: %w", err)
	}

	return &pb.GetUserByLoginResponse{
		Result: &pb.User{
			Id:          int32(user.Id),
			Login:       user.Login,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			IsConfirmed: user.IsConfirmed,
			Password:    user.Password,
		},
	}, nil
}
