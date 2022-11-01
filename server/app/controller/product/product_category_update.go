package product

import (
	"net/http"
	"sol-shop-server/app/controller/body"
	"sol-shop-server/app/controller/util"

	"github.com/labstack/echo/v4"
)

// @Summary     Update Product Category
// @Description Update Product Category
// @ID          Update-Product-Category
// @Tags        product-categories
// @Accept      json
// @Produce     json
// @Param       category body     body.ProductCategoryBody true "New Category Data"
// @Param       id       path     int                      true "Category ID"
// @Success     200      {object} body.ProductCategoryBody
// @Failure     500      {object} body.Error
// @Router      /products/category/{id} [put]
func (pc *ProductController) UpdateProductCategory(c echo.Context) error {
	id := util.ParamInt(c, "id")
	category := body.ProductCategoryBody{}

	if err := category.BindAndValidate(c); err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewError(err.Error()))
	}

	payload, err := pc.service.UpdateProductCategory(id, category.MapToEnt())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewError(err.Error()))
	}

	return c.JSON(http.StatusOK, payload)
}
