package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	cache *redis.Client
}

func NewRedisCache(client *redis.Client) IRedisCache {
	return &RedisCache{
		cache: client,
	}
}

var ctx = context.Background()

func (c *RedisCache) Set(key string, value interface{}) error {
	_, err := c.cache.Set(ctx, key, value, redis.KeepTTL).Result()
	return err
}

func (c *RedisCache) SetTTL(key string, value interface{}, ttl time.Duration) error {
	_, err := c.cache.Set(ctx, key, value, ttl).Result()
	return err
}

func (c *RedisCache) Get(key string) (string, error) {
	value, err := c.cache.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return value, err
}

func (c *RedisCache) Close() {
	c.Close()
}

func (c *RedisCache) Del(key string) error {
	_, err := c.cache.Del(ctx, key).Result()
	return err
}
