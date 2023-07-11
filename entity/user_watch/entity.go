package user_watch

import (
	"context"
	"time"
)

const PlaybackUniqueIDCacheKey = "user:playback_unique_id"
const PlaybackUniqueIDCacheLockKey = "user:playback_unique_id:lock"

type UserWatch struct {
	ID         int64
	UserID     int64
	MovieID    int64
	StartTime  time.Duration
	EndTime    time.Duration
	Duration   time.Duration
	ExpireTime time.Time
}
type UserWatchUseCase interface {
	StartPlay(ctx context.Context, input StartPlayInput) (userWatch UserWatch, playbackID string, err error)
	Playback(ctx context.Context, input PlaybackInput) (err error)
	History(ctx context.Context, input HistoryInput) (output []UserWathHistory, err error)
}

type UserWatchRepository interface {
	CreateUserWatch(ctx context.Context, userWath *UserWatch) error
	UpdateUserWatch(ctx context.Context, userWath *UserWatch) error
	GetByID(ctx context.Context, ID int64) (userWath *UserWatch, err error)
	GetUserWathHistory(ctx context.Context, userID int64) (userWathhistories []UserWathHistory, err error)
}
