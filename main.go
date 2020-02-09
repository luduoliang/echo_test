package main

import (
	"echo_test/router"
	"echo_test/config"
	"echo_test/common"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	//默认中间件是否使用
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	//初始化SESSION
	common.InitSession(e)

	//初始化视图目录
	common.InitTemplate(e)

	//初始化静态资源目录
	e.Static("/static", "./static")

	//初始化路由
	router.InitRouter(e)

	//运行服务
	e.Logger.Fatal(e.Start(":" + config.Config.Httpport))
}