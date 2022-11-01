package product

import (
	"net/http"
	"sol-shop-server/app/controller/body"
	"sol-shop-server/app/controller/util"

	"github.com/labstack/echo/v4"
)

// @Summary     Delete Product
// @Description Delete Product
// @ID          Delete-Product
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       id  path     int true "Product ID"
// @Success     200 {object} body.ProductBody
// @Failure     500 {object} body.Error
// @Router      /products/{id} [delete]
func (pc *ProductController) DeleteProduct(c echo.Context) error {
	id := util.ParamInt(c, "id")

	err := pc.service.DeleteProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewError(err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "deleted successfully"})
}
