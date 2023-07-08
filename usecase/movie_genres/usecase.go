package movie_genres

import (
	enMovieGenres "github.com/IbnAnjung/movie_fest/entity/movie_genres"
)

type MovieGenresUC struct {
	movieGenresRepository enMovieGenres.MovieGenresRepository
}

func NewMovieUC(
	movieGenresRepository enMovieGenres.MovieGenresRepository,
) enMovieGenres.MovieGenresUseCase {
	return MovieGenresUC{
		movieGenresRepository: movieGenresRepository,
	}
}
