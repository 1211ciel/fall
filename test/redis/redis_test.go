package redis

import (
	"fmt"
	"github.com/1211ciel/word-of-wind/calendar/dao"
	"github.com/gomodule/redigo/redis"
	"testing"
)

func TestRedis77(t *testing.T) {
	c := dao.NewDefaultRedis(":6379").Get()
	// c
	//fmt.Println(redis.String(c.Do("set", "ciel", "come on")))
	// u
	//fmt.Println(redis.String(c.Do("set", "ciel", "加油")))
	// r
	//fmt.Println(redis.String(c.Do("get","ciel")))
	// d
	//fmt.Println(redis.Int(c.Do("del", "ciel")))
	//hash map c
	//fmt.Println(redis.Int(c.Do("hset", "book", "go", "golang")))
	// hash map r
	//fmt.Println(redis.String(c.Do("hget", "book", "go")))
	fmt.Println(redis.StringMap(c.Do("hgetall", "book")))
}
func TestRedis78(t *testing.T) {
	c := dao.NewDefaultRedis(":6379").Get()
	// c
	//fmt.Println(c.Do("set", "ciel", "123"))
	// u
	//fmt.Println(c.Do("set","ciel","234"))
	// r
	//fmt.Println(redis.Int(c.Do("get", "ciel")))
	// d
	//fmt.Println(redis.Int(c.Do("del", "ciel")))
	// hash map c
	//fmt.Println(redis.StringMap(c.Do("hset", "books", "go", "123")))
	// hash map r
	//fmt.Println(redis.String(c.Do("hget", "books", "go")))
	fmt.Println(redis.Int(c.Do("del", "books")))
}
