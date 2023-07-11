package user_watch

import (
	"context"
	"fmt"
	"strconv"
	"time"

	enUserWatch "github.com/IbnAnjung/movie_fest/entity/user_watch"
	"github.com/IbnAnjung/movie_fest/utils"
)

type startPlayBackVlidationObject struct {
	PlaybackID string        `validate:"required"`
	EndTime    time.Duration `validate:"required"`
}

func (i *startPlayBackVlidationObject) set(src enUserWatch.PlaybackInput) {
	i.PlaybackID = src.PlaybackID
	i.EndTime = src.EndTime
}

func (uc userWatchUC) Playback(ctx context.Context, input enUserWatch.PlaybackInput) (err error) {

	iv := startPlayBackVlidationObject{}
	iv.set(input)

	if err = uc.validator.ValidateStruct(&iv); err != nil {
		return
	}

	uniqueKey := fmt.Sprintf("%s:%s", enUserWatch.PlaybackUniqueIDCacheKey, input.PlaybackID)
	lockKey := fmt.Sprintf("%s:%s", enUserWatch.PlaybackUniqueIDCacheLockKey, input.PlaybackID)

	lock, err := uc.cache.Get(ctx, lockKey)
	if err != nil {
		return err
	}

	if lock != "" {
		e := utils.UnprocessableEntityError
		e.Message = "playback locked"
		err = e

		return err
	}

	val, err := uc.cache.Get(ctx, uniqueKey)
	if err != nil {
		return err
	}

	if val == "" {
		e := utils.UnprocessableEntityError
		e.Message = "expire playback"
		err = e
		return err
	}

	id, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return
	}

	userWatch, err := uc.userWatchRepository.GetByID(ctx, id)
	if err != nil {
		return
	}

	movie, err := uc.movieRepository.GetMovieByID(&ctx, userWatch.MovieID)
	if err != nil {
		return err
	}

	userWatch.EndTime = input.EndTime
	userWatch.Duration = userWatch.Duration + (time.Second * 3)
	movie.WatchDurationCounter = movie.WatchDurationCounter + 3

	txContext := uc.uow.Begin(ctx)

	if err = uc.userWatchRepository.UpdateUserWatch(txContext, userWatch); err != nil {
		uc.uow.Rollback(txContext)
		return
	}

	if err = uc.movieRepository.UpdateMovieMetaData(&txContext, &movie); err != nil {
		uc.uow.Rollback(txContext)
		return
	}

	if err = uc.cache.Set(ctx, uniqueKey, userWatch.ID, 10*time.Minute); err != nil {
		uc.uow.Rollback(txContext)
		return
	}

	if err = uc.cache.Set(ctx, lockKey, true, time.Duration(3)*time.Second); err != nil {
		uc.uow.Rollback(txContext)
		return
	}

	uc.uow.Commit(txContext)

	return
}
