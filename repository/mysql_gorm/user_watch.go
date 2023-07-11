package mysql_gorm

import (
	"context"

	enUserWatch "github.com/IbnAnjung/movie_fest/entity/user_watch"
	"github.com/IbnAnjung/movie_fest/repository/mysql_gorm/models"

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
