package router

// DocsRouter implementation.
func (r *Router) DocsRoutes() {
	d := r.Echo.Group("/docs")

	d.File("/spec", "docs/swagger.yaml")
	d.File("", "docs/docs.html")
}
