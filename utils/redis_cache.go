package utils

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	enUtil "github.com/IbnAnjung/movie_fest/entity/utils"
	"github.com/redis/go-redis/v9"
)

type RedisCaching struct {
	conn *redis.Client
}

var mutex = sync.Mutex{}

func NewRedisCaching(
	conn *redis.Client,
) enUtil.Caching {
	return RedisCaching{
		conn: conn,
	}
}

func (r RedisCaching) Set(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	return r.conn.Set(ctx, key, value, exp).Err()
}

func (r RedisCaching) Get(ctx context.Context, key string) (string, error) {
	val, err := r.conn.Get(ctx, key).Result()

	if err != redis.Nil {
		log.Printf("GET ERROR %s", err.Error())
		return "", errors.New("fail get redis data")
	}

	return val, nil
}
func (r RedisCaching) Del(ctx context.Context, key string) error {
	if err := r.conn.Del(ctx, key).Err(); err != nil {
		return errors.New("fail delete redis data")
	}

	return nil
}
