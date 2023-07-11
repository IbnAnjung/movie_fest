package user_watch

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	enUserWatch "github.com/IbnAnjung/movie_fest/entity/user_watch"
	"github.com/IbnAnjung/movie_fest/utils"
)

type startPlayInputVlidationObject struct {
	UserID    int64         `validate:"required"`
	MovieID   int64         `validate:"required"`
	StartTime time.Duration `validate:"required"`
}

func (i *startPlayInputVlidationObject) set(src enUserWatch.StartPlayInput) {
	i.UserID = src.UserID
	i.MovieID = src.MovieID
	i.StartTime = src.StartTime
}

func (uc userWatchUC) StartPlay(ctx context.Context, input enUserWatch.StartPlayInput) (output enUserWatch.UserWatch, playbackId string, err error) {

	iv := startPlayInputVlidationObject{}
	iv.set(input)

	if err = uc.validator.ValidateStruct(&iv); err != nil {
		return
	}

	_, err = uc.movieRepository.GetMovieByID(&ctx, input.MovieID)
	if err != nil {
		if errors.Is(err, utils.DataNotFoundError) {
			e := utils.DataNotFoundError
			e.Message = "movie not found"
			err = e
		}

		return
	}
	log.Println("strt => ", input.StartTime)

	txContext := uc.uow.Begin(ctx)

	userWatch := enUserWatch.UserWatch{
		UserID:     input.UserID,
		MovieID:    input.MovieID,
		StartTime:  input.StartTime,
		EndTime:    input.StartTime,
		ExpireTime: time.Now().Add(6 * time.Hour),
	}

	if err = uc.userWatchRepository.CreateUserWatch(txContext, &userWatch); err != nil {
		uc.uow.Rollback(txContext)
		return
	}

	playbackId = fmt.Sprintf("%d-%s", userWatch.ID, uc.stringGenerator.UUID())
	unqueKey := fmt.Sprintf("%s:%s", enUserWatch.PlaybackUniqueIDCacheKey, playbackId)
	if err = uc.cache.Set(ctx, unqueKey, userWatch.ID, 3*time.Second); err != nil {
		uc.uow.Rollback(txContext)
		return
	}

	uc.uow.Commit(txContext)

	output = userWatch

	return
}
