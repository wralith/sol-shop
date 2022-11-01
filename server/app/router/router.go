package router

import (
	"sol-shop-server/app/controller"
	mw "sol-shop-server/app/router/middleware"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	Echo       *echo.Echo
	controller *controller.Controller
}

func NewRouter(controller *controller.Controller) *Router {
	router := Router{}
	e := echo.New()
	e.Validator = &ReqValidator{validator: validator.New()}

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:        true,
		LogStatus:     true,
		LogValuesFunc: mw.EchoZerologValuesMiddleware,
	}))

	router.Echo = e
	router.controller = controller

	router.ProductRoutes()

	return &router
}

type ReqValidator struct {
	validator *validator.Validate
}

func NewReqValidator() *ReqValidator {
	return &ReqValidator{validator: validator.New()}
}

func (v *ReqValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
