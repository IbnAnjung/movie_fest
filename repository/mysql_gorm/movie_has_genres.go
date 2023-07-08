package mysql_gorm

import (
	"context"

	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
	"github.com/IbnAnjung/movie_fest/repository/mysql_gorm/models"
	"gorm.io/gorm"
)

type movieHasGenresRepository struct {
	db *gorm.DB
}

func NewMovieHasGenresRepository(db *gorm.DB) enMovie.MovieHasGenresRepository {
	return &movieHasGenresRepository{
		db: db,
	}
}

func (r *movieHasGenresRepository) Add(ctx *context.Context, genres []enMovie.MovieHasGenres) error {
	db := getTxSessionDB(*ctx, r.db)

	ms := []models.MovieHasGenres{}
	for _, v := range genres {
		m := models.MovieHasGenres{}
		m.FillFromEntity(v)

		ms = append(ms, m)
	}

	return db.Create(ms).Error

}
