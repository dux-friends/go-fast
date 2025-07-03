package routes

import (
	"dux-project/app/home/web"

	"github.com/duxweb/go-fast/route"
	"github.com/labstack/echo/v4"
)

func RouteWeb(router *route.RouterData) {
	router.Get("/home/index", web.Index, "home.index")

	router.Get("/", func(ctx echo.Context) error {
		return ctx.JSON(200, "dsadsad")
	}, "首页")
}
