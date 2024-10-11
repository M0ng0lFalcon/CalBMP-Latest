package RedisUtil

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func GetStringVal(keyName string) string {
	Val, err := redis.String(c.Do("GET", keyName))
	if err != nil {
		fmt.Println("redis get failed:", err)
	}
	return Val
}

func SetStringKey(KeyName string, Val string) {
	_, err := c.Do("SET", KeyName, Val)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}

func SetFloatKey(KeyName string, Val float64) {
	_, err := c.Do("SET", KeyName, Val)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}

func GetFloatVal(KeyName string) float64 {
	val, err := redis.Float64(c.Do("GET", KeyName))
	if err != nil {
		fmt.Println("redis get failed:", err)
		val = -1
	}
	return val
}

func SetIntVal(KeyName string, Val int) {
	_, err := c.Do("SET", KeyName, Val)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}

func GetIntVal(KeyName string) int {
	Val, err := redis.Int(c.Do("GET", KeyName))
	if err != nil {
		fmt.Println("redis get failed:", err)
	}
	return Val
}
