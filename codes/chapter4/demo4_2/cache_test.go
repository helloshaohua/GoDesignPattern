package demo4_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheFactory_Create(t *testing.T) {
	c := &CacheFactory{}

	// Test RedisCache
	redis, err := c.Create(CacheTypeOfRedis)
	assert.NoError(t, err)
	redis.Set("hello", "world")
	hello, err := redis.Get("hello")
	assert.NoError(t, err)
	assert.Equal(t, "world", hello)

	memcache, err := c.Create(CacheTypeOfMemcache)
	assert.NoError(t, err)
	memcache.Set("foo", "bar")
	bar, err := memcache.Get("foo")
	assert.NoError(t, err)
	assert.Equal(t, "bar", bar)
}
