package cache

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	Name string
	Age  int
}

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
	var test Test
	err := Get(r, "test", &test)
	if err != nil {
		assertions.Error(err)
	}
}

func TestRedisStore_Set(t *testing.T) {
	assertions := assert.New(t)
	r := NewRedisStore("127.0.0.1:6379", "", 0)
	if r == nil {
		assertions.Error(errors.New("NewRedisStore error"))
	}
	var test Test
	var test2 Test
	test = Test{
		Name: "test",
		Age:  18,
	}

	err := Set(r, "test", test, 0)
	if err != nil {
		assertions.Error(err)
	}
	err = Get(r, "test", &test2)
	if err != nil {
		assertions.Error(err)
	}
	fmt.Printf("%+v\n", test)
	fmt.Printf("%+v\n", test2)
	assertions.Equal(test, test2)
}
