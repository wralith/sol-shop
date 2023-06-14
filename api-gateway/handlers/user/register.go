package user

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wralith/sol-shop/api-gateway/grpc"
	pb "github.com/wralith/sol-shop/pb/gen-go/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}

func (r RegisterRequest) toProto() *pb.RegisterRequest {
	return &pb.RegisterRequest{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}
}

func Register(rpc grpc.UserClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input RegisterRequest
		if err := c.BodyParser(&input); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		res, err := rpc.Register(context.TODO(), input.toProto())
		if err != nil {
			log.Println(err)
			st := status.Convert(err)
			if st.Code() == codes.AlreadyExists {
				return c.SendStatus(fiber.StatusConflict)
			}
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(res)
	}
}
