package user_vote

import (
	"context"
	"time"
)

type UserVote struct {
	UserID   int64
	MovieID  int64
	VoteTime time.Time
}

type UserVoteRepository interface {
	CreateVote(ctx context.Context, userVote *UserVote) error
	RemoveVote(ctx context.Context, userVote *UserVote) error
	GetVotedMoviesByUserID(ctx context.Context, userID int64) (movies []VotedMovieList, err error)
}

type UserVoteUseCase interface {
	Vote(ctx context.Context, userVote *UserVote) error
	Unvote(ctx context.Context, userVote *UserVote) error
	GetVotedMovies(ctx context.Context, userID int64) (movies []VotedMovieList, err error)
}
