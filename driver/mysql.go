package driver

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type sqlConfig struct {
	host               string
	port               int
	user               string
	password           string
	schema             string
	maxLifeConnection  time.Duration
	maxConnection      int
	maxIddleConnection int
}

func LoadMySqlConfig(
	host string,
	port int,
	user string,
	password string,
	schema string,
	maxLifeConnection int64,
	maxConnection int,
	maxIddleConnection int,
) sqlConfig {

	maxLfConn := time.Second * time.Duration(maxLifeConnection)

	return sqlConfig{
		host:               host,
		port:               port,
		user:               user,
		password:           password,
		schema:             schema,
		maxLifeConnection:  maxLfConn,
		maxConnection:      maxConnection,
		maxIddleConnection: maxIddleConnection,
	}
}

func NewMysqlConnection(ctx context.Context, config sqlConfig) (*sql.DB, func(), error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.user,
		config.password,
		config.host,
		config.port,
		config.schema,
	)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, nil, err
	}

	if err = conn.Ping(); err != nil {
		fmt.Printf("dsn: %s", dsn)
		return nil, nil, err
	}

	// See "Important settings" section.
	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)

	cleanup := func() {
		if err := conn.Close(); err != nil {
			log.Printf("mysql close connection fail, error: %s", err.Error())
			return
		}
		fmt.Println("mysql close connection success")
	}
	return conn, cleanup, nil
}
