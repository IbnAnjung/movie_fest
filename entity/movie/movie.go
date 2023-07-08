package movie

import (
	"context"
	"time"
)

type Movie struct {
	ID           int64
	Filename     string
	Title        string
	Duration     int64
	Artists      string
	Genres       []int32
	Description  string
	ViewsCounter int64
	VotesCounter int64
	UploadTime   time.Time
}

type MovieRepository interface {
	AddMovie(ctx *context.Context, movie *Movie) error
	GetMovieByID(ctx *context.Context, movieID int64) (movie Movie, err error)
	UpdateMovieMetaData(ctx *context.Context, newMovie *Movie) (err error)
}

type MovieUseCase interface {
	UploadMovie(ctx context.Context, input UploadMovieInput) (newMoview UploadMovieOutput, err error)
	UpdateMetaData(ctx context.Context, input UpdateMetaDataInput) (err error)
}
