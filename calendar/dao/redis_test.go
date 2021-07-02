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
