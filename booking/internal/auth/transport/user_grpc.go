package transport

import (
	"booking/internal/auth/config"
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "service/pkg/userservice/gw"
)

type UserGrpcTransport struct {
	config config.UserGrpcTransport
	client pb.UserServiceClient
}

func NewUserGrpcTransport(config config.UserGrpcTransport) *UserGrpcTransport {
	opts := []grpc.DialOption{grpc.WithInsecure()}

	conn, _ := grpc.Dial(config.Host, opts...)

	client := pb.NewUserServiceClient(conn)

	return &UserGrpcTransport{
		client: client,
		config: config,
	}
}

func (t *UserGrpcTransport) GetUserByLoin(ctx context.Context, login string) (*pb.User, error) {
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
