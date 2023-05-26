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

type LoginRequest struct {
	Email    string
	Password string
}

func (r LoginRequest) toProto() *pb.LoginRequest {
	return &pb.LoginRequest{
		Email:    r.Email,
		Password: r.Password,
	}
}

func Login(rpc grpc.UserClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input LoginRequest
		if err := c.BodyParser(&input); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		res, err := rpc.Login(context.TODO(), input.toProto())
		if err != nil {
			log.Println(err)
			st := status.Convert(err)
			if st.Code() == codes.NotFound {
				return c.SendStatus(fiber.StatusUnauthorized)
			}
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(res)
	}
}
