package mysql_gorm

import (
	"context"
	"errors"

	enUser "github.com/IbnAnjung/movie_fest/entity/users"
	"github.com/IbnAnjung/movie_fest/repository/mysql_gorm/models"
	"github.com/IbnAnjung/movie_fest/utils"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) enUser.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindUserByUsername(ctx *context.Context, username string) (user enUser.User, err error) {
	db := getTxSessionDB(*ctx, r.db)
	m := models.User{}

	err = db.Where("username = ?", username).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e := utils.DataNotFoundError
			e.Message = "user not found"
			err = e
		}

		return
	}

	m.ToEntity(&user)
	return
}

func (r *userRepository) CreateUser(ctx *context.Context, user *enUser.User) (err error) {
	db := getTxSessionDB(*ctx, r.db)

	m := models.User{}
	m.FillFromEntity(*user)

	err = db.Create(&m).Error

	m.ToEntity(user)
	return
}
