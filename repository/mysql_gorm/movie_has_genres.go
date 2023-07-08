package mysql_gorm

import (
	"context"
	"fmt"

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

func (r *movieHasGenresRepository) GetMovieGenres(ctx *context.Context, movieID int64) (res []enMovie.MovieHasGenres, err error) {
	db := getTxSessionDB(*ctx, r.db)

	type model struct {
		models.MovieHasGenres
		GenreName string `gorm:"column:genre_name"`
	}

	m := []model{}
	mhg := models.MovieHasGenres{}.TableName()
	mg := models.MovieGenres{}.TableName()
	if err = db.Model(models.MovieHasGenres{}).
		Select("movie_id, movie_genre_id, name genre_name").
		Joins(fmt.Sprintf("JOIN %s ON %s.movie_genre_id = %s.id", mg, mhg, mg)).
		Where(fmt.Sprintf("%s.movie_id = ?", mhg), movieID).
		Find(&m).Error; err != nil {
		return
	}

	for _, v := range m {
		en := enMovie.MovieHasGenres{}
		v.ToEntity(&en)
		en.GenreName = v.GenreName
		res = append(res, en)
	}

	return
}
