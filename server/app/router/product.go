package router

// ProductRoutes implementation.
func (r *Router) ProductRoutes() {
	p := r.Echo.Group("/products")

	p.GET("/:id", r.controller.ProductController.GetProductByID)
	p.GET("/list", r.controller.ProductController.ListProducts)
	p.POST("", r.controller.ProductController.CreateProduct)
	p.DELETE("/:id", r.controller.ProductController.DeleteProduct)
	p.PUT("/:id", r.controller.ProductController.UpdateProduct)

	p.GET("/category/:id", r.controller.ProductController.GetProductCategoryByID)
	p.GET("/category/list", r.controller.ProductController.ListProductCategories)
	p.POST("/category", r.controller.ProductController.CreateProductCategory)
	p.DELETE("/category/:id", r.controller.ProductController.DeleteProductCategory)
	p.PUT("/category/:id", r.controller.ProductController.UpdateProductCategory)
}
