package product

import (
	"net/http"
	"sol-shop-server/app/controller/body"

	"github.com/labstack/echo/v4"
)

//TODO: Returns categories as an empty Array, how to fix without querying again?

func (pc *ProductController) CreateProduct(c echo.Context) error {
	product := body.ProductBody{}
	if err := product.BindAndValidate(c); err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewErrorMessage(err.Error()))
	}

	createdProduct, err := pc.service.CreateProduct(product.MapToEntWithoutCategories(), product.Categories)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewErrorMessage(err.Error()))
	}

	payload := body.MapToProductBody(createdProduct)

	return c.JSON(http.StatusOK, payload)
}
