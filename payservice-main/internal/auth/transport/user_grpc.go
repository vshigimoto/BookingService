package transport

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"payservice/internal/auth/config"
	"payservice/internal/auth/entity"
	pb "payservice/pkg/protobuf/userservice/gw"
)

type UserGrpcTransport struct {
	config config.UserGrpcTransport
	client pb.UserServiceClient
}

func NewUserGrpcTransport(config config.UserGrpcTransport) *UserGrpcTransport {
	//nolint:all
	opts := []grpc.DialOption{grpc.WithInsecure()}

	conn, _ := grpc.Dial(config.Host, opts...)

	client := pb.NewUserServiceClient(conn)

	return &UserGrpcTransport{
		client: client,
		config: config,
	}
}

func (t *UserGrpcTransport) GetUserByLogin(ctx context.Context, login string) (*pb.User, error) {
	resp, err := t.client.GetUserByLogin(ctx, &pb.GetUserByLoginRequest{
		Login: login,
	})

	if err != nil {
		return nil, fmt.Errorf("cannot GetUserByLogin: %w", err)
	}

	if resp == nil {
		return nil, fmt.Errorf("not found")
	}

	return resp.Result, nil
}

func (t *UserGrpcTransport) RegisterUser(ctx context.Context, rq entity.UserRegister) (int, error) {
	resp, err := t.client.RegisterUser(ctx, &pb.RegisterUserRequest{
		FirstName:   rq.FirstName,
		LastName:    rq.LastName,
		Phone:       rq.Phone,
		Login:       rq.Login,
		Password:    rq.Password,
		IsConfirmed: rq.IsConfirmed,
	})

	if err != nil {
		return 0, fmt.Errorf("cannot GetUserByLogin: %w", err)
	}

	if resp == nil {
		return 0, fmt.Errorf("not found")
	}

	return int(resp.Id), nil
}
