package movie

import (
	"context"

	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
)

func (uc MovieUC) ViewMovie(ctx context.Context, mov enMovie.MovieDetail) (err error) {

	txContext := uc.unitOfWork.Begin(ctx)

	if err = uc.movieRepository.IncreaseViews(&txContext, mov.ID); err != nil {
		uc.unitOfWork.Rollback(txContext)
		return err
	}

	genresIDs := []int32{}
	for _, genre := range mov.Genres {
		genresIDs = append(genresIDs, genre.ID)
	}

	if err = uc.movieGenresRepository.IncreaseViews(&txContext, genresIDs); err != nil {
		uc.unitOfWork.Rollback(txContext)
		return err
	}

	uc.unitOfWork.Commit(txContext)

	return
}
