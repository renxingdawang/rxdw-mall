package mysql

import (
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

func (m *AuthManager) CreateToken(token *Token) (*Token, error) {
	if token.Token == "" {
		return nil, errno.AuthorizeFail.WithMessage("Token is nil")
	}
	err := m.db.Create(&token).Error
	if err != nil {
		return nil, err
	}

	return token, nil
}
