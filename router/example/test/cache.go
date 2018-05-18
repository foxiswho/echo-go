package test

import (
	"github.com/foxiswho/shop-go/router/base"
	"github.com/foxiswho/shop-go/module/cache"
	"time"
	"fmt"
)

func CacheHandler(c *base.BaseContext) error {

	err := cache.Client().Set("test", time.Now(), 10*time.Minute)
	fmt.Println("err", err)
	var cache_test time.Time
	err = cache.Client().Get("test", &cache_test)
	fmt.Println("get test err", err,cache_test)
	c.Set("tmpl", "example/test/cache")
	c.Set("data", map[string]interface{}{
		"title":      "测试 缓存",
		"cache_test": cache_test,
	})

	return nil
}
