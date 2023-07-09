package models

import (
	"time"

	enUser "github.com/IbnAnjung/movie_fest/entity/users"
)

type UserToken struct {
	ID        int64  `gorm:"column:id;<-:false"`
	UserID    int64  `gorm:"column:user_id"`
	TokenID   string `gorm:"-"`
	Token     string `gorm:"column:token"`
	ExpiresAt string `gorm:"-"`
}

func (UserToken) TableName() string {
	return "user_tokens"
}

func (m UserToken) ToEntity(en *enUser.UserToken) error {
	expireAt, _ := time.Parse("2006-01-02 15:04:05", m.ExpiresAt)

	en.ID = m.ID
	en.UserID = m.UserID
	en.Token.ID = m.TokenID
	en.Token.Token = m.Token
	en.Token.ExpiresAt = expireAt
	return nil
}

func (m *UserToken) FillFromEntity(en enUser.UserToken) {
	m.ID = en.ID
	m.UserID = en.UserID
	m.TokenID = en.Token.ID
	m.Token = en.Token.Token
	m.ExpiresAt = en.Token.ExpiresAt.Format("2006-01-02 15:04:05")
}
