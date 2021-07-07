package redis

import (
	"fmt"
	"github.com/1211ciel/word-of-wind/calendar/dao"
	"github.com/gomodule/redigo/redis"
	"testing"
)

func TestRedis77(t *testing.T) {
	c := dao.NewDefaultRedis(":6379").Get()
	//fmt.Println(redis.String(c.Do("set", "ciel", "come on"))) c
	//fmt.Println(redis.String(c.Do("set", "ciel", "加油"))) u
	//fmt.Println(redis.String(c.Do("get","ciel")))  r
	//fmt.Println(redis.Int(c.Do("del", "ciel"))) // d
	//fmt.Println(redis.Int(c.Do("hset", "book", "go", "golang")))
	//fmt.Println(redis.String(c.Do("hget", "book", "go")))
	fmt.Println(redis.StringMap(c.Do("hgetall", "book")))
}
