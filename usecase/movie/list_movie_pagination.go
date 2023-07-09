package movie

import (
	"context"

	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
)

func (uc MovieUC) GetListMovieWithPagination(ctx context.Context, input enMovie.ListMovieWithPaginationInput) (output enMovie.ListMovieWithPaginationOutput, err error) {

	uc.pagination.Init(&input.MetaPagination)

	movies, totalRaw, err := uc.movieRepository.GetListPagination(&ctx, input)
	if err != nil {
		return output, err
	}

	uc.pagination.SetTotalRaw(totalRaw)

	output.Meta = uc.pagination.GetMeta()
	output.Movies = movies

	return
}
