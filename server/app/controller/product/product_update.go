package product

import (
	"net/http"
	"sol-shop-server/app/controller/body"
	"sol-shop-server/app/controller/util"

	"github.com/labstack/echo/v4"
)

func (pc *ProductController) UpdateProduct(c echo.Context) error {
	id := util.ParamInt(c, "id")
	product := body.ProductBody{}
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewErrorMessage(err.Error()))
	}

	updatedProduct, err := pc.service.UpdateProduct(id, product.MapToEntWithoutCategories(), product.Categories)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewErrorMessage(err.Error()))
	}

	payload := body.MapToProductBody(updatedProduct)

	return c.JSON(http.StatusOK, payload)
}
