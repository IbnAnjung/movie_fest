package utils

import (
	"context"
	"time"
)

type Caching interface {
	Set(key string, value interface{}) Caching
	PushList(key string, value interface{}) Caching
	Expire(duration time.Duration) Caching
	ExpireAt(t time.Time) Caching
	Do(ctx context.Context) error
	Get(ctx context.Context, key string) (string, error)
	GetList(ctx context.Context, key string, from, to int64) ([]string, error)
	Del(ctx context.Context, key string) error
}
