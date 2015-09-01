package game_web_server

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

/*
	使用方法
		var pool *redis.Pool
 		pool = newPool(":6379")
 		conn := pool.Get()
 		defer conn.Close()
 		size, _ := redis.Int(conn.Do("DBSIZE"))
		print("db size is:", size)

*/
func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
