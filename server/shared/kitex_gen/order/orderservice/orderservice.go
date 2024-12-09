// Code generated by Kitex v0.11.3. DO NOT EDIT.

package orderservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	order "github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/order"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"PlaceOrder": kitex.NewMethodInfo(
		placeOrderHandler,
		newOrderServicePlaceOrderArgs,
		newOrderServicePlaceOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ListOrder": kitex.NewMethodInfo(
		listOrderHandler,
		newOrderServiceListOrderArgs,
		newOrderServiceListOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"MarkOrderPaid": kitex.NewMethodInfo(
		markOrderPaidHandler,
		newOrderServiceMarkOrderPaidArgs,
		newOrderServiceMarkOrderPaidResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CancelOrder": kitex.NewMethodInfo(
		cancelOrderHandler,
		newOrderServiceCancelOrderArgs,
		newOrderServiceCancelOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	orderServiceServiceInfo                = NewServiceInfo()
	orderServiceServiceInfoForClient       = NewServiceInfoForClient()
	orderServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return orderServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return orderServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return orderServiceServiceInfoForClient
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
	serviceName := "OrderService"
	handlerType := (*order.OrderService)(nil)
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
		"PackageName": "order",
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

func placeOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*order.OrderServicePlaceOrderArgs)
	realResult := result.(*order.OrderServicePlaceOrderResult)
	success, err := handler.(order.OrderService).PlaceOrder(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOrderServicePlaceOrderArgs() interface{} {
	return order.NewOrderServicePlaceOrderArgs()
}

func newOrderServicePlaceOrderResult() interface{} {
	return order.NewOrderServicePlaceOrderResult()
}

func listOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*order.OrderServiceListOrderArgs)
	realResult := result.(*order.OrderServiceListOrderResult)
	success, err := handler.(order.OrderService).ListOrder(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOrderServiceListOrderArgs() interface{} {
	return order.NewOrderServiceListOrderArgs()
}

func newOrderServiceListOrderResult() interface{} {
	return order.NewOrderServiceListOrderResult()
}

func markOrderPaidHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*order.OrderServiceMarkOrderPaidArgs)
	realResult := result.(*order.OrderServiceMarkOrderPaidResult)
	success, err := handler.(order.OrderService).MarkOrderPaid(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOrderServiceMarkOrderPaidArgs() interface{} {
	return order.NewOrderServiceMarkOrderPaidArgs()
}

func newOrderServiceMarkOrderPaidResult() interface{} {
	return order.NewOrderServiceMarkOrderPaidResult()
}

func cancelOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*order.OrderServiceCancelOrderArgs)
	realResult := result.(*order.OrderServiceCancelOrderResult)
	success, err := handler.(order.OrderService).CancelOrder(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newOrderServiceCancelOrderArgs() interface{} {
	return order.NewOrderServiceCancelOrderArgs()
}

func newOrderServiceCancelOrderResult() interface{} {
	return order.NewOrderServiceCancelOrderResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (r *order.PlaceOrderResp, err error) {
	var _args order.OrderServicePlaceOrderArgs
	_args.Req = req
	var _result order.OrderServicePlaceOrderResult
	if err = p.c.Call(ctx, "PlaceOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListOrder(ctx context.Context, req *order.ListOrderReq) (r *order.ListOrderResp, err error) {
	var _args order.OrderServiceListOrderArgs
	_args.Req = req
	var _result order.OrderServiceListOrderResult
	if err = p.c.Call(ctx, "ListOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (r *order.MarkOrderPaidResp, err error) {
	var _args order.OrderServiceMarkOrderPaidArgs
	_args.Req = req
	var _result order.OrderServiceMarkOrderPaidResult
	if err = p.c.Call(ctx, "MarkOrderPaid", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CancelOrder(ctx context.Context, req *order.CancelOrderReq) (r *order.CancelOrderResp, err error) {
	var _args order.OrderServiceCancelOrderArgs
	_args.Req = req
	var _result order.OrderServiceCancelOrderResult
	if err = p.c.Call(ctx, "CancelOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
