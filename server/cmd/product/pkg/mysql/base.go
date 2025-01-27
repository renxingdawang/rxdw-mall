package mysql

import (
	"time"
)

type Base struct {
	ID       int `gorm:"primarykey"`
	CreateAt time.Time
	UpdateAt time.Time
}
