package user_vote

import (
	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
	enUserVote "github.com/IbnAnjung/movie_fest/entity/user_vote"
)

type userVoteUC struct {
	movieRepository    enMovie.MovieRepository
	userVoteRepository enUserVote.UserVoteRepository
}

func NewUserVoteUseCase(movieRepository enMovie.MovieRepository, userVoteRepository enUserVote.UserVoteRepository) enUserVote.UserVoteUseCase {
	return userVoteUC{
		movieRepository:    movieRepository,
		userVoteRepository: userVoteRepository,
	}
}
