package main

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/renxingdawang/rxdw-mall/server/shared/errno"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/cart"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/checkout"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/order"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/payment"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/product"
	"strconv"
)

// CheckoutServiceImpl implements the last service interface defined in the IDL.
type CheckoutServiceImpl struct {
	CartManager
	OrderManager
	PaymentManager
	ProductManager
}
type CartManager interface {
	AddItem(ctx context.Context, req *cart.AddItemReq, callOptions ...callopt.Option) (r *cart.AddItemResp, err error)
	GetCart(ctx context.Context, req *cart.GetCartReq, callOptions ...callopt.Option) (r *cart.GetCartResp, err error)
	EmptyCart(ctx context.Context, req *cart.EmptyCartReq, callOptions ...callopt.Option) (r *cart.EmptyCartResp, err error)
}
type OrderManager interface {
	PlaceOrder(ctx context.Context, req *order.PlaceOrderReq, callOptions ...callopt.Option) (r *order.PlaceOrderResp, err error)
	ListOrder(ctx context.Context, req *order.ListOrderReq, callOptions ...callopt.Option) (r *order.ListOrderResp, err error)
	MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq, callOptions ...callopt.Option) (r *order.MarkOrderPaidResp, err error)
	CancelOrder(ctx context.Context, req *order.CancelOrderReq, callOptions ...callopt.Option) (r *order.CancelOrderResp, err error)
}
type PaymentManager interface {
	Charge(ctx context.Context, req *payment.ChargeReq, callOptions ...callopt.Option) (r *payment.ChargeResp, err error)
	CancelPayment(ctx context.Context, req *payment.CancelPaymentReq, callOptions ...callopt.Option) (r *payment.CancelPaymentResp, err error)
	TimedCancelPayment(ctx context.Context, req *payment.TimedCancelPaymentReq, callOptions ...callopt.Option) (r *payment.TimedCancelPaymentResp, err error)
}
type ProductManager interface {
	ListProducts(ctx context.Context, req *product.ListProductsReq, callOptions ...callopt.Option) (r *product.ListProductsResp, err error)
	GetProduct(ctx context.Context, req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error)
	SearchProducts(ctx context.Context, req *product.SearchProductsReq, callOptions ...callopt.Option) (r *product.SearchProductsResp, err error)
}

/*
Checkout
1. get cart
2. calculate cart
3. create order
4. empty cart
5. pay
6. change order result
7. finish
*/
// Checkout implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) Checkout(ctx context.Context, req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	cartResult, err := s.CartManager.GetCart(ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err)
		return
	}
	if cartResult == nil || cartResult.Cart == nil || len(cartResult.Cart.Items) == 0 {
		err = errno.CheckOutSrvErr.WithMessage("cart is empty")
		return
	}
	var (
		oi    []*order.OrderItem
		total float64
	)
	for _, cartItem := range cartResult.Cart.Items {
		productResp, resultErr := s.ProductManager.GetProduct(ctx, &product.GetProductReq{Id: cartItem.ProductId})
		if resultErr != nil {
			klog.Error(resultErr)
			err = resultErr
			return
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		cost := p.Price * float64(cartItem.Quantity)
		total += cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{ProductId: cartItem.ProductId, Quantity: cartItem.Quantity},
			Cost: int32(cost),
		})
	}
	// create order
	orderReq := &order.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: "USD",
		OrderItems:   oi,
		Email:        req.Email,
	}
	if req.Address != nil {
		addr := req.Address
		zipCodeInt, _ := strconv.Atoi(string(addr.ZipCode))
		orderReq.Address = &order.Address{
			StreetAddress: addr.StreetAddress,
			City:          addr.City,
			Country:       addr.Country,
			State:         addr.State,
			ZipCode:       int32(zipCodeInt),
		}
	}
	orderResult, err := s.OrderManager.PlaceOrder(ctx, orderReq)
	if err != nil {
		err = errno.CheckOutSrvErr.WithMessage("placeOrder error")
		return
	}
	klog.Info("orderResult", orderResult)
	// empty cart
	emptyResult, err := s.CartManager.EmptyCart(ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		err = errno.CheckOutSrvErr.WithMessage("EmptyCart error")
		return
	}
	klog.Info(emptyResult)

	// charge
	var orderId string
	if orderResult != nil || orderResult.Order != nil {
		orderId = orderResult.Order.OrderId
	}
	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
		},
	}
	paymentResult, err := s.PaymentManager.Charge(ctx, payReq)
	if err != nil {
		err = errno.CheckOutSrvErr.WithMessage("Charge error")
		return
	}
	klog.Info(paymentResult)
	// change order state
	klog.Info(orderResult)
	_, err = s.OrderManager.MarkOrderPaid(ctx, &order.MarkOrderPaidReq{UserId: req.UserId, OrderId: orderId})
	if err != nil {
		klog.Error(err)
		return
	}

	resp = &checkout.CheckoutResp{
		OrderId:     orderId,
		Transaction: paymentResult.TransactionId,
	}
	return
}
