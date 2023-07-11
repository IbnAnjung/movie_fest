package movie_genres

import (
	"context"

	enMovieGenres "github.com/IbnAnjung/movie_fest/entity/movie_genres"
)

func (uc MovieGenresUC) GetMostVote(ctx context.Context) (mov enMovieGenres.MovieGenres, err error) {
	return uc.movieGenresRepository.GetMostVote(&ctx)
}
