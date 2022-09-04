package redisclient

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
)

const (
	defaultTimeout time.Duration = 2 * time.Second
)

type RedisClient struct {
	redis   *redis.Client
	timeout time.Duration
}

func NewRedisClient() *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &RedisClient{
		redis:   rdb,
		timeout: defaultTimeout,
	}
}

func (r RedisClient) Insert(key string, value interface{}) error {
	if r.redis == nil {
		return ErrClientNotConnect
	}

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	return r.redis.Set(ctx, key, value, r.timeout).Err()
}
