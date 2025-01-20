package main

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/renxingdawang/rxdw-mall/server/cmd/user/pkg/mysql"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/auth"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	UserMysqlManager
	AuthManager
	EncryptManager
}
type UserMysqlManager interface {
	CreateUser(ctx context.Context, email string, password string) (*mysql.User, error)
	GetUserByEmail(ctx context.Context, Email string) (*mysql.User, error)
}
type AuthManager interface {
	DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq, callOptions ...callopt.Option) (r *auth.DeliveryResp, err error)
	VerifyTokenByRpc(ctx context.Context, req *auth.VerifyTokenReq, callOptions ...callopt.Option) (r *auth.VerifyResp, err error)
	RenewTokenByRpc(ctx context.Context, req *auth.RenewTokenReq, callOptions ...callopt.Option) (r *auth.RenewTokenResp, err error)
}
type EncryptManager interface {
	EncryptPassword(code string) string
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	existingUser, err := s.UserMysqlManager.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("emile already registered")
	}
	newUser, err := s.UserMysqlManager.CreateUser(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	_, err = s.AuthManager.DeliverTokenByRPC(ctx, &auth.DeliverTokenReq{
		UserId: int32(newUser.UserID),
	})
	if err != nil {
		return nil, err
	}
	resp = user.NewRegisterResp()
	resp.SetUserId(int32(newUser.UserID))
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	// TODO: Your code here...
	existingUser, err := s.UserMysqlManager.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, errors.New("user not found")
	}
	SaltPassword := s.EncryptManager.EncryptPassword(req.GetPassword())
	if SaltPassword != existingUser.Password {
		klog.Infof("%s login err", req.GetEmail())
		return nil, errors.New("invalid password")
	}
	// 生成 token
	_, err = s.AuthManager.DeliverTokenByRPC(ctx, &auth.DeliverTokenReq{
		UserId: int32(existingUser.UserID),
	})
	if err != nil {
		return nil, err
	}

	// 返回登录结果
	resp = user.NewLoginResp()
	resp.SetUserId(int32(existingUser.UserID))
	return resp, nil
}
