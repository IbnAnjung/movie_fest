package user_vote

import (
	"context"
	"errors"
	"fmt"

	enUserVote "github.com/IbnAnjung/movie_fest/entity/user_vote"
	"github.com/IbnAnjung/movie_fest/utils"
)

func (uc userVoteUC) Vote(ctx context.Context, userVote *enUserVote.UserVote) error {

	if _, err := uc.movieRepository.GetMovieByID(&ctx, userVote.MovieID); err != nil {
		if errors.Is(err, utils.DataNotFoundError) {
			e := utils.DataNotFoundError
			e.Message = "movie not found"
			err = e
		}

		return err
	}

	if err := uc.userVoteRepository.CreateVote(ctx, userVote); err != nil {
		if errors.Is(err, utils.DuplicatedDataError) {
			e := utils.DuplicatedDataError
			e.Message = "you already vote this movie"
			err = e
		}
		fmt.Println("error => ", err)

		return err
	}

	return nil
}
