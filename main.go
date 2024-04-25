package main

import (
	"dux-project/app/home"
	"dux-project/runtime"
	"embed"
	dux "github.com/duxweb/go-fast"
)

//go:embed views/* app/*/views/*
var ViewsFs embed.FS

func main() {

	// 创建框架服务
	app := dux.New()
	app.SetDatabaseStatus(true)
	app.RegisterAnnotations(runtime.Annotations)
	app.RegisterApp(home.App)

	app.Run()

	//// 设置基础配置
	//server.SetConfigDir("./config/")
	//server.SetDatabaseStatus(true)
	//server.SetRedisStatus(true)
	//server.SetQueueStatus(true)
	//server.SetWebsocketStatus(true)
	//server.SetMongodbStatus(false)
	//
	//// 注册应用模块
	//server.RegisterApp(home.App, admin.New)
	//
	//// 注册框架服务
	//server.RegisterService(ui.New, func(bootstrap *bootstrap.Bootstrap) {
	//	template.Must(core.Tpl.ParseFS(ViewsFs, "views/*", "app/*/views /*"))
	//})
	//
	//// 启动框架服务
	//server.Start()
}
