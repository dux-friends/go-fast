package runtime

import (
	appHomeWeb "dux-project/app/home/web"
	appSystemModels "dux-project/app/system/models"
	"github.com/duxweb/go-fast/annotation"
)

var Annotations = []*annotation.File{
	{
		Name: "dux-project/app/home/web",
		Annotations: []*annotation.Annotation{
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
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.Config{},
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.LogApi{},
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.LogLogin{},
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.LogOperate{},
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.LogVisit{},
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.LogVisitData{},
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.LogVisitSpider{},
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.LogVisitUv{},
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.SystemApi{},
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.SystemRole{},
			},
		},
	},
	{
		Name: "dux-project/app/system/models",
		Annotations: []*annotation.Annotation{
			{
				Name:   "AutoMigrate",
				Params: map[string]any{},
				Func:   appSystemModels.SystemUser{},
			},
		},
	},
}
