package main

import (
	"context"
	"github.com/renxingdawang/rxdw-mall/server/cmd/user/pkg/mysql"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	UserMysqlManager
}
type UserMysqlManager interface {
	CreateUser(ctx context.Context, email string, password string) (*mysql.User, error)
	GetUserByEmail(ctx context.Context, Email string) (*mysql.User, error)
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	// TODO: Your code here...
	return
}
