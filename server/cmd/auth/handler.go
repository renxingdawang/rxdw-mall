package main

import (
	"context"
	"github.com/rxdw-mall/server/shared/kitex_gen/auth"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// TODO: Your code here...
	return
}

// VerifyTokenByRpc implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRpc(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// TODO: Your code here...
	return
}
