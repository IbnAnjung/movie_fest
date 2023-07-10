package authentication

import (
	"context"
)

func (uc *authenticationUC) Logout(ctx context.Context, id string) (err error) {
	return uc.userTokenRepository.DeleteToken(&ctx, id)
}
