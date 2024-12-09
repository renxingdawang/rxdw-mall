// Code generated by Kitex v0.11.3. DO NOT EDIT.

package cartservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	cart "github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/cart"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AddItem(ctx context.Context, req *cart.AddItemReq, callOptions ...callopt.Option) (r *cart.AddItemResp, err error)
	GetCart(ctx context.Context, req *cart.GetCartReq, callOptions ...callopt.Option) (r *cart.GetCartResp, err error)
	EmptyCart(ctx context.Context, req *cart.EmptyCartReq, callOptions ...callopt.Option) (r *cart.EmptyCartResp, err error)
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
	return &kCartServiceClient{
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

type kCartServiceClient struct {
	*kClient
}

func (p *kCartServiceClient) AddItem(ctx context.Context, req *cart.AddItemReq, callOptions ...callopt.Option) (r *cart.AddItemResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddItem(ctx, req)
}

func (p *kCartServiceClient) GetCart(ctx context.Context, req *cart.GetCartReq, callOptions ...callopt.Option) (r *cart.GetCartResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCart(ctx, req)
}

func (p *kCartServiceClient) EmptyCart(ctx context.Context, req *cart.EmptyCartReq, callOptions ...callopt.Option) (r *cart.EmptyCartResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EmptyCart(ctx, req)
}
