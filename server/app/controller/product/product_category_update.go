package product

import (
	"net/http"
	"sol-shop-server/app/controller/body"
	"sol-shop-server/app/controller/util"

	"github.com/labstack/echo/v4"
)

func (pc *ProductController) UpdateProductCategory(c echo.Context) error {
	id := util.ParamInt(c, "id")
	category := body.ProductCategoryBody{}

	if err := category.BindAndValidate(c); err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewErrorMessage(err.Error()))
	}

	payload, err := pc.service.UpdateProductCategory(id, category.MapToEnt())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewErrorMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, payload)
}
