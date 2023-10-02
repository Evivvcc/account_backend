package main

import (
	"account_backend/service"
	"account_backend/service/router"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	// 设置 gin 模式
	if os.Getenv("env") == "test" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化各种资源
	service.InitResource()

	// 初始化路由
	r := router.InitRouter()

	// 启动服务
	service.RunServer(r)

}
