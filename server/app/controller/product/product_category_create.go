package product

import (
	"net/http"
	"sol-shop-server/app/controller/body"

	"github.com/labstack/echo/v4"
)

// @Summary     Create Product Category
// @Description Create Product Category
// @ID          Create-Product-Category
// @Tags        product-categories
// @Accept      json
// @Produce     json
// @Param       category body     body.ProductCategoryBody true "New Category Data"
// @Success     201      {object} body.ProductCategoryBody
// @Failure     500      {object} body.Error
// @Router      /products/category [post]
func (pc *ProductController) CreateProductCategory(c echo.Context) error {
	category := body.ProductCategoryBody{}
	if err := category.BindAndValidate(c); err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewError(err.Error()))
	}

	payload, err := pc.service.CreateProductCategory(category.MapToEnt())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, body.NewError(err.Error()))
	}

	return c.JSON(http.StatusCreated, payload)
}
