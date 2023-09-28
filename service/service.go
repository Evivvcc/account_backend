package service

import (
	"account_backend/conf"
	"github.com/gin-gonic/gin"
	"runtime"
)

func InitResource() {
	// 设置核心数
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 初始化配置
	conf.InitConfig("./conf/application.yml")

	// 初始化前置应用

}
func Run(router *gin.Engine) {

}
