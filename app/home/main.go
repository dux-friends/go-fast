package home

import (
	"github.com/duxweb/go-fast/app"
	"github.com/duxweb/go-fast/route"
	"github.com/labstack/echo/v4"

	"dux-project/app/home/routes"
)

var config = struct {
}{}

func App() {
	app.Register(&app.Config{
		Name:     "home",
		Config:   &config,
		Init:     Init,
		Register: Register,
	})
}

func Init(t *app.Dux) {
	route.Set("web", route.New(""))
}

func Register(t *app.Dux) {
	group := route.Get("web")
	routes.RouteWeb(group)

	group.Get("/test", func(c echo.Context) error {
		// return c.String(http.StatusOK, "Hello, World!")
		return c.JSON(200, map[string]interface{}{
			"name": "dux",
			"age":  10,
		})
	}, "web.home")

}
