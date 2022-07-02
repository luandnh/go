package cache

import (
	"time"

	"github.com/jellydator/ttlcache/v2"
)

type MemCache struct {
	ttlCache ttlcache.SimpleCache
}

func NewMemCache() IMemCache {
	cache := ttlcache.NewCache()
	cache.SkipTTLExtensionOnHit(true)
	return &MemCache{
		ttlCache: cache,
	}
}

func (c *MemCache) Set(key string, value interface{}) error {
	err := c.ttlCache.Set(key, value)
	return err
}

func (c *MemCache) Del(key string) error {
	err := c.ttlCache.Remove(key)
	return err
}

func (c *MemCache) SetTTL(key string, value interface{}, ttl time.Duration) error {
	err := c.ttlCache.SetWithTTL(key, value, ttl)
	return err
}

func (c *MemCache) Get(key string) (interface{}, error) {
	value, err := c.ttlCache.Get(key)
	if err == ttlcache.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func (c *MemCache) Close() {
	c.ttlCache.Close()
}
