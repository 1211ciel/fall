package dao

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
)

func TestRedisC(t *testing.T) {
	pool := NewDefaultRedis("127.0.0.1:6379")
	c := pool.Get()
	err := c.Send("SET", "ciel", "hello")
	if err != nil {
		t.Fatal(err.Error())
	}
	c.Flush()
}
func TestRedisU(t *testing.T) {
	c := NewDefaultRedis(":6379").Get()
	err := c.Send("get", "ciel")
	if err != nil {
		t.Fatal(err.Error())
	}
	c.Flush()
	s, err := redis.String(c.Receive())
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf(s)
}
func TestRedisD(t *testing.T) {
	c := NewDefaultRedis(":6379").Get()
	c.Send("del", "ciel")
	c.Flush()
	fmt.Println(redis.String(c.Receive()))
}
func TestRedis202172(t *testing.T) {
	// add
	c := NewDefaultRedis(":6379").Get()
	//c.Send("set", "ciel", "come on ")
	//c.Flush()
	//c.Send("get", "ciel")
	//c.Flush()
	//s, err := redis.String(c.Receive())
	//if err != nil {
	//	t.Fatal(err.Error())
	//}
	//fmt.Printf(s)
	// 批量删除
	c.Send("MULTI")
	c.Send("del", "1")
	c.Send("del", "2")
	c.Send("del", "3")
	c.Send("EXEC")
	c.Do("EXEC")
	//c.Flush()  do 和 flush 都可以使用
}
func TestRedis1(t *testing.T) {
	c := NewDefaultRedis(":6379").Get()
	// 基本
	// c
	//c.Do("set", "hello", "123")
	// u
	//c.Do("set", "hello", "1234")
	// r
	//content, err := redis.String(c.Do("get", "hello"))
	//if err != nil {
	//	t.Fatal(err.Error())
	//}
	//fmt.Printf(content)
	// d
	//c.Do("del", "hello")
	// 哈希表
	// c u
	//s, err := redis.Int(c.Do("hset", "website", "google", "www.g.cn"))
	//if err != nil {
	//	t.Fatal(err.Error())
	//}
	//fmt.Println(s)
	// r
	//s, err2 := redis.String(c.Do("hget", "website", "google"))
	//if err2 != nil {
	//	t.Fatal(err2.Error())
	//}
	//fmt.Println(s)

	// r all
	//stringMap, err := redis.StringMap(c.Do("hgetall", "website"))
	//if err != nil {
	//	t.Fatal(err.Error())
	//}
	//fmt.Println(stringMap)
	c.Do("hdel", "website", "google")
}

func TestRedis74(t *testing.T) {
	c := NewDefaultRedis(":6379").Get()
	// c
	//i, err := redis.String(c.Do("set", "ciel", "sunnie"))
	//if err != nil {
	//	t.Fatal(err.Error())
	//}
	// u
	//s, err := redis.String(c.Do("set", "ciel", "goodbye"))
	//if err != nil {
	//	t.Fatal(err.Error())
	//}
	//fmt.Println(s)
	// d
	//s, err := redis.Int(c.Do("del", "ciel"))
	//if err != nil {
	//	t.Fatal(err.Error())
	//}
	//fmt.Println(s)

	// hash map
	//c.Do("hset", "books", "golang", 1)
	//i, err := redis.Int(c.Do("hget", "books", "golang"))
	//if err != nil {
	//	t.Fatal(err.Error())
	//}
	//fmt.Println(i)
	//c.Do("hset", "books", "java", "3")
	i, err := redis.Int(c.Do("hdel", "books", "golang"))
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(i)
}
