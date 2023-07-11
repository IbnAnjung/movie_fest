package user_vote

import (
	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
	enMovieGenres "github.com/IbnAnjung/movie_fest/entity/movie_genres"
	enUserVote "github.com/IbnAnjung/movie_fest/entity/user_vote"
	enUtil "github.com/IbnAnjung/movie_fest/entity/utils"
)

type userVoteUC struct {
	uow                      enUtil.UnitOfWork
	movieRepository          enMovie.MovieRepository
	movieGenresRepository    enMovieGenres.MovieGenresRepository
	movieHasGenresRepository enMovie.MovieHasGenresRepository
	userVoteRepository       enUserVote.UserVoteRepository
}

func NewUserVoteUseCase(
	uow enUtil.UnitOfWork,
	movieRepository enMovie.MovieRepository,
	movieGenresRepository enMovieGenres.MovieGenresRepository,
	movieHasGenresRepository enMovie.MovieHasGenresRepository,
	userVoteRepository enUserVote.UserVoteRepository,
) enUserVote.UserVoteUseCase {
	return userVoteUC{
		uow:                      uow,
		movieRepository:          movieRepository,
		movieGenresRepository:    movieGenresRepository,
		movieHasGenresRepository: movieHasGenresRepository,
		userVoteRepository:       userVoteRepository,
	}
}
