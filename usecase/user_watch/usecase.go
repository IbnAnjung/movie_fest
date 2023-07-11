package user_watch

import (
	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
	enUserWath "github.com/IbnAnjung/movie_fest/entity/user_watch"
	enUtil "github.com/IbnAnjung/movie_fest/entity/utils"
)

type userWatchUC struct {
	cache               enUtil.Caching
	uow                 enUtil.UnitOfWork
	validator           enUtil.Validator
	stringGenerator     enUtil.StringGenerator
	movieRepository     enMovie.MovieRepository
	userWatchRepository enUserWath.UserWatchRepository
}

func NewUserVoteUseCase(
	uow enUtil.UnitOfWork,
	cache enUtil.Caching,
	validator enUtil.Validator,
	stringGenerator enUtil.StringGenerator,
	movieRepository enMovie.MovieRepository,
	userWatchRepository enUserWath.UserWatchRepository,
) enUserWath.UserWatchUseCase {
	return userWatchUC{
		uow:                 uow,
		cache:               cache,
		validator:           validator,
		stringGenerator:     stringGenerator,
		movieRepository:     movieRepository,
		userWatchRepository: userWatchRepository,
	}
}
