package authentication

import (
	"context"
	"fmt"

	enAuth "github.com/IbnAnjung/movie_fest/entity/authentication"
	enUser "github.com/IbnAnjung/movie_fest/entity/users"
	"github.com/IbnAnjung/movie_fest/utils"
)

type loginValidationObject struct {
	Username string `validate:"required,min=5,max=50"`
	Password string `validate:"required,min=5,max=255"`
}

func (i *loginValidationObject) set(src enAuth.Login) {
	i.Username = src.Username
	i.Password = src.Password
}

func (uc *authenticationUC) Login(ctx context.Context, input enAuth.Login) (output enAuth.LogedinUser, err error) {

	iv := loginValidationObject{}
	iv.set(input)

	if err = uc.validator.ValidateStruct(&iv); err != nil {
		return
	}

	user, err := uc.userRepository.FindUserByUsername(&ctx, input.Username)
	if err != nil {
		return
	}

	if ok := uc.crypt.VerifyHash(user.Password, input.Password); !ok {
		e := utils.UnprocessableEntityError
		e.Message = "username and password not match"
		err = e
		return
	}

	tokenDetail, err := uc.jwt.GenerateToken(&enAuth.UserJwtClaims{
		UserID:   user.ID,
		Username: user.Username,
	})
	if err != nil {
		return
	}

	newUserToken := enUser.UserToken{
		UserID: user.ID,
		Token: enUser.UserTokenDetail{
			ID:        tokenDetail.ID,
			Token:     tokenDetail.Token,
			ExpiresAt: tokenDetail.ExpiresAt,
		},
	}

	fmt.Println("expire => ", tokenDetail.ExpiresAt.Format("2006-01-02 15:04:05"))

	if err = uc.userTokenRepository.StoreToken(&ctx, &newUserToken); err != nil {
		return
	}

	output.ID = user.ID
	output.Username = user.Username
	output.Token = tokenDetail.Token

	return
}
