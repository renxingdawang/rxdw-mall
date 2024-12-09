// Code generated by Kitex v0.11.3. DO NOT EDIT.

package cartservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	cart "github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/cart"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"AddItem": kitex.NewMethodInfo(
		addItemHandler,
		newCartServiceAddItemArgs,
		newCartServiceAddItemResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetCart": kitex.NewMethodInfo(
		getCartHandler,
		newCartServiceGetCartArgs,
		newCartServiceGetCartResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"EmptyCart": kitex.NewMethodInfo(
		emptyCartHandler,
		newCartServiceEmptyCartArgs,
		newCartServiceEmptyCartResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	cartServiceServiceInfo                = NewServiceInfo()
	cartServiceServiceInfoForClient       = NewServiceInfoForClient()
	cartServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return cartServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return cartServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return cartServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "CartService"
	handlerType := (*cart.CartService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "cart",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.11.3",
		Extra:           extra,
	}
	return svcInfo
}

func addItemHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*cart.CartServiceAddItemArgs)
	realResult := result.(*cart.CartServiceAddItemResult)
	success, err := handler.(cart.CartService).AddItem(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCartServiceAddItemArgs() interface{} {
	return cart.NewCartServiceAddItemArgs()
}

func newCartServiceAddItemResult() interface{} {
	return cart.NewCartServiceAddItemResult()
}

func getCartHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*cart.CartServiceGetCartArgs)
	realResult := result.(*cart.CartServiceGetCartResult)
	success, err := handler.(cart.CartService).GetCart(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCartServiceGetCartArgs() interface{} {
	return cart.NewCartServiceGetCartArgs()
}

func newCartServiceGetCartResult() interface{} {
	return cart.NewCartServiceGetCartResult()
}

func emptyCartHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*cart.CartServiceEmptyCartArgs)
	realResult := result.(*cart.CartServiceEmptyCartResult)
	success, err := handler.(cart.CartService).EmptyCart(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCartServiceEmptyCartArgs() interface{} {
	return cart.NewCartServiceEmptyCartArgs()
}

func newCartServiceEmptyCartResult() interface{} {
	return cart.NewCartServiceEmptyCartResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) AddItem(ctx context.Context, req *cart.AddItemReq) (r *cart.AddItemResp, err error) {
	var _args cart.CartServiceAddItemArgs
	_args.Req = req
	var _result cart.CartServiceAddItemResult
	if err = p.c.Call(ctx, "AddItem", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCart(ctx context.Context, req *cart.GetCartReq) (r *cart.GetCartResp, err error) {
	var _args cart.CartServiceGetCartArgs
	_args.Req = req
	var _result cart.CartServiceGetCartResult
	if err = p.c.Call(ctx, "GetCart", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EmptyCart(ctx context.Context, req *cart.EmptyCartReq) (r *cart.EmptyCartResp, err error) {
	var _args cart.CartServiceEmptyCartArgs
	_args.Req = req
	var _result cart.CartServiceEmptyCartResult
	if err = p.c.Call(ctx, "EmptyCart", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
