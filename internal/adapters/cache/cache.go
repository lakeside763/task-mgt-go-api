package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(addr string) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisCache{client: client}
}

func (c *RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(context.TODO(), key, value, expiration).Err()
}

func (c *RedisCache) Get(key string) (string, error) {
	return c.client.Get(context.TODO(), key).Result()
}