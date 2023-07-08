package movie

import "context"

type MovieRepository interface {
	AddMovie(ctx *context.Context, movie *Movie) error
}
