package service

import (
	"Go-Mall/app/user/biz/dal/mysql"
	"Go-Mall/app/user/biz/model"
	"Go-Mall/rpc_gen/kitex_gen/user"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" || req.ConfirmPassword == "" || req.Role == 0 {
		return nil, errors.New("邮箱或密码不能为空")
	}
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("密码必须与确认密码相同")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(hashedPassword),
		Role:           req.Role,
	}
	if err = model.Create(mysql.DB, s.ctx, newUser); err != nil {
		return nil, err
	}

	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
