package user

import (
	"context"
	"time"
)

type UserToken struct {
	ID     int64
	UserID int64
	Token  UserTokenDetail
}

type UserTokenDetail struct {
	ID        string
	Token     string
	ExpiresAt time.Time
}

type UserTokenRepository interface {
	StoreToken(ctx *context.Context, userToken *UserToken) (err error)
	GetToken(ctx *context.Context, userToken *UserToken) (err error)
	DeleteToken(ctx *context.Context, id string) (err error)
}
