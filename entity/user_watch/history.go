package user_watch

import "time"

type HistoryInput struct {
	UserID int64
}

type UserWathHistory struct {
	MovieID     int64
	Title       string
	Description string
	Artists     string
	Duration    time.Duration
}
