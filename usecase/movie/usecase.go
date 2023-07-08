package movie

import (
	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
	enMovieGenres "github.com/IbnAnjung/movie_fest/entity/movie_genres"
	enUtil "github.com/IbnAnjung/movie_fest/entity/utils"
)

type MovieUC struct {
	unitOfWork               enUtil.UnitOfWork
	storageService           enUtil.Storage
	stringGenerator          enUtil.StringGenerator
	validator                enUtil.Validator
	movieRepository          enMovie.MovieRepository
	movieGenresRepository    enMovieGenres.MovieGenresRepository
	movieHasGenresRepository enMovie.MovieHasGenresRepository
}

func NewMovieUC(

	unitOfwork enUtil.UnitOfWork,
	storageService enUtil.Storage,
	validator enUtil.Validator,
	stringGenerator enUtil.StringGenerator,
	movieRepository enMovie.MovieRepository,
	movieGenresRepository enMovieGenres.MovieGenresRepository,
	movieHasGenresRepository enMovie.MovieHasGenresRepository,
) enMovie.MovieUseCase {
	return MovieUC{
		unitOfWork:               unitOfwork,
		storageService:           storageService,
		validator:                validator,
		stringGenerator:          stringGenerator,
		movieRepository:          movieRepository,
		movieGenresRepository:    movieGenresRepository,
		movieHasGenresRepository: movieHasGenresRepository,
	}
}
