package cache

import (
	ec "github.com/foxiswho/echo-go/middleware/cache"
	. "github.com/foxiswho/echo-go/conf"
	"time"
)

var client ec.CacheStore

func NewClient() ec.CacheStore {
	var store ec.CacheStore

	switch Conf.CacheStore {
	case MEMCACHED:
		store = ec.NewMemcachedStore([]string{Conf.Memcached.Server}, time.Hour)
	case REDIS:
		store = ec.NewRedisCache(Conf.Redis.Server, Conf.Redis.Pwd, DefaultExpiration)
	default:
		store = ec.NewInMemoryStore(time.Hour)
	}
	return store
}

func Client() ec.CacheStore {
	if client == nil {
		client = NewClient()
	}
	return client
}
