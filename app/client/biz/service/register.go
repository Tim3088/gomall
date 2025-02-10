package service

import (
	"Go-Mall/app/client/hertz_gen/client/common"
	"Go-Mall/app/client/hertz_gen/client/user"
	"Go-Mall/app/client/infra/rpc"
	rpcuser "Go-Mall/rpc_gen/kitex_gen/user"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *user.RegisterReq) (resp *common.Empty, err error) {
	_, err = rpc.UserClient.Register(h.Context, &rpcuser.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.Password,
		Role:            req.Role,
	})
	if err != nil {
		return nil, err
	}

	return
}
