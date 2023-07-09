package authentication

import "context"

type AuthenticationUsecase interface {
	RegisterUser(ctx context.Context, input Register) (output RegisteredUser, err error)
}
