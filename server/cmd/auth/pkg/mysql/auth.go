package mysql

import (
	"context"
	"github.com/renxingdawang/rxdw-mall/server/shared/errno"
	"gorm.io/gorm"
	"time"
)

type Token struct {
	TokenID   int32     `gorm:"column:token_id;primaryKey;autoIncrement:true" json:"token_id"`
	UserID    int32     `gorm:"column:user_id" json:"user_id"`
	Token     string    `gorm:"column:token;not null" json:"token"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	ExpiredAt time.Time `gorm:"column:expired_at" json:"expired_at"`
}

type AuthManager struct {
	db *gorm.DB
}

func (m *AuthManager) CreateToken(ctx context.Context, token *Token) (*Token, error) {
	if token.Token == "" {
		return nil, errno.AuthSrvErr.WithMessage("Token is nil")
	}
	//err := m.db.withContext(ctx).Create(&token).Error
	err := m.db.WithContext(ctx).Create(&token).Error
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (m *AuthManager) VerifyToken(ctx context.Context, token *Token) (bool, error) {
	if token.Token == "" {
		return false, errno.AuthSrvErr.WithMessage("Token is nil")
	}
	//查询token是否存在 存在返回true
	var count int64
	err := m.db.WithContext(ctx).Model(&Token{}).Where("token=?", token.Token).Count(&count).Error
	if err != nil {
		return false, errno.AuthSrvErr.WithMessage("Get auth count error")
	}

	return count > 0, nil
}
