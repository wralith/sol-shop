package product

import (
	"net/http"
	"sol-shop-server/app/controller/body"
	"sol-shop-server/app/controller/util"

	"github.com/labstack/echo/v4"
)

// @Summary     Update Product
// @Description Update Product
// @ID          Update-Product
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       product body     body.ProductBody true "New Product Data"
// @Param       id      path     int              true "Product ID"
// @Success     200     {object} body.ProductBody
// @Failure     500     {object} body.Error
// @Router      /products/{id} [put]
func (pc *ProductController) UpdateProduct(c echo.Context) error {
	id := util.ParamInt(c, "id")
	product := body.ProductBody{}

	if err := product.BindAndValidate(c); err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewError(err.Error()))
	}

	updatedProduct, err := pc.service.UpdateProduct(id, product.MapToEntWithoutCategories(), product.Categories)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewError(err.Error()))
	}

	payload := body.MapToProductBody(updatedProduct)

	return c.JSON(http.StatusOK, payload)
}
