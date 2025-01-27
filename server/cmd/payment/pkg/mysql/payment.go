package mysql

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type PaymentLog struct {
	Base
	UserId        uint32    `json:"user_id"`
	OrderId       string    `json:"order_id"`
	TransactionId string    `json:"transaction_id"`
	Amount        float64   `json:"amount"`
	PayAt         time.Time `json:"pay_at"`
}
type PaymentLogMysqlManager struct {
	db *gorm.DB
}

func NewUserMysqlManager(db *gorm.DB) *PaymentLogMysqlManager {
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
