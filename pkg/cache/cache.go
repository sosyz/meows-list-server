package cache

import (
	"encoding/json"
)

var cache *RedisStore

func Init(host, port, password string, db int) {
	cache = NewRedisStore(
		host+":"+port,
		password,
		db,
	)
}

func Get[T any](key string) (*T, error) {
	v, err := cache.Get(key)
	if err != nil {
		return nil, err
	}
	var ret T
	// 反序列化
	if err := json.Unmarshal([]byte(v), &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func Put(key string, value any, ttl int) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	if err := cache.Set(key, string(v), ttl); err != nil {
		return err
	}
	return nil
}
