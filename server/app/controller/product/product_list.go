package product

import (
	"net/http"
	"sol-shop-server/app/controller/body"
	"sol-shop-server/app/controller/util"
	"sol-shop-server/ent"

	"github.com/labstack/echo/v4"
)

func (pc *ProductController) ListProducts(c echo.Context) error {
	var err error
	var res ent.Products

	limit := util.QueryParamInt(c, "limit")
	offset := util.QueryParamInt(c, "offset")
	categoryId := util.QueryParamInt(c, "category")

	if categoryId == 0 {
		res, err = pc.service.ListProducts(limit, offset)
	} else {
		res, err = pc.service.ListProductsByCategoryID(categoryId, limit, offset)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewErrorMessage(err.Error()))
	}

	payload := body.MapMultipleProductBodies(res)

	return c.JSON(http.StatusOK, payload)
}
