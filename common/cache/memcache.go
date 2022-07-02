package cache

import (
	"context"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/eko/gocache/v3/cache"
	"github.com/eko/gocache/v3/store"
)

type MemCacheManager struct {
	cache *cache.Cache[interface{}]
}

func NewMemCacheManager() IMemCache {
	bigcacheClient, _ := bigcache.NewBigCache(bigcache.Config{})
	bigcacheStore := store.NewBigcache(bigcacheClient)
	cacheManager := cache.New[interface{}](bigcacheStore)
	return &MemCacheManager{
		cache: cacheManager,
	}
}

func (m *MemCacheManager) Close() {
	m.Close()
}

func (m *MemCacheManager) Set(key string, value interface{}) error {
	return m.cache.Set(context.Background(), key, value)

}

func (m *MemCacheManager) Get(key string) (interface{}, error) {
	return m.cache.Get(context.Background(), key)
}

func (m *MemCacheManager) Del(key string) error {
	return m.cache.Delete(context.Background(), key)
}

func (m *MemCacheManager) SetTTL(key string, value interface{}, ttl time.Duration) error {
	return m.cache.Set(context.Background(), key, value, store.WithExpiration(ttl))
}
