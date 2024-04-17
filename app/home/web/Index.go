package web

import "github.com/labstack/echo/v4"

// @RouteGroup(app="web", name = "text", route = "/test")

type Annotation struct {
	Name  string
	Route string
	Func  string
}

// Index @Route(method = "GET", name = "index", route = "/")
func Index(ctx echo.Context) error {
	return ctx.JSON(200, "dsadsad")
}
