package main

import (
	"context"
	"github.com/hertz-contrib/paseto"
	"github.com/renxingdawang/rxdw-mall/server/cmd/auth/pkg/mysql"
	"github.com/renxingdawang/rxdw-mall/server/shared/consts"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/auth"
	"time"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct {
	AuthManger
	TokenGenerator
}
type AuthManger interface {
	CreateToken(ctx context.Context, token *mysql.Token) (*mysql.Token, error)
	VerifyToken(ctx context.Context, token string) (bool, error)
	RenewToken(ctx context.Context, token string) error
}
type TokenGenerator interface {
	CreateToken(claims *paseto.StandardClaims) (token string, err error)
}

var jwtSecret = []byte("your-secret-key")

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {

	//token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
	//	"user_id": req.GetUserId(),
	//})
	//tokenString, err := token.SignedString(jwtSecret)
	//if err != nil {
	//	return nil, err
	//}
	tokenString, err := s.TokenGenerator.CreateToken(&paseto.StandardClaims{
		ID:        string(req.GetUserId()),
		Issuer:    consts.Issuer,
		Audience:  consts.User,
		IssuedAt:  time.Now(),
		NotBefore: time.Now(),
		ExpiredAt: time.Now().Add(consts.ThirtyDays),
	})
	resp = auth.NewDeliveryResp()
	resp.SetToken(tokenString)
	//将token存入mysql token table
	_, _ = s.AuthManger.CreateToken(ctx, &mysql.Token{
		UserID:    req.GetUserId(),
		Token:     tokenString,
		ExpiredAt: time.Now().Add(time.Hour),
	})
	return resp, nil
}

// VerifyTokenByRpc implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRpc(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {

	isExist, _ := s.AuthManger.VerifyToken(ctx, req.GetToken())
	resp = auth.NewVerifyResp()
	resp.SetRes(isExist)
	return resp, nil
}

// RenewTokenByRpc implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) RenewTokenByRpc(ctx context.Context, req *auth.RenewTokenReq) (resp *auth.RenewTokenResp, err error) {
	err = s.AuthManger.RenewToken(ctx, req.GetToken())
	resp = auth.NewRenewTokenResp()
	if err != nil {
		resp.SetRes(false)
		return resp, err
	}
	resp.SetRes(true)
	return resp, err
}
