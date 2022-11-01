package product

import (
	"net/http"
	"sol-shop-server/app/controller/body"

	"github.com/labstack/echo/v4"
)

//TODO: Returns categories as an empty Array, how to fix without querying again?

// @Summary     Create Product.
// @Description Create Product.
// @ID          Create-Product
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       product body     body.ProductBody true "New Product Data"
// @Success     201     {object} body.ProductBody
// @Failure     500     {object} body.Error
// @Router      /products [post]
func (pc *ProductController) CreateProduct(c echo.Context) error {
	product := body.ProductBody{}
	if err := product.BindAndValidate(c); err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewError(err.Error()))
	}

	createdProduct, err := pc.service.CreateProduct(product.MapToEntWithoutCategories(), product.Categories)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewError(err.Error()))
	}

	payload := body.MapToProductBody(createdProduct)

	return c.JSON(http.StatusCreated, payload)
}
