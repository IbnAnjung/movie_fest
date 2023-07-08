package movie

import (
	"context"

	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
	"github.com/IbnAnjung/movie_fest/utils"
)

type input struct {
	Title       string  `validate:"required"`
	Description string  `validate:"required,min=5"`
	Duration    int64   `validate:"required,min=5"`
	Artists     string  `validate:"required,max=255"`
	Genres      []int32 `validate:"required,dive"`
	MovieID     int64   `validate:"required,min=1"`
}

func (i *input) set(src enMovie.UpdateMetaDataInput) {
	i.Title = src.Title
	i.Description = src.Description
	i.Duration = src.Duration
	i.Artists = src.Artists
	i.Genres = src.Genres
	i.MovieID = src.MovieID
}

func (uc MovieUC) UpdateMetaData(ctx context.Context, enInput enMovie.UpdateMetaDataInput) (err error) {
	iv := input{}
	iv.set(enInput)

	if err = uc.validator.ValidateStruct(&iv); err != nil {
		return
	}

	movie, err := uc.movieRepository.GetMovieByID(&ctx, enInput.MovieID)
	if err != nil {
		return err
	}

	movieGenres, err := uc.movieGenresRepository.GetGenresByIDs(&ctx, enInput.Genres)
	if err != nil {
		return err
	}

	if len(movieGenres) != len(enInput.Genres) {
		e := utils.UnprocessableEntityError
		e.Message = "Some Genres not exists"
		return e
	}

	movie.Title = enInput.Title
	movie.Duration = enInput.Duration
	movie.Artists = enInput.Artists
	movie.Description = enInput.Description
	movie.Genres = enInput.Genres

	txContext := uc.unitOfWork.Begin(ctx)
	if err := uc.movieRepository.UpdateMovieMetaData(&txContext, &movie); err != nil {
		uc.unitOfWork.Rollback(txContext)
		return err
	}

	mg := []enMovie.MovieHasGenres{}
	for _, movieGenderID := range enInput.Genres {
		mg = append(mg, enMovie.MovieHasGenres{
			MovieID:      movie.ID,
			MovieGenreID: movieGenderID,
		})
	}

	if err := uc.movieHasGenresRepository.Add(&txContext, mg); err != nil {
		uc.unitOfWork.Rollback(txContext)
		return err
	}

	uc.unitOfWork.Commit(txContext)

	return nil
}
