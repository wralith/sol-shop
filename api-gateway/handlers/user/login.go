package user

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wralith/sol-shop/api-gateway/grpc"
	pb "github.com/wralith/sol-shop/pb/gen-go/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TODO: Add validation
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r LoginRequest) toProto() *pb.LoginRequest {
	return &pb.LoginRequest{
		Email:    r.Email,
		Password: r.Password,
	}
}

type LoginResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func newLoginResponse(proto *pb.LoginResponse, token string) LoginResponse {
	return LoginResponse{
		ID:       uint(proto.User.Id),
		Username: proto.User.Username,
		Email:    proto.User.Email,
		Token:    token,
	}
}

// TODO: Duration and secret from config
func Login(rpc grpc.UserClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input LoginRequest
		if err := c.BodyParser(&input); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		user, err := rpc.Login(context.TODO(), input.toProto())
		if err != nil {
			log.Println(err)
			st := status.Convert(err)
			if st.Code() == codes.NotFound {
				return c.SendStatus(fiber.StatusUnauthorized)
			}
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		token, err := generateToken(user, time.Hour, []byte("secret"))
		if err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(newLoginResponse(user, token))
	}
}

func generateToken(data *pb.LoginResponse, duration time.Duration, key []byte) (string, error) {
	user := data.GetUser()
	exp := time.Now().Add(duration).Unix()
	claims := jwt.MapClaims{
		"sub":      user.GetId(),
		"email":    user.GetEmail(),
		"username": user.GetUsername(),
		"exp":      exp,
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
}
