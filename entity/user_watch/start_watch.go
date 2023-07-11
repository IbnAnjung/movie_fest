package user_watch

import "time"

type StartPlayInput struct {
	UserID    int64
	MovieID   int64
	StartTime time.Duration
}
