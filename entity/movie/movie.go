package movie

import (
	"time"
)

type Movie struct {
	ID           int64
	Filename     string
	ViewsCounter int64
	VotesCounter int64
	UploadTime   time.Time
	UploadUserID string
}
