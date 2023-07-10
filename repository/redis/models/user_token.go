package models

import (
	"fmt"
	"time"

	enUser "github.com/IbnAnjung/movie_fest/entity/users"
)

type UserToken struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	TokenID   string `json:"token_id"`
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}

func (m UserToken) Key() string {
	return fmt.Sprintf("user_token:%s", m.TokenID)
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
