package models

import (
	"time"

	enUserWatch "github.com/IbnAnjung/movie_fest/entity/user_watch"
)

type UserWatch struct {
	ID        int64  `gorm:"column:id"`
	UserID    int64  `gorm:"column:user_id"`
	MovieID   int64  `gorm:"column:movie_id"`
	StartTime int64  `gorm:"column:start_time"`
	EndTime   int64  `gorm:"column:end_time"`
	Duration  int64  `gorm:"column:duration"`
	ExpiretAt string `gorm:"column:expired_at"`
}

func (UserWatch) TableName() string {
	return "user_watches"
}

func (m UserWatch) ToEntity(en *enUserWatch.UserWatch) error {
	expireTime, _ := time.Parse("2006-01-02 15:04:05", m.ExpiretAt)
	en.ID = m.ID
	en.UserID = m.UserID
	en.MovieID = m.MovieID
	en.Duration = time.Duration(m.Duration) * time.Second
	en.StartTime = time.Duration(m.StartTime) * time.Second
	en.EndTime = time.Duration(m.EndTime) * time.Second
	en.ExpireTime = expireTime

	return nil
}

func (m *UserWatch) FillFromEntity(en enUserWatch.UserWatch) {
	m.ID = en.ID
	m.UserID = en.UserID
	m.MovieID = en.MovieID
	m.StartTime = int64(en.StartTime / time.Second)
	m.EndTime = int64(en.EndTime / time.Second)
	m.Duration = int64(en.Duration / time.Second)
	m.ExpiretAt = en.ExpireTime.Format("2006-01-02 15:04:05")
}
