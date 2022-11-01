package body

import "sol-shop-server/ent"

type ProductBody struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
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
