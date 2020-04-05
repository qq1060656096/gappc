package generate

const apiRoutersTemplate = `
package routers

import (
	"github.com/gin-gonic/gin"
	"{{projectName}}/app/http/controller/api/v1/demo"
)

// ApiV1 接口v1版本路由
func ApiV1(engine *gin.Engine) {

	g := engine.Group("/api/v1")
	{
		// 简单 get post put delete请求示例
		g.POST("/demo/demoApi/post", demo.DemoApiCreate)
		g.DELETE("/demo/demoApi/delete", demo.DemoApiDelete)
		g.PUT("/demo/demoApi/put", demo.DemoApiUpdate)
		g.GET("/demo/demoApi/get", demo.DemoApiGet)
		g.GET("/demo/demoApi/get/detail", demo.DemoApiGetDetail)

		// 中间件
		//g.Use(middleware.DemoAuth()).GET("/demo/middleware", demo.Middleware)
	}
}

// ApiV2 接口v2版本路由
func ApiV2(engine *gin.Engine) {
	//g := engine.Group("/api/v2")
	//{
	//	g.GET("/demo/json")
	//}
}

// ApiV3 接口v3版本路由
func ApiV3(engine *gin.Engine) {
	//g := engine.Group("/api/v3")
	//{
	//	g.GET("/demo/json")
	//}
}
`


const appRoutersTemplate = `package routers

import (
	"github.com/gin-gonic/gin"
	"{{projectName}}/app/http/controller/app/v1/demo"
)

// AppV1 应用v1版本路由
func AppV1(engine *gin.Engine) {
	g := engine.Group("/app/v1")
	{
		// 简单 get post put delete请求示例
		g.POST("/demo/DemoApp/post", demo.DemoAppCreate)
		g.DELETE("/demo/DemoApp/delete", demo.DemoAppDelete)
		g.PUT("/demo/DemoApp/put", demo.DemoAppUpdate)
		g.GET("/demo/DemoApp/get", demo.DemoAppGet)
		g.GET("/demo/DemoApp/get/detail", demo.DemoAppGetDetail)
	}
}

// AppV2 应用v2版本路由
func AppV2(engine *gin.Engine) {
	//g := engine.Group("/app/v2")
	//{
	//	g.GET("/demo/json")
	//}
}

// AppV3 应用v3版本路由
func AppV3(engine *gin.Engine) {
	//g := engine.Group("/app/v3")
	//{
	//	g.GET("/demo/json")
	//}
}

`

const routersTemplate = `package routers

import "github.com/gin-gonic/gin"

// InitRouter 初始化路由
func InitRouter(engine *gin.Engine) *gin.Engine {
	// api 接口路由配置
	ApiV1(engine)
	ApiV2(engine)
	ApiV3(engine)

	// app 应用路由配置
	AppV1(engine)
	AppV2(engine)
	AppV3(engine)
	return engine
}

`