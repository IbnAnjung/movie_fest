package models

import (
	enUser "github.com/IbnAnjung/movie_fest/entity/users"
)

type UserToken struct {
	ID      int64  `gorm:"column:id;<-:false"`
	UserID  int64  `gorm:"column:user_id"`
	Token   string `gorm:"column:token"`
	IsBlock bool   `gorm:"column:is_block"`
}

func (UserToken) TableName() string {
	return "user_tokens"
}

func (m UserToken) ToEntity(en *enUser.UserToken) error {
	en.ID = m.ID
	en.UserID = m.UserID
	en.Token = m.Token
	en.IsBlock = m.IsBlock
	return nil
}

func (m *UserToken) FillFromEntity(en enUser.UserToken) {
	m.ID = en.ID
	m.UserID = en.UserID
	m.Token = en.Token
	m.IsBlock = en.IsBlock
}
