package appredis

import (
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	size      int           = 1000
	cacheTime time.Duration = 10 * time.Minute
)

type Object struct {
	Str string
	Num int
}

func GetCache() *cache.Cache {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server": ":6379",
		},
	})
	return cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(size, cacheTime),
	})
}
