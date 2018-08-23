package redis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

// Config is a conf for the redis KV store
type Config struct {
	Host                 string
	Port                 string
	Password             string
	MaxActiveConnections int
}

// NewConnPool connects to redis and return a connection pool
func NewConnPool(redisConf Config) (*redis.Pool, error) {
	redPool := &redis.Pool{
		MaxIdle:     3,
		MaxActive:   redisConf.MaxActiveConnections,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisConf.Host+":"+redisConf.Port)
			if err != nil {
				return nil, err
			}
			if redisConf.Password != "" {
				if _, errDo := c.Do("AUTH", redisConf.Password); errDo != nil {
					err2 := c.Close()
					return nil, fmt.Errorf("redis: %v, close err? %v", err, err2)
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	if err := ConnPoolPing(redPool); err != nil {
		return nil, fmt.Errorf("NewConnPool: %v", err)
	}

	return redPool, nil
}

// ConnPoolPing allows a check on redis server status
func ConnPoolPing(redPool *redis.Pool) error {
	var err error
	red := redPool.Get()
	defer func() {
		err = red.Close()
	}()

	_, err = red.Do("PING")
	if err != nil {
		return fmt.Errorf("ConnPoolPing: %v", err)
	}

	return nil
}
