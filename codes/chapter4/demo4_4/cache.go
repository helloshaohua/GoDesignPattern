package demo4_4

import (
	"errors"
	"sync"
)

// ------------------------------------------------------------------------
// 定义一个抽象的Cache接口
// ------------------------------------------------------------------------
type Cache interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, error)
}

// ------------------------------------------------------------------------
// 实现具体的Cache: RedisCache
// ------------------------------------------------------------------------
type RedisCache struct {
	data map[string]interface{}
	mux  sync.RWMutex
}

func (r *RedisCache) Set(key string, value interface{}) {
	r.mux.Lock()
	defer r.mux.Unlock()
	r.data[key] = value
	return
}

func (r *RedisCache) Get(key string) (interface{}, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()
	if v, ok := r.data[key]; ok {
		return v, nil
	}
	return nil, errors.New("key is not found")
}

// ------------------------------------------------------------------------
// 实现具体的Cache: Memcache
// ------------------------------------------------------------------------
type Memcache struct {
	data map[string]interface{}
	mux  sync.RWMutex
}

func (m *Memcache) Set(key string, value interface{}) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.data[key] = value
	return
}

func (m *Memcache) Get(key string) (interface{}, error) {
	m.mux.RLock()
	defer m.mux.RUnlock()
	if v, ok := m.data[key]; ok {
		return v, nil
	}
	return nil, errors.New("key is not found")
}

// ------------------------------------------------------------------------
// 定义一个抽象的缓存工厂
// ------------------------------------------------------------------------
type CacheFactory interface {
	Create() Cache
}

// ------------------------------------------------------------------------
// 实现具体的缓存工厂
// ------------------------------------------------------------------------
type RedisCacheFactory struct{}

func (r *RedisCacheFactory) Create() Cache {
	return &RedisCache{data: make(map[string]interface{})}
}

type MemcacheFactory struct{}

func (m *MemcacheFactory) Create() Cache {
	return &Memcache{data: make(map[string]interface{})}
}
