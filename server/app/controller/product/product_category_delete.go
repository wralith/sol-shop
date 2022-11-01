package product

import (
	"net/http"
	"sol-shop-server/app/controller/body"
	"sol-shop-server/app/controller/util"

	"github.com/labstack/echo/v4"
)

// @Summary     Delete Product Category
// @Description Delete Product Category
// @ID          Delete-Product-Category
// @Tags        product-categories
// @Accept      json
// @Produce     json
// @Param       id  path     int true "Category ID"
// @Success     200 {object} map[string]string
// @Failure     500 {object} body.Error
// @Router      /products/category/{id} [delete]
func (pc *ProductController) DeleteProductCategory(c echo.Context) error {
	id := util.ParamInt(c, "id")

	err := pc.service.DeleteProductCategory(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewError(err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "category deleted successfully"})
}
