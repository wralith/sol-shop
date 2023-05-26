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

	user := model.User{Username: "wra", Email: "wra@wra.com", Password: "1234"}
	err = repo.Create(user)
	if err != nil {
		log.Println(err)
	}

	server := grpc.NewGRPCController(repo)
	server.Run("8081")
}
