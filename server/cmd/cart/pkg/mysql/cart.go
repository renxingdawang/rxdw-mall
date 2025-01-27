package mysql

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type Cart struct {
	Base
	UserId    int32 `json:"user_id"`
	ProductId int32 `json:"product_id"`
	Qty       int32 `json:"qty"`
}

func (c *Cart) TableName() string {
	return "cart"
}

type CartMysqlManager struct {
	db *gorm.DB
}

func NewCartMysqlManager(db *gorm.DB) *CartMysqlManager {
	m := db.Migrator()
	if !m.HasTable(&Cart{}) {
		if err := m.CreateTable(&Cart{}); err != nil {
			panic(err)
		}
	}
	return &CartMysqlManager{
		db: db,
	}
}

func (m *CartMysqlManager) GetCartByUserId(ctx context.Context, userId int32) (cartList []*Cart, err error) {
	err = m.db.Debug().WithContext(ctx).Model(&Cart{}).Find(&cartList, "user_id = ?", userId).Error
	return cartList, err
}

func (m *CartMysqlManager) AddCart(ctx context.Context, c *Cart) error {
	var find Cart
	err := m.db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: c.UserId, ProductId: c.ProductId}).First(&find).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if find.ID != 0 {
		err = m.db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: c.UserId, ProductId: c.ProductId}).UpdateColumn("qty", gorm.Expr("qty+?", c.Qty)).Error
	} else {
		err = m.db.WithContext(ctx).Model(&Cart{}).Create(c).Error
	}
	return err
}

func (m *CartMysqlManager) EmptyCart(ctx context.Context, userId int32) error {
	if userId == 0 {
		return errors.New("user_is is required")
	}
	return m.db.WithContext(ctx).Delete(&Cart{}, "user_id = ?", userId).Error
}
