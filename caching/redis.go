package caching

import "github.com/redis/go-redis/v9"

// CacheDriver ...
type CacheDriver struct {
	addr       string
	pwd        string
	db         int
	maxRetries int
}

// NewCacheDriver factory to initialize CacheDriver
func NewCacheDriver(addr, pwd string, db, maxRetries int) *CacheDriver {
	return &CacheDriver{
		addr:       addr,
		pwd:        pwd,
		db:         db,
		maxRetries: maxRetries,
	}
}

func (d CacheDriver) GetClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:       d.addr,
		Password:   d.pwd,
		DB:         d.db,
		MaxRetries: d.maxRetries,
	})
	return client
}
