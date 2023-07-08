package movie_genres

import "context"

type MovieGenres struct {
	ID           int32
	Name         string
	ViewsCounter int64
	VotesCounter int64
}

type MovieGenresUseCase interface {
	GetMostView(ctx context.Context) (mg MovieGenres, err error)
}

type MovieGenresRepository interface {
	GetMostView(ctx *context.Context) (mg MovieGenres, err error)
	IncreaseViews(ctx *context.Context, genresIDS []int32) error
	GetGenresByIDs(ctx *context.Context, genresIDs []int32) (movieGenres []MovieGenres, err error)
}
