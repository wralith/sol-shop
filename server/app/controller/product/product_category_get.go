package product

import (
	"net/http"
	"sol-shop-server/app/controller/body"
	"sol-shop-server/app/controller/util"

	"github.com/labstack/echo/v4"
)

// @Summary     Get Product Category
// @Description Get Product Category
// @ID          Get-Product-Category
// @Tags        product-categories
// @Accept      json
// @Produce     json
// @Param       id  path     int true "Category ID"
// @Success     200 {object} body.ProductCategoryBody
// @Failure     500 {object} body.Error
// @Router      /products/category/{id} [get]
func (pc *ProductController) GetProductCategoryByID(c echo.Context) error {
	id := util.ParamInt(c, "id")
	res, err := pc.service.GetProductCategoryByID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewError(err.Error()))
	}

	return c.JSON(http.StatusOK, res)
}
