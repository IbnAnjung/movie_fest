package mysql_gorm

import (
	"context"
	"errors"

	enMovie "github.com/IbnAnjung/movie_fest/entity/movie"
	"github.com/IbnAnjung/movie_fest/repository/mysql_gorm/models"
	"github.com/IbnAnjung/movie_fest/utils"
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

func (r *movieRepository) GetMovieByID(ctx *context.Context, movieID int64) (movie enMovie.Movie, err error) {
	db := getTxSessionDB(*ctx, r.db)

	m := models.Movie{}
	if err = db.Where("id = ?", movieID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e := utils.DataNotFoundError
			e.Message = "Movie Not Found"
			err = e
		}

		return
	}

	m.ToEntity(&movie)
	return
}

func (r *movieRepository) UpdateMovieMetaData(ctx *context.Context, newMovie *enMovie.Movie) (err error) {
	db := getTxSessionDB(*ctx, r.db)
	m := models.Movie{}
	m.FillFromEntity(*newMovie)

	if err := db.Save(&m).Error; err != nil {
		return err
	}

	m.ToEntity(newMovie)
	return nil
}
