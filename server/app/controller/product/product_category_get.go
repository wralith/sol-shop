package product

import (
	"net/http"
	"sol-shop-server/app/controller/body"
	"sol-shop-server/app/controller/util"

	"github.com/labstack/echo/v4"
)

func (pc *ProductController) GetProductCategoryByID(c echo.Context) error {
	id := util.ParamInt(c, "id")
	res, err := pc.service.GetProductCategoryByID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewErrorMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, res)
}
