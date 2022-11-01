package product

import (
	"net/http"
	"sol-shop-server/app/controller/body"
	"sol-shop-server/app/controller/util"

	"github.com/labstack/echo/v4"
)

func (pc *ProductController) GetProductByID(c echo.Context) error {
	id := util.ParamInt(c, "id")
	res, err := pc.service.GetProductByID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewErrorMessage(err.Error()))
	}

	payload := body.MapToProductBody(res)

	return c.JSON(http.StatusOK, payload)
}
