package grpc

import (
	"context"

	"payservice/internal/user/entity"
	"payservice/internal/user/repository"
	pb "payservice/pkg/protobuf/userservice/gw"
)

type Service struct {
	pb.UnimplementedUserServiceServer
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetUserByLogin(ctx context.Context, request *pb.GetUserByLoginRequest) (*pb.GetUserByLoginResponse, error) {
	user := s.repo.GetUserByLogin(request.Login)

	return &pb.GetUserByLoginResponse{
		Result: &pb.User{
			Id:          int32(user.Id),
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Phone:       user.Phone,
			Login:       user.Login,
			Password:    user.Password,
			IsConfirmed: user.IsConfirmed,
		},
	}, nil
}

func (s *Service) RegisterUser(ctx context.Context, request *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	user := entity.User{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Phone:       request.Phone,
		Login:       request.Login,
		Password:    request.Password,
		IsConfirmed: request.IsConfirmed,
	}

	id, err := s.repo.CreateUser(user)
	if err != nil {
		return &pb.RegisterUserResponse{}, err
	}

	return &pb.RegisterUserResponse{
		Id: int32(id),
	}, nil
}
