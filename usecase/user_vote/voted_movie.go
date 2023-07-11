package user_vote

import (
	"context"

	enUserVote "github.com/IbnAnjung/movie_fest/entity/user_vote"
)

func (uc userVoteUC) GetVotedMovies(ctx context.Context, userID int64) (movies []enUserVote.VotedMovieList, err error) {
	return uc.userVoteRepository.GetVotedMoviesByUserID(ctx, userID)
}
