package mysql

import (
	"time"
)

type Base struct {
	ID       int32 `gorm:"primarykey"`
	CreateAt time.Time
	UpdateAt time.Time
}
