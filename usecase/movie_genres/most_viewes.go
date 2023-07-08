package movie_genres

import (
	"context"

	enMovieGenres "github.com/IbnAnjung/movie_fest/entity/movie_genres"
)

func (uc MovieGenresUC) GetMostView(ctx context.Context) (mov enMovieGenres.MovieGenres, err error) {
	return uc.movieGenresRepository.GetMostView(&ctx)
}
