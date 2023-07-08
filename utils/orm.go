package utils

import (
	"database/sql"
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormOrm(
	dialect string,
	conn *sql.DB,
) (*gorm.DB, error) {

	var gormDialect gorm.Dialector
	if dialect == "mysql" {
		gormDialect = mysql.New(mysql.Config{Conn: conn})
	} else {
		return nil, errors.New("sql dialect is not defined")
	}

	db, err := gorm.Open(gormDialect)

	return db, err

}
