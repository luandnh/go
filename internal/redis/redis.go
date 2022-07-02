package redis

import (
	"context"
	"errors"
	"go-project/common/log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type IRedis interface {
	GetClient() *redis.Client
	Connect() error
	Ping() error
	Set(key string, value interface{}) (string, error)
	SetTTL(key string, value interface{}, t time.Duration) (string, error)
	Get(key string) (string, error)
	IsExisted(key string) (bool, error)
	IsHExisted(list, key string) (bool, error)
	HGet(list, key string) (string, error)
	HGetAll(list string) (map[string]string, error)
	HSet(key string, values []interface{}) (int64, error)
	HMGet(key string, fields ...string) ([]interface{}, error)
	HMSet(key string, values ...interface{}) error
	HMDel(key string, fields ...string) error
	FLUSHALL() interface{}
	Del(key []string) error
	HDel(key string, fields ...string) error
	GetKeysPattern(pattern string) ([]string, error)
}

var Redis IRedis

type RedisClient struct {
	Client *redis.Client
	config Config
}

type Config struct {
	Addr         string
	Password     string
	DB           int
	PoolSize     int
	PoolTimeout  int
	IdleTimeout  int
	ReadTimeout  int
	WriteTimeout int
}

func NewRedis(config Config) (IRedis, error) {
	r := &RedisClient{
		config: config,
	}
	if err := r.Connect(); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *RedisClient) GetClient() *redis.Client {
	return r.Client
}

func (r *RedisClient) Connect() error {
	Client := redis.NewClient(&redis.Options{
		Addr:         r.config.Addr,
		Password:     r.config.Password,
		DB:           r.config.DB,
		PoolSize:     r.config.PoolSize,
		PoolTimeout:  time.Duration(r.config.PoolTimeout) * time.Second,
		IdleTimeout:  time.Duration(r.config.IdleTimeout) * time.Second,
		ReadTimeout:  time.Duration(r.config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(r.config.WriteTimeout) * time.Second,
	})
	str, err := Client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Info(str)
	r.Client = Client
	return nil
}

func (r *RedisClient) Ping() error {
	_, err := r.Client.Ping(ctx).Result()
	return err
}

func (r *RedisClient) Set(key string, value interface{}) (string, error) {
	ret, err := r.Client.Set(ctx, key, value, 0).Result()
	return ret, err
}

//Set - Set a value with key to Redis DB
func (r *RedisClient) SetTTL(key string, value interface{}, t time.Duration) (string, error) {
	ret, err := r.Client.Set(ctx, key, value, t).Result()
	return ret, err
}

func (r *RedisClient) Get(key string) (string, error) {
	ret, err := r.Client.Get(ctx, key).Result()
	return ret, err
}

func (r *RedisClient) IsExisted(key string) (bool, error) {
	res, err := r.Client.Exists(ctx, key).Result()
	if res == 0 || err != nil {
		return false, err
	}
	return true, nil
}

func (r *RedisClient) IsHExisted(list, key string) (bool, error) {
	res, err := r.Client.HExists(ctx, list, key).Result()
	if res == false || err != nil {
		return false, err
	}
	return true, nil
}

func (r *RedisClient) HGet(list, key string) (string, error) {
	ret, err := r.Client.HGet(ctx, list, key).Result()
	return ret, err
}

func (r *RedisClient) HGetAll(list string) (map[string]string, error) {
	ret, err := r.Client.HGetAll(ctx, list).Result()
	return ret, err
}

func (r *RedisClient) HSet(key string, values []interface{}) (int64, error) {
	ret, err := r.Client.HSet(ctx, key, values...).Result()
	return ret, err
}

func (r *RedisClient) Del(key []string) error {
	err := r.Client.Del(ctx, key...).Err()
	return err
}

func (r *RedisClient) HMSet(key string, values ...interface{}) error {
	ret, err := r.Client.HMSet(ctx, key, values...).Result()
	if err != nil {
		return err
	}
	if !ret {
		err = errors.New("HashMap Set failed")
	}
	return err
}

func (r *RedisClient) HMDel(key string, fields ...string) error {
	err := r.Client.HDel(ctx, key, fields...).Err()
	return err
}

func (r *RedisClient) FLUSHALL() interface{} {
	ret := r.Client.FlushAll(ctx)
	return ret
}

func (r *RedisClient) HMGet(key string, fields ...string) ([]interface{}, error) {
	ret, err := r.Client.HMGet(ctx, key, fields...).Result()
	return ret, err
}

func (r *RedisClient) HDel(key string, fields ...string) error {
	err := r.Client.HDel(ctx, key, fields...).Err()
	return err
}

func (r *RedisClient) GetKeysPattern(pattern string) ([]string, error) {
	ret, err := r.Client.Keys(ctx, pattern).Result()
	return ret, err
}
