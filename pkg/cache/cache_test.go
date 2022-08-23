package cache

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type test struct {
	Name string
	Age  int
}

func TestPut(t *testing.T) {
	Init("127.0.0.1", "6379", "", 0)

	assertions := assert.New(t)
	if cache.C == nil {
		assertions.Error(errors.New("cache is nil"))
	}

	td := test{
		Name: "test",
		Age:  18,
	}

	err := Put("test", td, 0)
	if err != nil {
		return
	}
}

func TestGet(t *testing.T) {
	Init("127.0.0.1", "6379", "", 0)

	assertions := assert.New(t)
	if cache.C == nil {
		assertions.Error(errors.New("cache is nil"))
	}
	td := test{
		Name: "test",
		Age:  18,
	}
	value, err := Get[test]("test")
	if err != nil {
		assertions.Error(err)
	}
	assertions.Equal(td, *value)
}

func TestDel(t *testing.T) {
	Init("127.0.0.1", "6379", "", 0)

	assertions := assert.New(t)
	if cache.C == nil {
		assertions.Error(errors.New("cache is nil"))
	}

	err := Del("test")
	if err != nil {
		assertions.Error(err)
	}
}
