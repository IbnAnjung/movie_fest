package movie

import "context"

type MovieUseCase interface {
	UploadMovie(ctx context.Context, input UploadMovieInput) (newMoview UploadMovieOutput, err error)
}
