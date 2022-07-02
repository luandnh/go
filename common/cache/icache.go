package cache

import (
	"time"
)

type IMemCache interface {
	Set(key string, value interface{}) error
	SetTTL(key string, value interface{}, t time.Duration) error
	Get(key string) (interface{}, error)
	Del(key string) error
	Close()
}

type IRedisCache interface {
	Set(key string, value interface{}) error
	SetTTL(key string, value interface{}, t time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
	Close()
}

var RCache IRedisCache
var MCache IMemCache
