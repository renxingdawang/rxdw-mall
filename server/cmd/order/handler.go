package main

import (
	"context"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/order"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// TODO: Your code here...
	return
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// TODO: Your code here...
	return
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// TODO: Your code here...
	return
}

// CancelOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CancelOrder(ctx context.Context, req *order.CancelOrderReq) (resp *order.CancelOrderResp, err error) {
	// TODO: Your code here...
	return
}
