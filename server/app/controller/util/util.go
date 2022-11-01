package util

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func ParamInt(c echo.Context, p string) int {
	param := c.Param(p)
	res, err := strconv.Atoi(param)
	if err != nil {
		return 0
	}
	return res
}

func QueryParamInt(c echo.Context, p string) int {
	param := c.QueryParam(p)
	res, err := strconv.Atoi(param)
	if err != nil {
		return 0
	}
	return res
}
