package authentication

import "context"

type AuthenticationUsecase interface {
	RegisterUser(ctx context.Context, input Register) (output RegisteredUser, err error)
	Login(ctx context.Context, input Login) (output LogedinUser, err error)
}
