package movie

import (
	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
	enUtil "github.com/IbnAnjung/movie_fest/entity/utils"
)

type MovieUC struct {
	unitOfWork      enUtil.UnitOfWork
	storageService  enUtil.Storage
	stringGenerator enUtil.StringGenerator
	movieRepository enMovie.MovieRepository
}

func NewMovieUC(

	unitOfwork enUtil.UnitOfWork,
	storageService enUtil.Storage,
	stringGenerator enUtil.StringGenerator,
	movieRepository enMovie.MovieRepository,
) enMovie.MovieUseCase {
	return MovieUC{
		unitOfWork:      unitOfwork,
		storageService:  storageService,
		stringGenerator: stringGenerator,
		movieRepository: movieRepository,
	}
}
