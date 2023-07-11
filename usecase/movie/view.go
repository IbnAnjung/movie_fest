package movie

import (
	"context"
	"log"

	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
)

type GetViewInputValidationObject struct {
	Sort string `validate:"omitempty,oneof=ASC asc DESC desc"`
}

func (i *GetViewInputValidationObject) set(src enMovie.GetViewsInput) {
	i.Sort = src.Sort
}

func (uc MovieUC) ViewMovie(ctx context.Context, mov enMovie.MovieDetail) (err error) {
	txContext := uc.unitOfWork.Begin(ctx)

	if err = uc.movieRepository.IncreaseViews(&txContext, mov.ID); err != nil {
		uc.unitOfWork.Rollback(txContext)
		return err
	}

	genresIDs := []int32{}
	for _, genre := range mov.Genres {
		genresIDs = append(genresIDs, genre.ID)
	}

	if err = uc.movieGenresRepository.IncreaseViews(&txContext, genresIDs); err != nil {
		uc.unitOfWork.Rollback(txContext)
		return err
	}

	uc.unitOfWork.Commit(txContext)

	return
}

func (uc MovieUC) GetViews(ctx context.Context, input enMovie.GetViewsInput) (movies []enMovie.Movie, err error) {
	iv := GetViewInputValidationObject{}
	iv.set(input)
	log.Print(input.MinViews)
	log.Print(input.MaxViews)
	if err = uc.validator.ValidateStruct(&iv); err != nil {
		return
	}
	log.Print(iv.Sort)

	if input.Sort == "" {
		input.Sort = "ASC"
	}

	movies, err = uc.movieRepository.GetAllMovieViews(&ctx, input)
	if err != nil {
		return
	}
	return
}
