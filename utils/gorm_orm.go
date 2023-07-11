package utils

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormOrm(
	appMode string,
	dialect string,
	conn *sql.DB,
) (*gorm.DB, error) {

	var gormDialect gorm.Dialector
	if dialect == "mysql" {
		gormDialect = mysql.New(mysql.Config{Conn: conn})
	} else {
		return nil, errors.New("sql dialect is not defined")
	}

	logLevel := logger.Info
	if appMode == "production" {
		logLevel = logger.Warn
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,    // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(gormDialect, &gorm.Config{
		Logger:         newLogger,
		TranslateError: true,
	})

	return db, err

}
