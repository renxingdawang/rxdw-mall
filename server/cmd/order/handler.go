package main

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/renxingdawang/rxdw-mall/server/cmd/order/pkg/mq"
	"github.com/renxingdawang/rxdw-mall/server/cmd/order/pkg/mysql"
	"github.com/renxingdawang/rxdw-mall/server/shared/consts"
	"github.com/renxingdawang/rxdw-mall/server/shared/errno"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/cart"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/order"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/saga"
	"gorm.io/gorm"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct {
	OrderMysqlManager
}
type OrderMysqlManager interface {
	GetDB() *gorm.DB
	ListOrder(ctx context.Context, userId int32) (orders []mysql.Order, err error)
	GetOrder(ctx context.Context, userId int32, orderId string) (order mysql.Order, err error)
	UpdateOrderState(ctx context.Context, userId int32, orderId string, state consts.OrderState) error
}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	if len(req.OrderItems) == 0 {
		err = errno.OrderSrvErr.WithMessage("OrderItems empty")
		return
	}
	err = s.OrderMysqlManager.GetDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		orderId, _ := uuid.NewUUID()
		o := &mysql.Order{
			OrderId:      orderId.String(),
			OrderState:   consts.OrderStatePlaced,
			UserId:       req.GetUserId(),
			UserCurrency: req.GetUserCurrency(),
			Consignee: mysql.Consignee{
				Email: req.GetEmail(),
			},
		}
		if req.Address != nil {
			a := req.Address
			o.Consignee.Country = a.Country
			o.Consignee.State = a.State
			o.Consignee.City = a.City
			o.Consignee.StreetAddress = a.StreetAddress
		}
		if err := tx.Create(o).Error; err != nil {
			return err
		}
		var itemList []*mysql.OrderItem
		for _, v := range req.OrderItems {
			itemList = append(itemList, &mysql.OrderItem{
				OrderIdRefer: o.OrderId,
				ProductId:    v.Item.ProductId,
				Quantity:     v.Item.Quantity,
				Cost:         float32(v.Cost),
			})
		}
		if err := tx.Create(&itemList).Error; err != nil {
			return err
		}
		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult_{
				OrderId: orderId.String(),
			},
		}
		return nil
	})
	return
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	orders, err := s.OrderMysqlManager.ListOrder(ctx, req.UserId)
	if err != nil {
		klog.Errorf("model.ListOrder.err:%v", err)
		return nil, err
	}
	var list []*order.Order
	for _, v := range orders {
		var items []*order.OrderItem
		for _, v := range v.OrderItems {
			items = append(items, &order.OrderItem{
				Cost: int32(v.Cost),
				Item: &cart.CartItem{
					ProductId: v.ProductId,
					Quantity:  v.Quantity,
				},
			})
		}
		o := &order.Order{
			OrderId:      v.OrderId,
			UserId:       v.UserId,
			UserCurrency: v.UserCurrency,
			Email:        v.Consignee.Email,
			CreatedAt:    int32(v.CreatedAt.Unix()),
			Address: &order.Address{
				Country:       v.Consignee.Country,
				City:          v.Consignee.City,
				StreetAddress: v.Consignee.StreetAddress,
				ZipCode:       v.Consignee.ZipCode,
			},
			OrderItems: items,
		}
		list = append(list, o)
	}
	resp = &order.ListOrderResp{
		Orders: list,
	}
	return
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	if req.UserId == 0 || req.OrderId == "" {
		err = errno.OrderSrvErr.WithMessage("user_id or order_id can not be empty")
		return
	}
	_, err = s.OrderMysqlManager.GetOrder(ctx, req.UserId, req.OrderId)
	if err != nil {
		klog.Errorf("model.ListOrder.err:%v", err)
		return nil, err
	}
	err = s.OrderMysqlManager.UpdateOrderState(ctx, req.UserId, req.OrderId, consts.OrderStatePaid)
	if err != nil {
		klog.Errorf("model.ListOrder.err:%v", err)
		return nil, err
	}
	resp = &order.MarkOrderPaidResp{}
	return
}

// CancelOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CancelOrder(ctx context.Context, req *order.CancelOrderReq) (resp *order.CancelOrderResp, err error) {
	err = s.OrderMysqlManager.UpdateOrderState(ctx, req.GetUserId(), req.GetOrderId(), consts.OrderStateCanceled)
	if err != nil {
		return nil, errno.OrderSrvErr.WithMessage("cancel order error")
	}
	return &order.CancelOrderResp{Success: true}, nil
}
func StartOrderCancelConsumer(rabbitMQ *mq.RabbitMQ, orderService *OrderServiceImpl) {
	_ = rabbitMQ.Consume(consts.PaymentCancelledQueue, func(msg []byte) {
		var event saga.PaymentCancelledEvent
		_ = json.Unmarshal(msg, &event)

		ctx := context.Background()
		// 调用订单取消
		_, err := orderService.CancelOrder(ctx, &order.CancelOrderReq{
			UserId:  event.UserId,
			OrderId: event.OrderId,
		})
		if err != nil {
			// 发布 OrderCancelFailedEvent 触发补偿
			failedEvent := &saga.OrderCancelFailedEvent{
				OrderId:     event.OrderId,
				UserId:      event.UserId,
				ErrorReason: err.Error(),
			}
			failedBytes, _ := json.Marshal(failedEvent)
			_ = rabbitMQ.Publish(consts.OrderCancelFailedQueue, failedBytes)
		}
	})
}
