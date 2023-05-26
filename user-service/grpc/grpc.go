package grpc

import (
	context "context"
	"log"
	"net"

	"github.com/wralith/sol-shop/pb/gen-go/user"
	"github.com/wralith/sol-shop/user-service/repo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCController struct {
	user.UnimplementedUserServiceServer
	repo *repo.Repo
}

func NewGRPCController(repo *repo.Repo) *GRPCController {
	return &GRPCController{repo: repo}
}

func (c *GRPCController) Run(port string) {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	user.RegisterUserServiceServer(server, c)

	err = server.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}

	defer server.Stop()
}

func (c *GRPCController) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	entity, err := c.repo.FindByEmail(req.GetEmail())
	if err != nil {
		// TODO: Check error type
		return nil, status.Error(codes.NotFound, "email was not found")
	}

	res := &user.LoginResponse{
		User: &user.User{
			Id:       uint32(entity.ID),
			Email:    entity.Email,
			Username: entity.Username,
		},
	}

	return res, nil
}
