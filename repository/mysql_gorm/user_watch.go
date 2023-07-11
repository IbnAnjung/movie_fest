package mysql_gorm

import (
	"context"
	"errors"
	"fmt"
	"time"

	enUserWatch "github.com/IbnAnjung/movie_fest/entity/user_watch"
	"github.com/IbnAnjung/movie_fest/repository/mysql_gorm/models"
	"github.com/IbnAnjung/movie_fest/utils"

	"gorm.io/gorm"
)

type userWatchRepository struct {
	db *gorm.DB
}

func NewUserWatchRepository(db *gorm.DB) enUserWatch.UserWatchRepository {
	return &userWatchRepository{
		db: db,
	}
}

func (r userWatchRepository) CreateUserWatch(ctx context.Context, userWath *enUserWatch.UserWatch) error {
	m := models.UserWatch{}
	m.FillFromEntity(*userWath)

	db := getTxSessionDB(ctx, r.db)
	if err := db.WithContext(ctx).Create(&m).Error; err != nil {
		return err
	}

	m.ToEntity(userWath)
	return nil

}

func (r userWatchRepository) UpdateUserWatch(ctx context.Context, userWath *enUserWatch.UserWatch) error {
	m := models.UserWatch{}
	m.FillFromEntity(*userWath)

	db := getTxSessionDB(ctx, r.db)
	if err := db.WithContext(ctx).Save(&m).Error; err != nil {
		return err
	}

	m.ToEntity(userWath)
	return nil

}

func (r userWatchRepository) GetByID(ctx context.Context, ID int64) (userWath *enUserWatch.UserWatch, err error) {
	m := models.UserWatch{}
	userWath = &enUserWatch.UserWatch{}

	db := getTxSessionDB(ctx, r.db)
	if err = db.WithContext(ctx).Where("id", ID).
		First(&m).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e := utils.DataNotFoundError
			e.Message = "session not found"
			err = e
		}

		return
	}

	m.ToEntity(userWath)

	return

}

func (r userWatchRepository) GetUserWathHistory(ctx context.Context, userID int64) (userWath []enUserWatch.UserWathHistory, err error) {

	type jmodel struct {
		models.UserWatch
		Title         string `gorm:"column:title"`
		Description   string `gorm:"column:description"`
		Artists       string `gorm:"column:artists"`
		TotalDuration int64  `gorm:"column:total_duration"`
	}

	m := models.Movie{}.TableName()
	uw := models.UserWatch{}.TableName()

	db := getTxSessionDB(ctx, r.db)

	model := []jmodel{}
	if err = db.WithContext(ctx).Where("user_id", userID).
		Select(fmt.Sprintf("%s.movie_id, %s.title, %s.description, %s.artists, sum(%s.duration) total_duration", uw, m, m, m, uw)).
		Joins(fmt.Sprintf("JOIN %s ON %s.movie_id = %s.id", m, uw, m)).
		Group("movie_id, title, description, artists").
		Order("total_duration desc").
		Find(&model).
		Error; err != nil {
		return
	}

	for _, v := range model {
		userWath = append(userWath, enUserWatch.UserWathHistory{
			MovieID:     v.MovieID,
			Title:       v.Title,
			Description: v.Description,
			Artists:     v.Artists,
			Duration:    time.Duration(v.TotalDuration) * time.Second,
		})
	}

	return

}
