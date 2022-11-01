package product

import (
	"sol-shop-server/app/service"
)

type ProductController struct {
	service *service.ProductService
}

func NewProductController(service *service.ProductService) *ProductController {
	return &ProductController{
		service: service,
	}
}
