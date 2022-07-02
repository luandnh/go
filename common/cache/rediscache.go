package cache

import (
	"context"
	"time"

	"github.com/eko/gocache/v3/cache"
	"github.com/eko/gocache/v3/store"
	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	cache *cache.Cache[interface{}]
}

func NewRedisCache(client *redis.Client) IRedisCache {
	redisStore := store.NewRedis(client)
	cacheManager := cache.New[interface{}](redisStore)
	return &RedisCache{
		cache: cacheManager,
	}
}

var ctx = context.Background()

func (c *RedisCache) Set(key string, value interface{}) error {
	return c.cache.Set(ctx, key, value)
}

func (c *RedisCache) SetTTL(key string, value interface{}, ttl time.Duration) error {
	return c.cache.Set(ctx, key, value, store.WithExpiration(ttl))
}

func (c *RedisCache) Get(key string) (interface{}, error) {
	value, err := c.cache.Get(ctx, key)
	return value, err
}

func (c *RedisCache) Close() {
	c.Close()
}

func (c *RedisCache) Del(key string) error {
	return c.cache.Delete(ctx, key)
}
