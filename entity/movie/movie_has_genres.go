package movie

import "context"

type MovieHasGenres struct {
	MovieID      int64
	MovieGenreID int32
}

type MovieHasGenresRepository interface {
	Add(ctx *context.Context, genres []MovieHasGenres) error
}
