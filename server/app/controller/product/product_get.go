package product

import (
	"net/http"
	"sol-shop-server/app/controller/body"
	"sol-shop-server/app/controller/util"

	"github.com/labstack/echo/v4"
)

// @Summary     Get Product
// @Description Get Product
// @ID          Get-Product
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       id  path     int true "Product ID"
// @Success     200 {object} body.ProductBody
// @Failure     500 {object} body.Error
// @Router      /products/{id} [get]
func (pc *ProductController) GetProductByID(c echo.Context) error {
	id := util.ParamInt(c, "id")
	res, err := pc.service.GetProductByID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewError(err.Error()))
	}

	payload := body.MapToProductBody(res)

	return c.JSON(http.StatusOK, payload)
}
