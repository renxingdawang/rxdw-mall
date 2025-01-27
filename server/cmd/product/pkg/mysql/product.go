package mysql

import (
	"context"
	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Picture     string     `json:"picture"`
	Price       float32    `json:"price"`
	Categories  []Category `json:"categories" gorm:"many2many:product_category"`
}

func (p Product) TableName() string {
	return "product"
}

type ProductMysqlManager struct {
	db *gorm.DB
}

func NewProductMysqlManager(db *gorm.DB) *ProductMysqlManager {
	m := db.Migrator()
	if !m.HasTable(&Product{}) {
		if err := m.CreateTable(&Product{}); err != nil {
			panic(err)
		}
	}
	return &ProductMysqlManager{
		db: db,
	}
}
func (p *ProductMysqlManager) GetById(ctx context.Context, productId int) (product Product, err error) {
	err = p.db.WithContext(ctx).Model(&Product{}).Where(&Product{Base: Base{ID: productId}}).First(&product).Error
	return
}
func (p *ProductMysqlManager) SearchProduct(ctx context.Context, q string) (product []*Product, err error) {
	err = p.db.WithContext(ctx).Model(&Product{}).Find(&product, "name like ? or description like ?", "%"+q+"%", "%"+q+"%").Error
	return product, err
}
