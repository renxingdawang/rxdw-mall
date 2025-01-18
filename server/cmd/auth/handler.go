package main

import (
	"context"
	"errors"
	"github.com/hertz-contrib/paseto"
	"github.com/renxingdawang/rxdw-mall/server/shared/consts"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/auth"
	"time"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct {
	TokenGenerator
	AuthRedisManager
}
type TokenGenerator interface {
	CreateToken(claims *paseto.StandardClaims) (token string, err error)
	ParseToken(token string) (int32, error)
}
type AuthRedisManager interface {
	StoreToken(ctx context.Context, UserId int32, token string, expiration time.Duration) error
	GetToken(ctx context.Context, UserId int32) (string, error)
}

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
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
	expiration := consts.ThirtyDays
	err = s.AuthRedisManager.StoreToken(ctx, req.GetUserId(), tokenString, expiration)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// VerifyTokenByRpc implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRpc(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	resp = auth.NewVerifyResp()
	userID, err := s.TokenGenerator.ParseToken(req.GetToken())
	if err != nil {
		resp.SetRes(false)
		return resp, err
	}
	storedToken, err := s.AuthRedisManager.GetToken(ctx, userID)
	if err != nil {
		resp.SetRes(false)
		return resp, err
	}
	resp.SetRes(storedToken == req.GetToken())
	return resp, nil
}

// RenewTokenByRpc implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) RenewTokenByRpc(ctx context.Context, req *auth.RenewTokenReq) (resp *auth.RenewTokenResp, err error) {
	resp = auth.NewRenewTokenResp()
	userID, err := s.TokenGenerator.ParseToken(req.GetToken())
	if err != nil {
		resp.SetRes(false)
		return resp, err
	}
	storedToken, err := s.AuthRedisManager.GetToken(ctx, userID)
	if err != nil {
		resp.SetRes(false)
		return resp, err
	}
	// 验证 token 是否一致
	if storedToken != req.GetToken() {
		return nil, errors.New("invalid token")
	}

	// 生成新的 token
	newTokenString, err := s.TokenGenerator.CreateToken(&paseto.StandardClaims{
		ID:        string(userID),
		Issuer:    consts.Issuer,
		Audience:  consts.User,
		IssuedAt:  time.Now(),
		NotBefore: time.Now(),
		ExpiredAt: time.Now().Add(consts.ThirtyDays),
	})
	if err != nil {
		return nil, err
	}
	// 存储新的 token 到 Redis
	expiration := consts.ThirtyDays
	err = s.AuthRedisManager.StoreToken(ctx, userID, newTokenString, expiration)
	if err != nil {
		return nil, err
	}
	resp = auth.NewRenewTokenResp()
	resp.SetRes(true)
	return resp, nil
}
