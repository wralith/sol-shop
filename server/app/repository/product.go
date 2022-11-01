package repository

import (
	"context"
	"sol-shop-server/ent"
	"sol-shop-server/ent/product"
	"sol-shop-server/ent/productcategory"
)

type ProductRepository struct {
	db *ent.Client
}

func newProductRepository(entClient *ent.Client) *ProductRepository {
	return &ProductRepository{db: entClient}
}

// GetProductByID implementation.
func (r *ProductRepository) GetProductByID(id int) (*ent.Product, error) {
	return r.db.Product.Query().WithProductCategory().Where(product.ID(id)).Only(context.Background())
}

// ListProducts implementation.
func (r *ProductRepository) ListProducts(limit int, offset int) (ent.Products, error) {
	return r.db.Product.Query().WithProductCategory().Limit(limit).Offset(offset).All(context.Background())
}

// ListProductsByCategoryID implementation.
func (r *ProductRepository) ListProductsByCategoryID(id int, limit int, offset int) (ent.Products, error) {
	return r.db.ProductCategory.
		Query().
		Where(productcategory.ID(id)).
		QueryProducts().
		WithProductCategory().
		Limit(limit).
		Offset(offset).
		All(context.Background())
}

// CreateProduct implementation.
func (r *ProductRepository) CreateProduct(p *ent.Product, categoryIds []int) (*ent.Product, error) {
	return r.db.Product.Create().
		SetName(p.Name).
		SetDescription(p.Description).
		SetPrice(p.Price).
		SetStock(p.Stock).
		AddProductCategoryIDs(categoryIds...).
		Save(context.Background())
}

// UpdateProduct implementation.
func (r *ProductRepository) UpdateProduct(id int, p *ent.Product, categoryIds []int) (*ent.Product, error) {
	return r.db.Product.UpdateOneID(id).
		SetName(p.Name).
		SetDescription(p.Description).
		SetStock(p.Stock).
		ClearProductCategory().
		AddProductCategoryIDs(categoryIds...).
		Save(context.Background())
}

// DeleteProduct implementation.
func (r *ProductRepository) DeleteProduct(id int) error {
	return r.db.Product.DeleteOneID(id).Exec(context.Background())
}

// Product Categories

// GetProductCategoryByID implementation.
func (r *ProductRepository) GetProductCategoryByID(id int) (*ent.ProductCategory, error) {
	return r.db.ProductCategory.Get(context.Background(), id)
}

// ListProductCategories implementation.
func (r *ProductRepository) ListProductCategories(limit int, offset int) (ent.ProductCategories, error) {
	return r.db.ProductCategory.Query().Limit(limit).Offset(offset).All(context.Background())
}

// CreateProductCategory implementation.
func (r *ProductRepository) CreateProductCategory(c *ent.ProductCategory) (*ent.ProductCategory, error) {
	return r.db.ProductCategory.Create().
		SetName(c.Name).
		SetDescription(c.Description).
		Save(context.Background())
}

// UpdateProductCategory implementation.
func (r *ProductRepository) UpdateProductCategory(id int, c *ent.ProductCategory) (*ent.ProductCategory, error) {
	return r.db.ProductCategory.UpdateOneID(id).
		SetName(c.Name).
		SetDescription(c.Description).
		Save(context.Background())
}

// DeleteProductCategory implementation.
func (r *ProductRepository) DeleteProductCategory(id int) error {
	return r.db.ProductCategory.DeleteOneID(id).Exec(context.Background())
}
