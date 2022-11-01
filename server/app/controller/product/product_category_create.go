package product

import (
	"net/http"
	"sol-shop-server/app/controller/body"
	"sol-shop-server/ent"

	"github.com/labstack/echo/v4"
)

func (pc *ProductController) CreateProductCategory(c echo.Context) error {
	category := ent.ProductCategory{}
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewErrorMessage(err.Error()))
	}

	payload, err := pc.service.CreateProductCategory(&category)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewErrorMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, payload)
}
