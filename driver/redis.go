package driver

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type redisConfig struct {
	host     string
	port     int
	user     string
	password string
	db       int
}

func LoadRedisConfig(
	host string,
	port int,
	user string,
	password string,
	db int,
) redisConfig {
	return redisConfig{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		db:       db,
	}
}

func NewRedisConnection(ctx context.Context, config redisConfig) (*redis.Client, func(), error) {
	redisOpt := redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.host, config.port),
		Username: config.user,
		Password: config.password,
		DB:       config.db,
	}

	rdsClient := redis.NewClient(&redisOpt)

	if err := rdsClient.Ping(ctx).Err(); err != nil {
		return nil, nil, fmt.Errorf("redis ping failed: %w", err)
	}

	cleanup := func() {
		if err := rdsClient.Close(); err != nil {
			log.Printf("redis close connection failed %e", err)
			return
		}
		fmt.Println("redis close connection success")
	}

	return rdsClient, cleanup, nil

}
