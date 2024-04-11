package web

import "github.com/labstack/echo/v4"

type Annotation struct {
	Name  string
	Route string
	Func  string
}

// @RouteGroup(name = "testGroup", route = "/list")

// Index @Route(name = "test", route = "/test")
func Index(ctx echo.Context) error {
	return ctx.JSON(200, "dsadsad")
}
