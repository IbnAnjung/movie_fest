package user_watch

import (
	"context"
	"time"
)

const PlaybackUniqueIDCacheKey = "user:playback_unque_id"

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
}

type UserWatchRepository interface {
	CreateUserWatch(ctx context.Context, userWath *UserWatch) error
}
