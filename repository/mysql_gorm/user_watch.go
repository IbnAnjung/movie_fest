package mysql_gorm

import (
	"context"
	"errors"

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
