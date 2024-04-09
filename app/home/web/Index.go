package web

import "github.com/labstack/echo/v4"

// RouteGroup('xxx', '/list')

// Route('xxx', '/list')
func Index(ctx echo.Context) error {
	return ctx.JSON(200, "dsadsad")
}
