package runtime

import (
	appHomeWeb "dux-project/app/home/web"
	"github.com/duxweb/go-fast/annotation"
)

var Annotations = []*annotation.File{
	{
		Name: "dux-project/app/home/web",
		Annotations: []*annotation.Annotation{
			{
				Name: "RouteGroup",
				Params: map[string]any{
					"app":   "web",
					"name":  "text",
					"route": "/test",
				},
			},
			{
				Name: "Route",
				Params: map[string]any{
					"method": "GET",
					"name":   "index",
					"route":  "/",
				},
				Func: appHomeWeb.Index,
			},
		},
	},
}
