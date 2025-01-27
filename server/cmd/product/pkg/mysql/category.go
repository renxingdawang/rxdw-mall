package mysql

import (
	"context"
	"gorm.io/gorm"
)

type Category struct {
	Base
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `json:"product" gorm:"many2many:product_category"`
}
type CategoryMysqlManager struct {
	db *gorm.DB
}

func NewCategoryMysqlManager(db *gorm.DB) *CategoryMysqlManager {
	m := db.Migrator()
	if !m.HasTable(&Product{}) {
		if err := m.CreateTable(&Product{}); err != nil {
			panic(err)
		}
	}
	return &CategoryMysqlManager{
		db: db,
	}
}
func (c *CategoryMysqlManager) TableName() string {
	return "category"
}

func (c *CategoryMysqlManager) GetProductsByCategoryName(ctx context.Context, name string) (category []Category, err error) {
	err = c.db.WithContext(ctx).Model(&Category{}).Where(&Category{Name: name}).Preload("Products").Find(&category).Error
	return category, err
}
