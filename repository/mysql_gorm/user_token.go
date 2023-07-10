package mysql_gorm

import (
	"context"
	"errors"
	"strconv"

	enUser "github.com/IbnAnjung/movie_fest/entity/users"
	"github.com/IbnAnjung/movie_fest/repository/mysql_gorm/models"
	"github.com/IbnAnjung/movie_fest/utils"
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

func (r *userTokenRepository) GetToken(ctx *context.Context, userToken *enUser.UserToken) (err error) {
	db := getTxSessionDB(*ctx, r.db)

	m := models.UserToken{}
	m.FillFromEntity(*userToken)

	err = db.Where("id = ?", m.ID).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e := utils.DataNotFoundError
			e.Message = "Token Not Found"
			err = e
		}

		return err
	}

	m.ToEntity(userToken)
	return
}

func (r *userTokenRepository) DeleteToken(ctx *context.Context, id string) (err error) {
	db := getTxSessionDB(*ctx, r.db)

	inID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	return db.Where("id = ?", inID).Delete(&models.UserToken{}).Error
}
