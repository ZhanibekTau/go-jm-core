package redis

import (
	"github.com/ZhanibekTau/go-jm-core/pkg/config/structures"
	"github.com/gomodule/redigo/redis"
	"time"
)

func GetPool(config *structures.RedisConfig) *redis.Pool {
	return &redis.Pool{
		IdleTimeout:     time.Duration(config.IdleTimeout) * time.Second,
		MaxConnLifetime: time.Duration(config.MaxConnLifetime) * time.Second,
		MaxActive:       config.PoolSize,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.RedisHost)
			if err != nil {
				return nil, err
			}

			return c, nil
		},
	}
}

func GetString(pool *redis.Pool, key string) (string, error) {
	conn := pool.Get()
	defer conn.Close()

	return redis.String(conn.Do("GET", key))
}

func SetString(pool *redis.Pool, key string, data interface{}, ttl int) (interface{}, error) {
	conn := pool.Get()
	defer conn.Close()

	return conn.Do("SETEX", key, ttl, data)
}
