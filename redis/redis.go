package redis

import (
	"time"

	"github.com/garyburd/redigo/redis"

	"fmt"
	"redis-service/config"
	"strconv"
	"sync"
)

var (
	pool *redis.Pool
	wg   sync.WaitGroup
)

func init() {
	host := config.GetString("redis_host", true, "127.0.0.1")
	port := config.GetString("redis_port", true, "6379")
	pass := config.GetString("redis_auth", true, "")
	maxIdle := config.GetString("redis_maxidle", true, "10")
	server := fmt.Sprint(host, ":", port)
	maxIdleInt, err := strconv.Atoi(maxIdle)
	if err != nil {
		panic(err)
	}
	pool = newPool(server, pass, maxIdleInt)
}

func newPool(server, password string, maxIdle int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func Exec(commandName string, args ...interface{}) (interface{}, error) {
	defer wg.Done()
	wg.Add(1)
	reply, err := pool.Get().Do(commandName, args...)
	return reply, err
}

func Select(db string) {
	Exec(db)
}

func Close() {
	wg.Wait()
	pool.Close()
}
