package mysql

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type PaymentLog struct {
	Base
	UserId        int32     `json:"user_id"`
	OrderId       string    `json:"order_id"`
	TransactionId string    `json:"transaction_id"`
	Amount        float64   `json:"amount"`
	PayAt         time.Time `json:"pay_at"`
	Status        string    `json:"status"` // 新增状态字段：created, paid, cancelled
}
type PaymentLogMysqlManager struct {
	db *gorm.DB
}

func NewPaymentLogMysqlManager(db *gorm.DB) *PaymentLogMysqlManager {
	m := db.Migrator()
	if !m.HasTable(&PaymentLog{}) {
		if err := m.CreateTable(&PaymentLog{}); err != nil {
			panic(err)
		}
	}
	return &PaymentLogMysqlManager{
		db: db,
	}
}
func (p *PaymentLog) TableName() string {
	return "payment"
}

func (p *PaymentLogMysqlManager) CreatePaymentLog(ctx context.Context, payment *PaymentLog) error {
	return p.db.WithContext(ctx).Model(&PaymentLog{}).Create(payment).Error
}

func (p *PaymentLogMysqlManager) GetPaymentLogByOrderID(ctx context.Context, OrderId string) (*PaymentLog, error) {
	var payment PaymentLog
	err := p.db.WithContext(ctx).Where("order_id = ?", OrderId).First(&payment).Error
	return &payment, err
}

func (p *PaymentLogMysqlManager) UpdatePaymentLog(ctx context.Context, payment *PaymentLog) error {
	return p.db.WithContext(ctx).Save(payment).Error
}

func (p *PaymentLogMysqlManager) GetPaymentLogByOrderIDAndUserID(ctx context.Context, orderID string, userID int32) (*PaymentLog, error) {
	var payment PaymentLog
	err := p.db.WithContext(ctx).
		Where("order_id = ? AND user_id = ?", orderID, userID).
		First(&payment).
		Error
	return &payment, err
}
func (p *PaymentLogMysqlManager) GetPaymentLogByTransactionID(ctx context.Context, TransactionID string) (*PaymentLog, error) {
	var payment PaymentLog
	err := p.db.WithContext(ctx).Where("transaction_id = ?", TransactionID).First(&payment).Error
	return &payment, err
}
