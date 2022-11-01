package repository

import "sol-shop-server/ent"

type Repository struct {
	ProductRepository *ProductRepository
}

func NewRepository(entClient *ent.Client) *Repository {
	return &Repository{
		ProductRepository: newProductRepository(entClient),
	}
}
