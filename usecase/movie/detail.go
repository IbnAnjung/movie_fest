package movie

import (
	"context"

	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
)

func (uc MovieUC) GetMovieDetail(ctx context.Context, movieID int64) (mov enMovie.MovieDetail, err error) {
	enMov, err := uc.movieRepository.GetMovieByID(&ctx, movieID)
	if err != nil {
		return
	}

	genres, err := uc.movieHasGenresRepository.GetMovieGenres(&ctx, enMov.ID)
	if err != nil {
		return
	}

	publicUrl := uc.storageService.GetPublicFullPath(enMov.Filename)

	mov.ID = enMov.ID
	mov.Filename = enMov.Filename
	mov.Title = enMov.Title
	mov.Duration = enMov.Duration
	mov.Artists = enMov.Artists
	mov.Description = enMov.Description
	mov.PublicUrl = publicUrl
	mov.ViewsCounter = enMov.ViewsCounter
	mov.VotesCounter = enMov.VotesCounter
	mov.Genres = []enMovie.MovieDetailGenres{}

	for _, genre := range genres {
		mov.Genres = append(mov.Genres, enMovie.MovieDetailGenres{
			ID:   genre.MovieGenreID,
			Name: genre.GenreName,
		})
	}

	return
}
