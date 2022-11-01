package service

import "sol-shop-server/app/repository"

type Service struct {
	ProductService *ProductService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		ProductService: newProductService(repo.ProductRepository),
	}
}
