package movie

import (
	"context"

	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
)

func (uc MovieUC) GetMostVotes(ctx context.Context) (mov enMovie.Movie, err error) {
	return uc.movieRepository.GetMostVotes(&ctx)
}

type GetVoteInputValidationObject struct {
	Sort string `validate:"omitempty,oneof=ASC asc DESC desc"`
}

func (i *GetVoteInputValidationObject) set(src enMovie.GetVotesInput) {
	i.Sort = src.Sort
}

func (uc MovieUC) GetVotes(ctx context.Context, input enMovie.GetVotesInput) (movies []enMovie.Movie, err error) {
	iv := GetVoteInputValidationObject{}
	iv.set(input)
	if err = uc.validator.ValidateStruct(&iv); err != nil {
		return
	}

	if input.Sort == "" {
		input.Sort = "ASC"
	}

	movies, err = uc.movieRepository.GetAllMovieVotes(&ctx, input)
	if err != nil {
		return
	}
	return
}
