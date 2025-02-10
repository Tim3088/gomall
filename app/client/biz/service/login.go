package service

import (
	"Go-Mall/app/client/hertz_gen/client/user"
	"Go-Mall/app/client/infra/rpc"
	rpcauth "Go-Mall/rpc_gen/kitex_gen/auth"
	rpcuser "Go-Mall/rpc_gen/kitex_gen/user"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

type LoginResp struct {
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Expire string `protobuf:"bytes,2,opt,name=expire,proto3" json:"expire,omitempty"`
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *user.LoginReq) (resp *LoginResp, err error) {
	res, err := rpc.UserClient.Login(h.Context, &rpcuser.LoginReq{Email: req.Email, Password: req.Password})
	if err != nil {
		return nil, err
	}

	resWithToken, err := rpc.AuthClient.DeliverTokenByRPC(h.Context,
		&rpcauth.DeliverTokenReq{UserId: res.UserId, Role: res.Role})
	if err != nil {
		return nil, err
	}

	// 过期时间为24小时
	expire := time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05")
	return &LoginResp{Token: resWithToken.Token, Expire: expire}, nil
}
