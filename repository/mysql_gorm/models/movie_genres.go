package models

import (
	enMovieGenres "github.com/IbnAnjung/movie_fest/entity/movie_genres"
)

type MovieGenres struct {
	ID           int32  `gorm:"column:id;<-:false"`
	Name         string `gorm:"column:name"`
	ViewsCounter int64  `gorm:"views_counter"`
	VotesCounter int64  `gorm:"votes_counter"`
}

func (MovieGenres) TableName() string {
	return "movie_genres"
}

func (m MovieGenres) ToEntity(en *enMovieGenres.MovieGenres) error {
	en.ID = m.ID
	en.Name = m.Name
	en.ViewsCounter = m.ViewsCounter
	en.VotesCounter = m.VotesCounter

	return nil
}

func (m *MovieGenres) FillFromEntity(en enMovieGenres.MovieGenres) {
	m.ID = en.ID
	m.Name = en.Name
	m.ViewsCounter = en.ViewsCounter
	m.VotesCounter = en.VotesCounter
}
