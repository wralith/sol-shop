package body

import (
	"sol-shop-server/ent"

	"github.com/labstack/echo/v4"
)

type ProductBody struct {
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required,min=1000"`
	Stock       int    `json:"stock" validate:"required,min=0"`
	Categories  []int  `json:"categories"`
}

func MapToProductBody(e *ent.Product) *ProductBody {
	categories := []int{}

	for i := 0; i < len(e.Edges.ProductCategory); i++ {
		categories = append(categories, e.Edges.ProductCategory[i].ID)
	}

	return &ProductBody{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		Price:       e.Price,
		Stock:       e.Stock,
		Categories:  categories,
	}
}

func MapMultipleProductBodies(e []*ent.Product) []*ProductBody {
	var productBodies []*ProductBody

	for i := 0; i < len(e); i++ {
		productBodies = append(productBodies, MapToProductBody(e[i]))
	}

	return productBodies
}

func (b *ProductBody) MapToEntWithoutCategories() *ent.Product {
	return &ent.Product{
		ID:          b.ID,
		Name:        b.Name,
		Description: b.Description,
		Price:       b.Price,
		Stock:       b.Stock,
	}
}

func (b *ProductBody) BindAndValidate(c echo.Context) error {
	if err := c.Bind(b); err != nil {
		return err
	}

	if err := c.Validate(b); err != nil {
		return err
	}

	return nil
}
