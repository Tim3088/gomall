package service

import (
	"Go-Mall/app/user/biz/dal/mysql"
	user "Go-Mall/rpc_gen/kitex_gen/user"
	"context"
	"github.com/joho/godotenv"
	"testing"
)

func TestLogin_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewLoginService(ctx)
	// init req and assert value

	req := &user.LoginReq{
		Email:    "test@test.com",
		Password: "123456",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
