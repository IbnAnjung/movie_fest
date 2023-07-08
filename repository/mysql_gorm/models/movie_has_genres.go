package models

import (
	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
)

type MovieHasGenres struct {
	MovieID      int64 `gorm:"column:movie_id"`
	MovieGenreID int32 `gorm:"column:movie_genre_id"`
}

func (MovieHasGenres) TableName() string {
	return "movie_has_genres"
}

func (m MovieHasGenres) ToEntity(en *enMovie.MovieHasGenres) error {
	en.MovieID = m.MovieID
	en.MovieGenreID = m.MovieGenreID

	return nil
}

func (m *MovieHasGenres) FillFromEntity(en enMovie.MovieHasGenres) {
	m.MovieID = en.MovieID
	m.MovieGenreID = en.MovieGenreID
}
