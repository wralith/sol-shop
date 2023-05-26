package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wralith/sol-shop/api-gateway/config"
	"github.com/wralith/sol-shop/api-gateway/grpc"
	"github.com/wralith/sol-shop/api-gateway/handlers/user"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	userGrpcClient, err := grpc.NewUserClient("8081")
	if err != nil {
		log.Fatal(err)
	}

	app.Post("/login", user.Login(userGrpcClient))

	app.Listen(":" + config.Http.Port)
}
