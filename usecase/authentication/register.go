package authentication

import (
	"context"
	"errors"

	enAuth "github.com/IbnAnjung/movie_fest/entity/authentication"
	enUser "github.com/IbnAnjung/movie_fest/entity/users"

	"github.com/IbnAnjung/movie_fest/utils"
)

type inputValidationObject struct {
	Username        string `validate:"required,min=5,max=50"`
	Password        string `validate:"required,min=5,max=255"`
	ConfirmPassword string `validate:"required,min=5,max=255,eqfield=Password"`
}

func (i *inputValidationObject) set(src enAuth.Register) {
	i.Username = src.Username
	i.Password = src.Password
	i.ConfirmPassword = src.ConfirmPassword
}

func (uc *authenticationUC) RegisterUser(ctx context.Context, input enAuth.Register) (output enAuth.RegisteredUser, err error) {

	iv := inputValidationObject{}
	iv.set(input)

	if err = uc.validator.ValidateStruct(&iv); err != nil {
		return
	}

	oldUser, err := uc.userRepository.FindUserByUsername(&ctx, input.Username)
	if err != nil && !errors.Is(err, utils.DataNotFoundError) {
		return
	}

	if oldUser.ID != 0 {
		e := utils.UnprocessableEntityError
		e.Message = "usename already exists"
		return output, e
	}

	hashPassword, err := uc.crypt.HashString(input.Password)
	if err != nil {
		return
	}

	newUser := enUser.User{
		Username: input.Username,
		Password: hashPassword,
	}

	txContext := uc.uow.Begin(ctx)

	if err = uc.userRepository.CreateUser(&txContext, &newUser); err != nil {
		uc.uow.Rollback(txContext)
		return
	}

	tokenDetail, err := uc.jwt.GenerateToken(&enAuth.UserJwtClaims{
		UserID:   newUser.ID,
		Username: newUser.Username,
	})
	if err != nil {
		return
	}

	newUserToken := enUser.UserToken{
		UserID: newUser.ID,
		Token: enUser.UserTokenDetail{
			ID:        tokenDetail.ID,
			Token:     tokenDetail.Token,
			ExpiresAt: tokenDetail.ExpiresAt,
		},
	}

	if err = uc.userTokenRepository.StoreToken(&txContext, &newUserToken); err != nil {
		uc.uow.Rollback(txContext)
		return
	}

	uc.uow.Commit(txContext)

	output.ID = newUser.ID
	output.Username = newUser.Username
	output.Token = tokenDetail.Token

	return
}
