package models

import (
	"time"

	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
)

type Movie struct {
	ID           int64  `gorm:"column:id"`
	Filename     string `gorm:"column:filename"`
	Title        string `gorm:"column:title"`
	Duration     int64  `gorm:"column:duration"`
	Artists      string `gorm:"column:artists"`
	Description  string `gorm:"column:description"`
	ViewsCounter int64  `gorm:"column:views_counter"`
	VotesCounter int64  `gorm:"column:votes_counter"`
	UploadedAt   string `gorm:"column:uploaded_at;<-:false"`
}

func (Movie) TableName() string {
	return "movies"
}

func (m Movie) ToEntity(en *enMovie.Movie) error {
	uploadTime, err := time.Parse("2006-01-02 15:04:05", m.UploadedAt)
	if err != nil {
		return err
	}

	en.ID = m.ID
	en.Filename = m.Filename
	en.Title = m.Title
	en.Duration = m.Duration
	en.Artists = m.Artists
	en.Description = m.Description
	en.ViewsCounter = m.ViewsCounter
	en.VotesCounter = m.VotesCounter
	en.UploadTime = uploadTime

	return nil
}

func (m *Movie) FillFromEntity(en enMovie.Movie) {
	m.ID = en.ID
	m.Filename = en.Filename
	m.Title = en.Title
	m.Duration = en.Duration
	m.Artists = en.Artists
	m.Description = en.Description
	m.ViewsCounter = en.ViewsCounter
	m.VotesCounter = en.VotesCounter
	m.UploadedAt = en.UploadTime.Format("2006-01-02 15:04:05")
}
