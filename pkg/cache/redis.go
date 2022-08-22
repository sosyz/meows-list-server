package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisStore struct {
	C *redis.Client
}

// NewRedisStore 创建新的redis存储
func NewRedisStore(address, password string, database int) *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       database,
	})
	return &RedisStore{C: rdb}
}

// Set 存储值
func (r *RedisStore) Set(key string, value string, ttl int) error {
	ctx := context.Background()
	if ttl > 0 {
		if err := r.C.Set(ctx, key, value, time.Duration(ttl)*time.Second).Err(); err != nil {
			return err
		}
	} else {
		if err := r.C.Set(ctx, key, value, 0).Err(); err != nil {
			return err
		}
	}
	return nil
}

// Get 取值
func (r *RedisStore) Get(key string) (string, error) {
	ctx := context.Background()
	v, err := r.C.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return v, nil
}
