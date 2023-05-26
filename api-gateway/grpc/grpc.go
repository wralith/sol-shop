package grpc

import (
	"github.com/wralith/sol-shop/pb/gen-go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient user.UserServiceClient

func NewUserClient(port string) (UserClient, error) {
	conn, err := grpc.Dial(":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return user.NewUserServiceClient(conn), nil
}
