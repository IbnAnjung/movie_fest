package movie_genres

import "context"

type MovieGenres struct {
	ID           int32
	Name         string
	ViewsCounter int64
	VotesCounter int64
}

type MovieGenresRepository interface {
	GetGenresByIDs(ctx *context.Context, genresIDs []int32) (movieGenres []MovieGenres, err error)
}
