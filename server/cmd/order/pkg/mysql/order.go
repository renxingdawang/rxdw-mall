package mysql

import (
	"context"
	"github.com/renxingdawang/rxdw-mall/server/shared/consts"
	"gorm.io/gorm"
)

type Consignee struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}
type Order struct {
	Base
	OrderId      string `gorm:"uniqueIndex;size:256"`
	UserId       int32
	UserCurrency string
	Consignee    Consignee   `gorm:"embedded"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
	OrderState   consts.OrderState
}
type OrderMysqlManager struct {
	db *gorm.DB
}

func (o *Order) TableName() string {
	return "order"
}
func NewOrderMysqlManager(db *gorm.DB) *OrderMysqlManager {
	m := db.Migrator()
	if !m.HasTable(&Order{}) {
		if err := m.CreateTable(&Order{}); err != nil {
			panic(err)
		}
	}
	return &OrderMysqlManager{
		db: db,
	}
}

func (m *OrderMysqlManager) GetDB() *gorm.DB {
	return m.db
}

func (m *OrderMysqlManager) ListOrder(ctx context.Context, userId int32) (orders []Order, err error) {
	err = m.db.WithContext(ctx).Model(&Order{}).Where(&Order{UserId: userId}).Preload("OrderItems").Find(&orders).Error
	return
}
func (m *OrderMysqlManager) GetOrder(ctx context.Context, userId int32, orderId string) (order Order, err error) {
	err = m.db.WithContext(ctx).Where(&Order{UserId: userId, OrderId: orderId}).First(&order).Error
	return
}

func (m *OrderMysqlManager) UpdateOrderState(ctx context.Context, userId int32, orderId string, state consts.OrderState) error {
	return m.db.WithContext(ctx).Model(&Order{}).Where(&Order{UserId: userId, OrderId: orderId}).Update("order_state", state).Error
}
