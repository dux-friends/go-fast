package web

import "github.com/labstack/echo/v4"

// @RouteGroup(app="web", name = "text", route = "/test")

// Index @Route(method = "GET", name = "index", route = "/")
func Index(ctx echo.Context) error {
	return ctx.JSON(200, "dsadsad")
}
