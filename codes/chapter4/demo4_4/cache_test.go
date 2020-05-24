package demo4_4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemcacheFactory_Create(t *testing.T) {
	var cacheFactory CacheFactory
	cacheFactory = &MemcacheFactory{}
	memcache := cacheFactory.Create()
	memcache.Set("hello", "world")
	world, err := memcache.Get("hello")
	assert.NoError(t, err)
	assert.Equal(t, "world", world)
}

func TestRedisCacheFactory_Create(t *testing.T) {
	var cacheFactory CacheFactory
	cacheFactory = &RedisCacheFactory{}
	redis := cacheFactory.Create()
	redis.Set("hello", "world")
	world, err := redis.Get("hello")
	assert.NoError(t, err)
	assert.Equal(t, "world", world)
}
