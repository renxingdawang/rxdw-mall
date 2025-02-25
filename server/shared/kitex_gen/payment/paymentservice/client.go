// Code generated by Kitex v0.11.3. DO NOT EDIT.

package paymentservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	payment "github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/payment"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Charge(ctx context.Context, req *payment.ChargeReq, callOptions ...callopt.Option) (r *payment.ChargeResp, err error)
	CancelPayment(ctx context.Context, req *payment.CancelPaymentReq, callOptions ...callopt.Option) (r *payment.CancelPaymentResp, err error)
	TimedCancelPayment(ctx context.Context, req *payment.TimedCancelPaymentReq, callOptions ...callopt.Option) (r *payment.TimedCancelPaymentResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kPaymentServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kPaymentServiceClient struct {
	*kClient
}

func (p *kPaymentServiceClient) Charge(ctx context.Context, req *payment.ChargeReq, callOptions ...callopt.Option) (r *payment.ChargeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Charge(ctx, req)
}

func (p *kPaymentServiceClient) CancelPayment(ctx context.Context, req *payment.CancelPaymentReq, callOptions ...callopt.Option) (r *payment.CancelPaymentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CancelPayment(ctx, req)
}

func (p *kPaymentServiceClient) TimedCancelPayment(ctx context.Context, req *payment.TimedCancelPaymentReq, callOptions ...callopt.Option) (r *payment.TimedCancelPaymentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TimedCancelPayment(ctx, req)
}
