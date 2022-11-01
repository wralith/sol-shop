package product

import (
	"net/http"
	"sol-shop-server/app/controller/body"
	"sol-shop-server/app/controller/util"

	"github.com/labstack/echo/v4"
)

// @Summary     List Product Categories
// @Description List Product Categories
// @ID          List-Product-Categories
// @Tags        product-categories
// @Accept      json
// @Produce     json
// @Param       limit  query    int false "Limit"
// @Param       offset query    int false "offset"
// @Success     200    {object} []body.ProductCategoryBody
// @Failure     500    {object} body.Error
// @Router      /products/category/list [get]
func (pc *ProductController) ListProductCategories(c echo.Context) error {
	limit := util.QueryParamInt(c, "limit")
	offset := util.QueryParamInt(c, "offset")

	res, err := pc.service.ListProductCategories(limit, offset)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewError(err.Error()))
	}

	payload := body.MapToProductCategoryBodies(res)

	return c.JSON(http.StatusOK, payload)
}
