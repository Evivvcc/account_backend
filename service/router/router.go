package router

import "github.com/gin-gonic/gin"

/*
InitRouter 路由初始化
*/
func InitRouter() *gin.Engine {
	r := gin.New()

	bizGroup := r.Group("/api/v1")
	// 登录相关路由
	addLoginRouter(bizGroup)

	return r
}

func addLoginRouter(bizGroup *gin.RouterGroup) {

}
