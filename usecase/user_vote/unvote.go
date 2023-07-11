package user_vote

import (
	"context"
	"errors"

	enUserVote "github.com/IbnAnjung/movie_fest/entity/user_vote"
	"github.com/IbnAnjung/movie_fest/utils"
)

func (uc userVoteUC) Unvote(ctx context.Context, userVote *enUserVote.UserVote) error {

	if _, err := uc.movieRepository.GetMovieByID(&ctx, userVote.MovieID); err != nil {
		if errors.Is(err, utils.DataNotFoundError) {
			e := utils.DataNotFoundError
			e.Message = "movie not found"
			err = e
		}

		return err
	}

	if err := uc.userVoteRepository.RemoveVote(ctx, userVote); err != nil {
		return err
	}

	return nil
}
