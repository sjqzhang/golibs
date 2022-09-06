package redis

import (
	redis6 "github.com/go-redis/redis/v8"
	"github.com/gomodule/redigo/redis"
	"github.com/sjqzhang/golibs/dsnparse"
	"strconv"
	"sync"
)

var globalRedisPool *redis.Pool
var once sync.Once

var globalRedisClient *redis6.Client
var onceClient sync.Once

func InitGlobalRedisPool(dsn string) (*redis.Pool, error) {
	var err error
	once.Do(func() {
		globalRedisPool, err = NewRedisPool(dsn)

	})
	if err != nil {
		return nil, err
	}
	return globalRedisPool, nil
}

func NewRedisPool(dsn string) (*redis.Pool, error) {
	dp, err := dsnparse.Parse(dsn)
	if err != nil {
		return nil, err
	}

	localRedisPool := &redis.Pool{
		// Other pool configuration not shown in this example.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", dp.HostWithPort())
			if err != nil {
				return nil, err
			}
			if dp.Password() != "" {
				if _, err := c.Do("AUTH", dp.Password()); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", dp.DatabaseName()); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
	return localRedisPool, nil

}

func GetGlobalRedisPool() *redis.Pool {
	return globalRedisPool
}

func InitGlobalRedisClient(dsn string) (*redis6.Client, error) {
	var err error
	onceClient.Do(func() {
		globalRedisClient, err = NewRedisClient(dsn)
	})
	if err != nil {
		return nil, err
	}
	return globalRedisClient, nil
}

func NewRedisClient(dsn string) (*redis6.Client, error) {
	dp, err := dsnparse.Parse(dsn)
	if err != nil {
		return nil, err
	}
	db, err := strconv.Atoi(dp.DatabaseName())
	if err != nil {
		return nil, err
	}
	return redis6.NewClient(&redis6.Options{
		Addr:     dp.HostWithPort(),
		Password: dp.Password(),
		Username: dp.Username(),
		DB:       db,
	}), nil
}

func GetGlobalRedisClient() *redis6.Client {
	return globalRedisClient
}
