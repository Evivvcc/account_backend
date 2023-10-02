package service

import (
	"account_backend/conf/config"
	"account_backend/integration/storage/mysql"
	"github.com/gin-gonic/gin"
	"runtime"
)

func InitResource() {
	// 设置核心数
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 加载配置文件
	config.InitConfig("./conf/application.yml")

	//mysql
	mysql.InitMysql()

	// redis
	//redis.InitRedis()

	// rsa
	//initRsa()

	// 注册到北极星

	// 注册链路追踪供应器

}

func RunServer(router *gin.Engine) {

}
