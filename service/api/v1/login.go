package v1

import (
	"account_backend/ao"
	"account_backend/request"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SayHello(ctx *gin.Context) {
	var req request.SayHelloReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Errorf("%s\n", err)
		return
	}
	_, err := ao.LoginSer.SayHello(&req)
	if err != nil {
		return
	}
}
