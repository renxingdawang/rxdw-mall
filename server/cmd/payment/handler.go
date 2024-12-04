package main

import (
	"context"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/payment"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// TODO: Your code here...
	return
}

// CancelPayment implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) CancelPayment(ctx context.Context, req *payment.CancelPaymentReq) (resp *payment.CancelPaymentResp, err error) {
	// TODO: Your code here...
	return
}

// TimedCancelPayment implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) TimedCancelPayment(ctx context.Context, req *payment.TimedCancelPaymentReq) (resp *payment.TimedCancelPaymentResp, err error) {
	// TODO: Your code here...
	return
}
