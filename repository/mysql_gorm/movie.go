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

func (r *movieRepository) GetMostView(ctx *context.Context) (movie enMovie.Movie, err error) {
	db := getTxSessionDB(*ctx, r.db)

	m := models.Movie{}
	if err = db.Order("views_counter DESC").
		Find(&m).Error; err != nil {
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

func (r *movieRepository) IncreaseViews(ctx *context.Context, movieID int64) error {
	db := getTxSessionDB(*ctx, r.db)

	return db.Model(&models.Movie{}).
		Where("id = ?", movieID).
		UpdateColumn("views_counter", gorm.Expr("views_counter + ? ", 1)).Error
}

func (r *movieRepository) GetListPagination(ctx *context.Context, offset, limit int) (movies []enMovie.Movie, totalRaw int64, err error) {
	db := getTxSessionDB(*ctx, r.db)

	query := db.Model(&models.Movie{}).
		Order("id").
		Session(&gorm.Session{})

	if err = query.Count(&totalRaw).Error; err != nil {
		return
	}

	mMovies := []models.Movie{}
	if err = query.Offset(offset).Limit(limit).Find(&mMovies).Error; err != nil {
		return
	}

	for _, m := range mMovies {
		movie := enMovie.Movie{}
		m.ToEntity(&movie)

		movies = append(movies, movie)
	}

	return
}
