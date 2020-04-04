package generate

const mainTemplate = `package main

import (
	"{{projectName}}/bootstrap"
	"{{projectName}}/routers"
	"github.com/gin-gonic/gin"
	"github.com/qq1060656096/go-common"
)

func main() {
	app := bootstrap.App()
	gin.DefaultWriter = app.LogWriter
	engine := gin.Default()
	engine.LoadHTMLGlob("resources/themes/***/***/*")
	// 初始化路由
	routers.InitRouter(engine)
	addr := common.OsEnvManager.Get("ADDR")
	engine.Run(addr)
}
`