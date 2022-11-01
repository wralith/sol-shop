package service

import (
	"sol-shop-server/app/repository"
	"sol-shop-server/ent"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func newProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// GetProductByID implementation.
func (s *ProductService) GetProductByID(id int) (*ent.Product, error) {
	return s.repo.GetProductByID(id)
}

// ListProducts implementation.
func (s *ProductService) ListProducts(limit int, offset int) (ent.Products, error) {
	return s.repo.ListProducts(limit, offset)
}

func (s *ProductService) ListProductsByCategoryID(id int, limit int, offset int) (ent.Products, error) {
	return s.repo.ListProductsByCategoryID(id, limit, offset)
}

// CreateProduct implementation.
func (s *ProductService) CreateProduct(product *ent.Product, categoryIds []int) (*ent.Product, error) {
	return s.repo.CreateProduct(product, categoryIds)
}

// UpdateProduct implementation.
func (s *ProductService) UpdateProduct(id int, p *ent.Product, categoryIds []int) (*ent.Product, error) {
	return s.repo.UpdateProduct(id, p, categoryIds)
}

// DeleteProduct implementation.
func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.DeleteProduct(id)
}

// Product Categories

// GetProductCategoryByID implementation.
func (s *ProductService) GetProductCategoryByID(id int) (*ent.ProductCategory, error) {
	return s.repo.GetProductCategoryByID(id)
}

// ListProductCategories implementation.
func (s *ProductService) ListProductCategories(limit int, offset int) (ent.ProductCategories, error) {
	return s.repo.ListProductCategories(limit, offset)
}

// CreateProductCategory implementation.
func (s *ProductService) CreateProductCategory(c *ent.ProductCategory) (*ent.ProductCategory, error) {
	return s.repo.CreateProductCategory(c)
}

// UpdateProductCategory implementation.
func (s *ProductService) UpdateProductCategory(id int, c *ent.ProductCategory) (*ent.ProductCategory, error) {
	return s.repo.UpdateProductCategory(id, c)
}

// DeleteProductCategory implementation.
func (s *ProductService) DeleteProductCategory(id int) error {
	return s.repo.DeleteProductCategory(id)
}
