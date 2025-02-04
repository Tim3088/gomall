package service

import (
	"Go-Mall/app/client/hertz_gen/client/user"
	"Go-Mall/app/client/infra/rpc"
	clientutils "Go-Mall/app/client/utils"
	rpcuser "Go-Mall/rpc_gen/kitex_gen/user"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

type LoginResp struct {
	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *user.LoginReq) (resp *LoginResp, err error) {
	res, err := rpc.UserClient.Login(h.Context, &rpcuser.LoginReq{Email: req.Email, Password: req.Password})
	if err != nil {
		return nil, err
	}

	clientutils.MustHandleError(err)

	if err != nil {
		return nil, err
	}

	return &LoginResp{UserId: res.UserId}, nil
}
