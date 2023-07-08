package movie

import (
	"context"

	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
)

func (uc MovieUC) GetMostView(ctx context.Context) (mov enMovie.Movie, err error) {
	return uc.movieRepository.GetMostView(&ctx)
}
