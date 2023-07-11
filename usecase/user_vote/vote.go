package user_vote

import (
	"context"
	"errors"

	enUserVote "github.com/IbnAnjung/movie_fest/entity/user_vote"
	"github.com/IbnAnjung/movie_fest/utils"
)

func (uc userVoteUC) Vote(ctx context.Context, userVote *enUserVote.UserVote) error {

	mov, err := uc.movieRepository.GetMovieByID(&ctx, userVote.MovieID)
	if err != nil {
		if errors.Is(err, utils.DataNotFoundError) {
			e := utils.DataNotFoundError
			e.Message = "movie not found"
			err = e
		}

		return err
	}

	txContext := uc.uow.Begin(ctx)

	if err = uc.userVoteRepository.CreateVote(txContext, userVote); err != nil {
		if errors.Is(err, utils.DuplicatedDataError) {
			e := utils.DuplicatedDataError
			e.Message = "you already vote this movie"
			err = e
		}

		uc.uow.Rollback(txContext)
		return err
	}

	if err = uc.movieRepository.IncreaseVotes(&txContext, userVote.MovieID); err != nil {
		uc.uow.Rollback(txContext)
		return err
	}

	genres, err := uc.movieHasGenresRepository.GetMovieGenres(&ctx, mov.ID)
	if err != nil {
		uc.uow.Rollback(txContext)
		return err
	}

	genresIDs := []int32{}
	for _, genre := range genres {
		genresIDs = append(genresIDs, genre.MovieGenreID)
	}

	if err = uc.movieGenresRepository.IncreaseVotes(&txContext, genresIDs); err != nil {
		uc.uow.Rollback(txContext)
		return err
	}

	uc.uow.Commit(txContext)

	return nil
}
