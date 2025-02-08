package service

import (
	"Go-Mall/app/client/biz/utils"
	"Go-Mall/app/client/hertz_gen/client/user"
	"Go-Mall/app/client/infra/rpc"
	"Go-Mall/app/user/biz/model"
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

	token, expire, err := utils.JwtUtil.TokenGenerator(&model.User{Base: model.Base{ID: int(res.UserId)}})
	if err != nil {
		return nil, err
	}
	return &LoginResp{Token: token, Expire: expire.In(time.FixedZone("CST", 8*3600)).Format(time.RFC3339)}, nil
}
