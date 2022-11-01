package product

import (
	"net/http"
	"sol-shop-server/app/controller/body"
	"sol-shop-server/app/controller/util"

	"github.com/labstack/echo/v4"
)

func (pc *ProductController) ListProductCategories(c echo.Context) error {
	limit := util.QueryParamInt(c, "limit")
	offset := util.QueryParamInt(c, "offset")

	res, err := pc.service.ListProductCategories(limit, offset)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewErrorMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, res)
}
