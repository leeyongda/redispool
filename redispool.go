package redis

import (
	"github.com/garyburd/redigo/redis"
)

type redisStorage struct {
	addr string
}

var (
	rds redis.Conn
)

func newPool(addr string) *redis.Pool {

	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 3,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
}

func PingRedis(url string) (redis.Conn, error) {

	pool := newPool(url)
	conn := pool.Get()
	_, err := conn.Do("PING")
	if err != nil {
		return nil, err
	}
	return conn, nil

}

// 连接redis 数据库
func ConnectRedis(url string) error {
	_, err := PingRedis(url)
	if err != nil {
		return err
	}
	return nil
}

func GetRedisStore() redis.Conn {
	return rds
}

func Close() error {
	err := rds.Close()
	return err
}
