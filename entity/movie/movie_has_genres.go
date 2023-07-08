package movie

import "context"

type MovieHasGenres struct {
	MovieID      int64
	MovieGenreID int32
	GenreName    string
}

type MovieHasGenresRepository interface {
	Add(ctx *context.Context, genres []MovieHasGenres) error
	GetMovieGenres(ctx *context.Context, movieID int64) (res []MovieHasGenres, err error)
}
