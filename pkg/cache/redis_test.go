package cache

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRedisStore(t *testing.T) {
	r := NewRedisStore("127.0.0.1:6379", "", 0)
	if r == nil {
		println("NewRedisStore error")
	}
}

func TestRedisStore_Get(t *testing.T) {
	assertions := assert.New(t)
	r := NewRedisStore("127.0.0.1:6379", "", 0)
	if r == nil {
		assertions.Error(errors.New("NewRedisStore error"))
	}
	value, err := r.Get("test")
	if err != nil {
		assertions.Error(err)
	}
	assertions.Equal("123", value)
}

func TestRedisStore_Set(t *testing.T) {
	assertions := assert.New(t)
	r := NewRedisStore("127.0.0.1:6379", "", 0)
	if r == nil {
		assertions.Error(errors.New("NewRedisStore error"))
	}

	err := r.Set("test", "123", 0)
	if err != nil {
		assertions.Error(err)
	}
}
