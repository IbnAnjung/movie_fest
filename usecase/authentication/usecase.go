package authentication

import (
	enAuth "github.com/IbnAnjung/movie_fest/entity/authentication"
	enUser "github.com/IbnAnjung/movie_fest/entity/users"
	enUtils "github.com/IbnAnjung/movie_fest/entity/utils"
)

type authenticationUC struct {
	jwt                 enUtils.Jwt
	crypt               enUtils.Crypt
	uow                 enUtils.UnitOfWork
	validator           enUtils.Validator
	userRepository      enUser.UserRepository
	userTokenRepository enUser.UserTokenRepository
}

func NewAuthenticationUC(
	jwt enUtils.Jwt,
	crypt enUtils.Crypt,
	uow enUtils.UnitOfWork,
	validator enUtils.Validator,
	userRepository enUser.UserRepository,
	userTokenRepository enUser.UserTokenRepository,
) enAuth.AuthenticationUsecase {
	return &authenticationUC{
		jwt:                 jwt,
		uow:                 uow,
		crypt:               crypt,
		validator:           validator,
		userRepository:      userRepository,
		userTokenRepository: userTokenRepository,
	}
}
