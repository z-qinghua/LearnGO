// @program:     LearnGo
// @file:        cache.go
// @create:      2023-03-18 21:24
// @description:

package demo3

import "errors"

//定义一个Cache接口，作为父类

type Cache interface {
	Set(key, value string)
	Get(key string) string
}

// 实现具体的Cache：RedisCache

type RedisCache struct {
	data map[string]string
}

func (redis *RedisCache) Set(key, value string) {
	redis.data[key] = value
}

func (redis *RedisCache) Get(key string) string {
	return "redis: " + redis.data[key]
}

//实现具体的Cache：MemCache

type MemCache struct {
	data map[string]string
}

func (mem *MemCache) Set(key, value string) {
	mem.data[key] = value
}

func (mem *MemCache) Get(key string) string {
	return "mem: " + mem.data[key]
}

type cacheType int

const (
	redis cacheType = iota
	mem
)

//实现cache的简单工厂

type CacheFactory struct {
}

func (factory *CacheFactory) Create(cacheType cacheType) (Cache, error) {
	if cacheType == redis {
		return &RedisCache{
			data: map[string]string{}, //map需要初始化
		}, nil
	}
	if cacheType == mem {
		return &MemCache{
			data: map[string]string{},
		}, nil
	}
	return nil, errors.New("error cache type")
}
