package user

import "context"

type UserRepository interface {
	FindUserByUsername(ctx *context.Context, username string) (user User, err error)
	CreateUser(ctx *context.Context, user *User) (err error)
}

type UserTokenRepository interface {
	StoreToken(ctx *context.Context, userToken *UserToken) (err error)
}
