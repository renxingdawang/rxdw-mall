package main

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/auth"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

var jwtSecret = []byte("your-secret-key")

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// TODO: Your code here...

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"user_id": req.GetUserId(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}
	resp = auth.NewDeliveryResp()
	resp.SetToken(tokenString)

	return resp, nil
}

// VerifyTokenByRpc implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRpc(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// TODO: Your code here...
	return
}

// RenewTokenByRpc implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) RenewTokenByRpc(ctx context.Context, req *auth.RenewTokenReq) (resp *auth.RenewTokenResp, err error) {
	// TODO: Your code here...
	return
}
