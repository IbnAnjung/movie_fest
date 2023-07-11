package models

import (
	"time"

	enUserVote "github.com/IbnAnjung/movie_fest/entity/user_vote"
)

type UserVote struct {
	UserID    int64  `gorm:"column:user_id"`
	MovieID   int64  `gorm:"column:movie_id"`
	CreatedAt string `gorm:"column:created_at;<-:false"`
}

func (UserVote) TableName() string {
	return "user_votes"
}

func (m UserVote) ToEntity(en *enUserVote.UserVote) error {
	createdAt, _ := time.Parse("2006-01-02 15:04:05", m.CreatedAt)

	en.UserID = m.UserID
	en.MovieID = m.MovieID
	en.VoteTime = createdAt

	return nil
}

func (m *UserVote) FillFromEntity(en enUserVote.UserVote) {
	m.UserID = en.UserID
	m.MovieID = en.MovieID
	m.CreatedAt = en.VoteTime.Format("2006-01-02 15:04:05")
}
