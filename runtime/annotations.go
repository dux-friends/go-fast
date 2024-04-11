package main

import (
	appHomeWeb "dux-project/app/home/web"
	appTestWeb "dux-project/app/test/web"
)

type Annotation struct {
	Name   string
	Params map[string]any
	Func   any
}

type File struct {
	Name        string
	Annotations []*Annotation
}

var Annotations = []*File{
	{
		Name: "dux-project/app/home/web",
		Annotations: []*Annotation{
			{
				Name: "RouteGroup",
				Params: map[string]any{
					"name":  "xxxGroup",
					"route": "/list",
				},
			},
			{
				Name: "Route",
				Params: map[string]any{
					"name":  "xxx",
					"route": "/",
				},
				Func: appHomeWeb.Index,
			},
		},
	},
	{
		Name: "dux-project/app/test/web",
		Annotations: []*Annotation{
			{
				Name: "RouteGroup",
				Params: map[string]any{
					"name":  "testGroup",
					"route": "/list",
				},
			},
			{
				Name: "Route",
				Params: map[string]any{
					"name":  "test",
					"route": "/test",
				},
				Func: appTestWeb.Index,
			},
		},
	},
}
