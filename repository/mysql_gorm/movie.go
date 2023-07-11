package mysql_gorm

import (
	"context"
	"errors"
	"fmt"

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

func (r *movieRepository) GetMovieByIDs(ctx *context.Context, movieIDs []int64) (movies []enMovie.Movie, err error) {
	db := getTxSessionDB(*ctx, r.db)

	m := []models.Movie{}
	if err = db.Where("id in (?)", movieIDs).
		Find(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e := utils.DataNotFoundError
			e.Message = "Movie Not Found"
			err = e
		}

		return
	}
	for _, v := range m {
		movie := enMovie.Movie{}
		v.ToEntity(&movie)

		movies = append(movies, movie)
	}
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

func (r *movieRepository) IncreaseVotes(ctx *context.Context, movieID int64) error {
	db := getTxSessionDB(*ctx, r.db)

	return db.Model(&models.Movie{}).
		Where("id = ?", movieID).
		UpdateColumn("votes_counter", gorm.Expr("votes_counter + ? ", 1)).Error
}

func (r *movieRepository) DecreaseVotes(ctx *context.Context, movieID int64) error {
	db := getTxSessionDB(*ctx, r.db)

	return db.Model(&models.Movie{}).
		Where("id = ?", movieID).
		UpdateColumn("votes_counter", gorm.Expr("votes_counter - ? ", 1)).Error
}

func (r *movieRepository) GetListPagination(ctx *context.Context, input enMovie.ListMovieWithPaginationInput) (movies []enMovie.Movie, totalRaw int64, err error) {
	db := getTxSessionDB(*ctx, r.db)

	m := models.Movie{}.TableName()
	mg := models.MovieGenres{}.TableName()
	mhg := models.MovieHasGenres{}.TableName()

	query := db.Model(&models.Movie{}).Select(fmt.Sprintf("DISTINCT %s.*", m)).
		Order("id")

	if input.Search != "" {
		query.Joins(fmt.Sprintf("JOIN %s ON %s.id = %s.movie_id", mhg, m, mhg)).
			Joins(fmt.Sprintf("JOIN %s ON %s.movie_genre_id = %s.id", mg, mhg, mg)).
			Where(fmt.Sprintf("%s.title LIKE ?", m), "%"+input.Search+"%").
			Or(fmt.Sprintf("%s.description LIKE ?", m), "%"+input.Search+"%").
			Or(fmt.Sprintf("%s.artists LIKE ?", m), "%"+input.Search+"%").
			Or(fmt.Sprintf("%s.name LIKE ?", mg), "%"+input.Search+"%")
	}

	query.Session(&gorm.Session{})

	if err = query.Count(&totalRaw).Error; err != nil {
		return
	}

	mMovies := []models.Movie{}
	if err = query.Offset(input.Offset).Limit(input.Limit).Find(&mMovies).Error; err != nil {
		return
	}

	for _, m := range mMovies {
		movie := enMovie.Movie{}
		m.ToEntity(&movie)

		movies = append(movies, movie)
	}

	return
}

func (r *movieRepository) GetAllMovieViews(ctx *context.Context, input enMovie.GetViewsInput) (movies []enMovie.Movie, err error) {
	db := getTxSessionDB(*ctx, r.db)

	m := []models.Movie{}

	query := db.Model(&models.Movie{}).Select("*").
		Order(fmt.Sprintf("views_counter %s", input.Sort))

	if input.MinViews != 0 {
		query.Where("views_counter >= ?", input.MinViews)
	}

	if input.MaxViews != 0 {
		query.Where("views_counter <= ?", input.MaxViews)
	}

	err = query.Find(&m).Error
	if err != nil {
		return
	}

	for _, v := range m {
		movie := enMovie.Movie{}
		v.ToEntity(&movie)

		movies = append(movies, movie)
	}

	return
}

func (r *movieRepository) GetMostVotes(ctx *context.Context) (movie enMovie.Movie, err error) {
	db := getTxSessionDB(*ctx, r.db)

	m := models.Movie{}
	if err = db.Order("votes_counter DESC").
		Find(&m).Error; err != nil {
		return
	}

	m.ToEntity(&movie)
	return
}

func (r *movieRepository) GetAllMovieVotes(ctx *context.Context, input enMovie.GetVotesInput) (movies []enMovie.Movie, err error) {
	db := getTxSessionDB(*ctx, r.db)

	m := []models.Movie{}

	query := db.Model(&models.Movie{}).Select("*").
		Order(fmt.Sprintf("votes_counter %s", input.Sort))

	if input.MinVotes != 0 {
		query.Where("votes_counter >= ?", input.MinVotes)
	}

	if input.MaxVotes != 0 {
		query.Where("votes_counter <= ?", input.MaxVotes)
	}

	err = query.Find(&m).Error
	if err != nil {
		return
	}

	for _, v := range m {
		movie := enMovie.Movie{}
		v.ToEntity(&movie)

		movies = append(movies, movie)
	}

	return
}
