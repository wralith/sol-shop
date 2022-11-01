package body

import (
	"sol-shop-server/ent"

	"github.com/labstack/echo/v4"
)

type ProductCategoryBody struct {
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// BindAndValidate implementation.
func (p *ProductCategoryBody) BindAndValidate(c echo.Context) error {
	if err := c.Bind(p); err != nil {
		return err
	}

	if err := c.Validate(p); err != nil {
		return err
	}

	return nil
}

func MapToProductCategoryBody(e *ent.ProductCategory) *ProductCategoryBody {
	return &ProductCategoryBody{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
	}
}

func MapToProductCategoryBodies(e []*ent.ProductCategory) []*ProductCategoryBody {
	var bodies []*ProductCategoryBody

	for i := 0; i < len(e); i++ {
		bodies = append(bodies, MapToProductCategoryBody(e[i]))
	}

	return bodies
}

func (p *ProductCategoryBody) MapToEnt() *ent.ProductCategory {
	return &ent.ProductCategory{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
	}
}
