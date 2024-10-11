package RedisUtil

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var c redis.Conn

func InitRedisConnector() {
	var err error
	c, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
}

func GetRedisClint() redis.Conn {
	return c
}
