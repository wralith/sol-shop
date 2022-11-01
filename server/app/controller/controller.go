package controller

import (
	"sol-shop-server/app/controller/product"
	"sol-shop-server/app/service"
)

type Controller struct {
	ProductController *product.ProductController
}

// New Controller.
func NewController(service *service.Service) *Controller {
	return &Controller{
		ProductController: product.NewProductController(service.ProductService),
	}
}
