package mysql_gorm

import (
	"context"

	enUser "github.com/IbnAnjung/movie_fest/entity/users"
	"github.com/IbnAnjung/movie_fest/repository/mysql_gorm/models"
	"gorm.io/gorm"
)

type userTokenRepository struct {
	db *gorm.DB
}

func NewUserTokenRepository(db *gorm.DB) enUser.UserTokenRepository {
	return &userTokenRepository{
		db: db,
	}
}

func (r *userTokenRepository) StoreToken(ctx *context.Context, userToken *enUser.UserToken) (err error) {
	db := getTxSessionDB(*ctx, r.db)

	m := models.UserToken{}
	m.FillFromEntity(*userToken)

	err = db.Create(&m).Error

	m.ToEntity(userToken)
	return
}
