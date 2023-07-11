package mysql_gorm

import (
	"context"
	"errors"
	"fmt"
	"time"

	enUserVote "github.com/IbnAnjung/movie_fest/entity/user_vote"
	"github.com/IbnAnjung/movie_fest/repository/mysql_gorm/models"
	"github.com/IbnAnjung/movie_fest/utils"

	"gorm.io/gorm"
)

type userVoteRepository struct {
	db *gorm.DB
}

func NewUserVoteRepository(db *gorm.DB) enUserVote.UserVoteRepository {
	return &userVoteRepository{
		db: db,
	}
}

func (r userVoteRepository) CreateVote(ctx context.Context, userVote *enUserVote.UserVote) error {
	m := models.UserVote{}
	m.FillFromEntity(*userVote)

	db := getTxSessionDB(ctx, r.db)
	if err := db.WithContext(ctx).Create(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			e := utils.DuplicatedDataError
			err = e
		}

		return err
	}
	m.ToEntity(userVote)
	return nil

}

func (r userVoteRepository) RemoveVote(ctx context.Context, userVote *enUserVote.UserVote) error {
	db := getTxSessionDB(ctx, r.db)
	if err := db.WithContext(ctx).Where("user_id = ?", userVote.UserID).
		Where("movie_id = ?", userVote.MovieID).Delete(models.UserVote{}).Error; err != nil {
		return err
	}

	return nil
}

func (r userVoteRepository) GetVotedMoviesByUserID(ctx context.Context, userID int64) (movies []enUserVote.VotedMovieList, err error) {
	type model struct {
		models.UserVote
		Title    string `gorm:"column:title"`
		Duration int64  `gorm:"column:duration"`
		Artist   string `gorm:"column:artists"`
	}

	m := []model{}
	db := getTxSessionDB(ctx, r.db)

	tnUserVote := models.UserVote{}.TableName()
	tnMovie := models.Movie{}.TableName()

	err = db.Model(&models.UserVote{}).
		Select(fmt.Sprintf("user_id, movie_id, %s.created_at, title, duration, artists", tnUserVote)).
		Joins(fmt.Sprintf("JOIN %s ON %s.movie_id = %s.id", tnMovie, tnUserVote, tnMovie)).
		Find(&m).
		Error
	if err != nil {
		return
	}

	for _, v := range m {
		movies = append(movies, enUserVote.VotedMovieList{
			ID:       v.MovieID,
			Title:    v.Title,
			Duration: time.Duration(v.Duration * int64(time.Second)),
			Artist:   v.Artist,
		})
	}

	return
}
