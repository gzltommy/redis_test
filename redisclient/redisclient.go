package redisclient

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var redisPool *redis.Pool

func init() {
	redisHost := "192.168.24.147"
	redisPort := "6379"
	password := ""

	url := fmt.Sprintf("%s:%s", redisHost, redisPort)
	redisPool = &redis.Pool{
		TestOnBorrow:    nil,
		MaxIdle:         20,
		MaxActive:       100,
		IdleTimeout:     20 * time.Second,
		Wait:            false,
		MaxConnLifetime: 0,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", url)
			if err != nil {
				return nil, err
			}
			if password != "" {
				_, err = c.Do("AUTH", password)
				if err != nil {
					return nil, err
				}
			}
			return c, nil
		},
	}

	return
}

func Get() redis.Conn {
	return redisPool.Get()
}

func Close() error {
	if redisPool == nil {
		return nil
	}
	if err := redisPool.Close(); err != nil {
		fmt.Printf("redis close err:%v", err)
		return err
	} else {
		fmt.Println("redis close")
	}
	return nil
}
