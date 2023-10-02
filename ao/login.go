package ao

import (
	"account_backend/request"
	"account_backend/respond"
	"fmt"
)

var LoginSer *LoginService

type LoginService struct{}

func inti() {
	LoginSer = &LoginService{}
}

func (l *LoginService) SayHello(req *request.SayHelloReq) (*respond.SayHelloResp, error) {
	fmt.Printf("Hello %v\n", req.Name)
	return nil, nil
}
