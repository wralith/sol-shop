package product

import (
	"net/http"
	"sol-shop-server/app/controller/body"

	"github.com/labstack/echo/v4"
)

func (pc *ProductController) CreateProductCategory(c echo.Context) error {
	category := body.ProductCategoryBody{}
	if err := category.BindAndValidate(c); err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewErrorMessage(err.Error()))
	}

	payload, err := pc.service.CreateProductCategory(category.MapToEnt())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewErrorMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, payload)
}
