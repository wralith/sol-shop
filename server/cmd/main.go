package main

import (
	"sol-shop-server/app/controller"
	"sol-shop-server/app/repository"
	"sol-shop-server/app/router"
	"sol-shop-server/app/service"
	"sol-shop-server/config"
	"sol-shop-server/pkg/db/postgres"
	"sol-shop-server/pkg/logger"

	"github.com/rs/zerolog/log"
)

func main() {
	c := config.NewConfig()
	logger.InitLogger(&c.Logger)

	client, err := postgres.NewPostgresEntClient(&c.Postgres)
	if err != nil {
		log.Fatal().Err(err).Msg("failed opening connection to postgres")
	}
	defer client.Close()

	if err := postgres.AutoMigrate(client); err != nil {
		log.Fatal().Err(err).Msg("failed migrate schemas to postgres")

	}

	repo := repository.NewRepository(client)
	service := service.NewService(repo)
	controller := controller.NewController(service)
	router := router.NewRouter(controller)

	router.Echo.Start(":" + c.Server.Port)
	// shoe, _ := service.ProductService.GetProductByID(2)
}
