package user_vote

import "time"

type VotedMovieList struct {
	ID       int64
	Title    string
	Duration time.Duration
	Artist   string
}
