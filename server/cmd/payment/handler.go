package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"github.com/renxingdawang/rxdw-mall/server/cmd/payment/pkg/mysql"
	"github.com/renxingdawang/rxdw-mall/server/shared/kitex_gen/payment"
	"strconv"
	"time"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct {
	PaymentLogMysqlManager
}
type PaymentLogMysqlManager interface {
	CreatePaymentLog(ctx context.Context, payment *mysql.PaymentLog) error
}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// TODO: Your code here...
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
		UserId:        uint32(req.GetUserId()),
		OrderId:       req.GetOrderId(),
		TransactionId: translationId.String(),
		Amount:        req.GetAmount(),
		PayAt:         time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &payment.ChargeResp{TransactionId: translationId.String()}, nil
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
