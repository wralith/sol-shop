package main

import (
	"log"

	"github.com/wralith/sol-shop/user-service/grpc"
	"github.com/wralith/sol-shop/user-service/model"
	"github.com/wralith/sol-shop/user-service/repo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		log.Fatal("failed to connect db")
	}

	repo := repo.New(db)
	db.AutoMigrate(&model.User{})

	server := grpc.NewGRPCController(repo)
	server.Run("8081")
}
