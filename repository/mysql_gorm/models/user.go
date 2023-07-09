package models

import (
	enUser "github.com/IbnAnjung/movie_fest/entity/users"
)

type User struct {
	ID       int64  `gorm:"column:id;<-:false"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

func (User) TableName() string {
	return "users"
}

func (m User) ToEntity(en *enUser.User) error {
	en.ID = m.ID
	en.Username = m.Username
	en.Password = m.Password
	return nil
}

func (m *User) FillFromEntity(en enUser.User) {
	m.ID = en.ID
	m.Username = en.Username
	m.Password = en.Password
}
