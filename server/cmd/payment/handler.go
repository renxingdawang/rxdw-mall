package main

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"github.com/renxingdawang/rxdw-mall/server/cmd/payment/pkg/mq"
	"github.com/renxingdawang/rxdw-mall/server/cmd/payment/pkg/mysql"
	"github.com/renxingdawang/rxdw-mall/server/shared/consts"
	"github.com/renxingdawang/rxdw-mall/server/shared/errno"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/order"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/payment"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/saga"
	"strconv"
	"time"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct {
	PaymentLogMysqlManager
	OrderManager
	RabbitMQ
}
type PaymentLogMysqlManager interface {
	CreatePaymentLog(ctx context.Context, payment *mysql.PaymentLog) error
	GetPaymentLogByOrderID(ctx context.Context, OrderId string) (*mysql.PaymentLog, error)
	UpdatePaymentLog(ctx context.Context, payment *mysql.PaymentLog) error
	GetPaymentLogByOrderIDAndUserID(ctx context.Context, orderID string, userID int32) (*mysql.PaymentLog, error)
	GetPaymentLogByTransactionID(ctx context.Context, TransactionID string) (*mysql.PaymentLog, error)
}
type OrderManager interface {
	PlaceOrder(ctx context.Context, req *order.PlaceOrderReq, callOptions ...callopt.Option) (r *order.PlaceOrderResp, err error)
	ListOrder(ctx context.Context, req *order.ListOrderReq, callOptions ...callopt.Option) (r *order.ListOrderResp, err error)
	MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq, callOptions ...callopt.Option) (r *order.MarkOrderPaidResp, err error)
	CancelOrder(ctx context.Context, req *order.CancelOrderReq, callOptions ...callopt.Option) (r *order.CancelOrderResp, err error)
}
type RabbitMQ interface {
	PublishWithDelay(queue string, body []byte, delayMs int) error
	Publish(queue string, event []byte) error
	Consume(queue string, handler func(string)) error
}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	card := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}
	err = card.Validate(true)
	if err != nil {
		return nil, kerrors.NewBizStatusError(400, err.Error())
	}
	translationId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	err = s.PaymentLogMysqlManager.CreatePaymentLog(ctx, &mysql.PaymentLog{
		Base:          mysql.Base{},
		UserId:        req.GetUserId(),
		OrderId:       req.GetOrderId(),
		TransactionId: translationId.String(),
		Amount:        req.GetAmount(),
		PayAt:         time.Now(),
		Status:        string(consts.PayStatePaid),
	})
	if err != nil {
		return nil, err
	}
	return &payment.ChargeResp{TransactionId: translationId.String()}, nil
}

// CancelPayment implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) CancelPayment(ctx context.Context, req *payment.CancelPaymentReq) (resp *payment.CancelPaymentResp, err error) {
	// 1. 查询支付记录（通过 OrderId 和 UserId）
	paymentLog, err := s.PaymentLogMysqlManager.GetPaymentLogByOrderIDAndUserID(ctx, req.GetOrderId(), req.GetUsrId())
	if err != nil {
		return nil, errno.PaymentSrvErr.WithMessage("payment not found")
	}
	// 2. 检查状态是否允许取消
	if paymentLog.Status != string(consts.PayStateCreated) {
		return nil, errno.PaymentSrvErr.WithMessage("payment cannot be cancelled")
	}

	// 3. 更新状态为已取消
	paymentLog.Status = string(consts.PayStateCanceled)
	if err := s.PaymentLogMysqlManager.UpdatePaymentLog(ctx, paymentLog); err != nil {
		return nil, err
	}
	//// 4. 调用 Order 服务取消订单
	//_, err = s.OrderManager.CancelOrder(ctx, &order.CancelOrderReq{
	//	UserId:  req.GetUsrId(),
	//	OrderId: paymentLog.OrderId,
	//})
	//if err != nil {
	//	klog.Errorf("Failed to cancel order: %v", err)
	//}
	event := &saga.PaymentCancelledEvent{
		OrderId:       req.GetOrderId(),
		UserId:        req.GetUsrId(),
		TransactionId: paymentLog.TransactionId,
	}
	eventBytes, _ := json.Marshal(event)
	err = s.RabbitMQ.Publish(
		consts.PaymentCancelledQueue,
		eventBytes,
	)
	if err != nil {
		paymentLog.Status = string(consts.PayStatePaid)
		_ = s.PaymentLogMysqlManager.UpdatePaymentLog(ctx, paymentLog)
		return nil, errno.PaymentSrvErr.WithMessage("can't cancel order by mq")
	}

	return &payment.CancelPaymentResp{Success: true, TransactionId: paymentLog.TransactionId}, nil
}

// TimedCancelPayment implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) TimedCancelPayment(ctx context.Context, req *payment.TimedCancelPaymentReq) (resp *payment.TimedCancelPaymentResp, err error) {
	// 将 TransactionID 发送到延迟队列（30分钟后过期）
	err = s.RabbitMQ.PublishWithDelay(
		consts.DelayQueue,
		[]byte(req.OrderId),
		30*60*1000, // 30分钟（毫秒）
	)
	if err != nil {
		return nil, errno.PaymentSrvErr.WithMessage("failed to schedule cancellation")
	}
	return &payment.TimedCancelPaymentResp{Success: true}, nil
}
func StartPaymentCancelConsumer(rabbitMQ *mq.RabbitMQ, paymentService *PaymentServiceImpl) {
	// 监听死信队列
	_ = rabbitMQ.Consume(consts.DelayQueue, func(transactionID string) {
		ctx := context.Background()
		// 1. 根据 TransactionID 查询支付记录
		paymentLog, err := paymentService.PaymentLogMysqlManager.GetPaymentLogByTransactionID(ctx, transactionID)
		if err != nil {
			klog.Errorf("no way to query (TransactionID=%s): %v", transactionID, err)
			return
		}
		// 2. 构造正确的 CancelPaymentReq
		req := &payment.CancelPaymentReq{
			OrderId: paymentLog.OrderId,
			UsrId:   paymentLog.UserId,
		}
		// 3. 调用取消接口
		_, _ = paymentService.CancelPayment(ctx, req)
	})
}
func StartPaymentCompensationConsumer(rabbitMQ *mq.RabbitMQ, paymentService *PaymentServiceImpl) {
	_ = rabbitMQ.ConsumeByte(consts.OrderCancelFailedQueue, func(msg []byte) {
		var event saga.OrderCancelFailedEvent
		_ = json.Unmarshal(msg, &event)

		ctx := context.Background()
		// 恢复支付状态（补偿操作）
		paymentLog, _ := paymentService.PaymentLogMysqlManager.GetPaymentLogByOrderIDAndUserID(ctx, event.OrderId, event.UserId)
		paymentLog.Status = "paid" // 恢复为已支付状态
		_ = paymentService.PaymentLogMysqlManager.UpdatePaymentLog(ctx, paymentLog)
	})
}
