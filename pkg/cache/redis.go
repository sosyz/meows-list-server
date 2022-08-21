package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

// NewRedisStore 创建新的redis存储
func NewRedisStore(address, password string, database int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       database,
	})
	return rdb
}

// Set 存储值
func Set(rdb *redis.Client, key string, value any, ttl int) error {
	ctx := context.Background()
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	if err := rdb.Set(ctx, key, v, time.Duration(ttl)).Err(); err != nil {
		return err
	}
	return nil
}

// Get 取值
func Get(rdb *redis.Client, key string, value any) error {
	ctx := context.Background()
	v, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(v), value)
	if err != nil {
		return err
	}
	return nil
}
