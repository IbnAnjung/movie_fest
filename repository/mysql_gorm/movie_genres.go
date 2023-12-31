package mysql_gorm

import (
	"context"

	enMovieGenres "github.com/IbnAnjung/movie_fest/entity/movie_genres"
	"github.com/IbnAnjung/movie_fest/repository/mysql_gorm/models"
	"gorm.io/gorm"
)

type movieGenresRepository struct {
	db *gorm.DB
}

func NewMovieGenresRepository(db *gorm.DB) enMovieGenres.MovieGenresRepository {
	return &movieGenresRepository{
		db: db,
	}
}

func (r *movieGenresRepository) GetGenresByIDs(ctx *context.Context, genresIDs []int32) (movieGenres []enMovieGenres.MovieGenres, err error) {
	db := getTxSessionDB(*ctx, r.db)

	m := []models.MovieGenres{}
	movieGenres = []enMovieGenres.MovieGenres{}

	err = db.WithContext(*ctx).
		Where("id IN (?)", genresIDs).
		Find(&m).Error
	if err != nil {
		return movieGenres, err
	}

	for _, v := range m {
		g := enMovieGenres.MovieGenres{}
		v.ToEntity(&g)
		movieGenres = append(movieGenres, g)
	}

	return
}

func (r *movieGenresRepository) IncreaseViews(ctx *context.Context, genresIDS []int32) error {
	db := getTxSessionDB(*ctx, r.db)

	return db.Model(&models.MovieGenres{}).
		Where("id IN (?)", genresIDS).
		UpdateColumn("views_counter", gorm.Expr("views_counter + ? ", 1)).Error
}

func (r *movieGenresRepository) IncreaseVotes(ctx *context.Context, genresIDS []int32) error {
	db := getTxSessionDB(*ctx, r.db)

	return db.Model(&models.MovieGenres{}).
		Where("id IN (?)", genresIDS).
		UpdateColumn("votes_counter", gorm.Expr("votes_counter + ? ", 1)).Error
}

func (r *movieGenresRepository) DecreaseVotes(ctx *context.Context, genresIDS []int32) error {
	db := getTxSessionDB(*ctx, r.db)

	return db.Model(&models.MovieGenres{}).
		Where("id IN (?)", genresIDS).
		UpdateColumn("votes_counter", gorm.Expr("votes_counter - ? ", 1)).Error
}

func (r *movieGenresRepository) GetMostView(ctx *context.Context) (mg enMovieGenres.MovieGenres, err error) {
	db := getTxSessionDB(*ctx, r.db)

	m := models.MovieGenres{}
	if err = db.Order("views_counter DESC").
		Find(&m).Error; err != nil {
		return
	}

	m.ToEntity(&mg)
	return
}

func (r *movieGenresRepository) GetMostVote(ctx *context.Context) (mg enMovieGenres.MovieGenres, err error) {
	db := getTxSessionDB(*ctx, r.db)

	m := models.MovieGenres{}
	if err = db.Order("votes_counter DESC").
		Find(&m).Error; err != nil {
		return
	}

	m.ToEntity(&mg)
	return
}
