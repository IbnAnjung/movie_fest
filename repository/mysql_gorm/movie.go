package mysql_gorm

import (
	"context"

	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
	"github.com/IbnAnjung/movie_fest/repository/mysql_gorm/models"
	"gorm.io/gorm"
)

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) enMovie.MovieRepository {
	return &movieRepository{
		db: db,
	}
}

func (r *movieRepository) AddMovie(ctx *context.Context, movie *enMovie.Movie) error {
	db := getTxSessionDB(*ctx, r.db)

	m := models.Movie{}
	m.FillFromEntity(*movie)

	err := db.WithContext(*ctx).Create(&m).Error
	if err != nil {
		return err
	}

	m.ToEntity(movie)
	return nil
}
