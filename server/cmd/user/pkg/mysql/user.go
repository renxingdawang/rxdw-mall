package mysql

import (
	"context"
	"errors"
	"github.com/renxingdawang/rxdw-mall/server/cmd/user/pkg/md5"
	"gorm.io/gorm"
)

type User struct {
	UserID   int    `gorm:"column:user_id;primaryKey;autoIncrement"`        // 对应 user_id 字段，主键且自增
	Email    string `gorm:"column:email;type:varchar(255);unique;not null"` // 对应 email 字段，唯一且不为空
	Password string `gorm:"column:password;type:varchar(255);not null"`     // 对应 password 字段，不为空
	// 可以添加其他用户相关字段
}
type UserMysqlManager struct {
	salt string
	db   *gorm.DB
}

func NewUserMysqlManager(db *gorm.DB, salt string) *UserMysqlManager {
	m := db.Migrator()
	if !m.HasTable(&User{}) {
		if err := m.CreateTable(&User{}); err != nil {
			panic(err)
		}
	}
	return &UserMysqlManager{
		db:   db,
		salt: salt,
	}
}

func (m *UserMysqlManager) CreateUser(ctx context.Context, email string, password string) (*User, error) {
	SaltPassword := md5.Crypt(password, m.salt)
	user := &User{
		Email:    email,
		Password: SaltPassword,
	}
	result := m.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
func (m *UserMysqlManager) GetUserByEmail(ctx context.Context, Email string) (*User, error) {
	var user User
	result := m.db.WithContext(ctx).Where("email=?", Email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在
		}
		return nil, result.Error
	}
	return &user, nil
}
