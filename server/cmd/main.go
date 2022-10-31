package main

import (
	"context"
	"fmt"
	"sol-shop-server/config"
	"sol-shop-server/ent"
	"sol-shop-server/ent/product"
	"sol-shop-server/pkg/logger"
	"sol-shop-server/pkg/postgres"

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

	ctx := context.Background()

	queryShoes, err := client.ProductCategory.Query().QueryProducts().Order(ent.Desc(product.FieldID)).All(ctx)
	if err != nil {
		fmt.Println("error while querying shoes category")
	}

	fmt.Println(queryShoes)
}
