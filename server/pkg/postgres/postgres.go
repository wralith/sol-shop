package postgres

import (
	"context"
	"fmt"
	"sol-shop-server/config"
	"sol-shop-server/ent"

	_ "github.com/lib/pq"
)

func NewPostgresEntClient(c *config.PostgresConfig) (*ent.Client, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DbName,
	)

	return ent.Open("postgres", dsn)
}

func AutoMigrate(e *ent.Client) error {
	return e.Schema.Create(context.Background())
}
