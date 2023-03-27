// @program:     LearnGo
// @file:        cache_test.go
// @create:      2023-03-18 21:24
// @description:

package demo3

import (
	"fmt"
	"testing"
)

func TestCacheFactory_Create(t *testing.T) {
	cacheFactory := &CacheFactory{}
	redis, error := cacheFactory.Create(redis)
	if error != nil {
		t.Error(error)
	}

	redis.Set("k1", "v1")
	fmt.Println(redis.Get("k1"))
}
