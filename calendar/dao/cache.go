package dao

import (
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/syncx"
	"gorm.io/gorm"
	"sync"
	"time"
)

const (
	notFoundPlaceholder = "*"
	expiryDeviation     = 0.05
)

var (
	sharedCalls = syncx.NewSharedCalls()
)

// indicates there is no such value associate with the key
var errPlaceholder = errors.New("placeholder")

type Shadow struct {
	pool        *redis.Pool
	lock        *sync.Mutex
	db          *gorm.DB
	errNotFound error
}

func NewShadow(pool *redis.Pool, db *gorm.DB) Shadow {
	return Shadow{
		pool:        pool,
		lock:        new(sync.Mutex),
		errNotFound: errPlaceholder,
		db:          db,
	}
}

func (s Shadow) Del(keys ...string) error {
	c := s.pool.Get()
	c.Send("MULTI")
	for _, key := range keys {
		c.Send("del", key)
	}
	_, err := c.Do("EXEC")
	return err
}

func (s Shadow) Get(key string, v interface{}) error {
	c := s.pool.Get()
	c.Send("get", key)
	c.Flush()
	data, err := redis.String(c.Receive())
	if err != nil {
		return s.errNotFound
	}
	err = json.Unmarshal([]byte(data), v)
	if err == nil {
		return nil
	}
	return s.errNotFound
}

func (s Shadow) IsNotFound(err error) bool {
	return s.errNotFound == err
}

// Set 默认7天
func (s Shadow) Set(key string, v interface{}) error {
	return s.SetWithExpire(key, v, time.Hour*24*7)
}

// SetWithExpire expire unit is seconds
func (s Shadow) SetWithExpire(key string, v interface{}, expire time.Duration) error {
	data, err := json.Marshal(v)
	if err != nil {
		logx.Error(err.Error())
		return err
	}
	c := s.pool.Get()
	c.Send("SETEX", key, expire.Seconds(), string(data))
	c.Flush()
	_, err = c.Receive()
	if err != nil {
		logx.Error(err.Error())
		return err
	}
	return nil
}

// Take  用key 去缓存获取数据绑定到v,如果不存在,则使用 query进行查询
func (s Shadow) Take(v interface{}, key string, query func(v interface{}) error) error {
	return s.doTake(v, key, query, func(v interface{}) error {
		return s.Set(key, v)
	})
}
func (s Shadow) doTake(v interface{}, key string, query func(v interface{}) error,
	cacheVal func(v interface{}) error) error {
	// 确保只有一个执行查询
	val, fresh, err := sharedCalls.DoEx(key, func() (interface{}, error) {
		// 想从缓存获取
		if err := s.Get(key, v); err != nil {
			if err != errPlaceholder {
				logx.Error(err.Error())
				return nil, s.errNotFound
			}
			// 如果缓存不存在则从 query获取
			if err = query(v); err != nil {

				logx.Error(err)
				return nil, err
			}
			// 将查询到的结果 通过 cacheVal方法设置到缓存
			if err = cacheVal(v); err != nil {
				logx.Error(err.Error())
			}
		}
		// 返回
		return json.Marshal(v)
	})
	// 如果出错返回错误
	if err != nil {
		return err
	}
	// 如果已经序列化则不需要在序列化一次
	if fresh {
		return nil
	}
	// 反序列化
	return json.Unmarshal(val.([]byte), v)
}

func (s Shadow) TakeWithExpire(v interface{}, key string, query func(v interface{}, expire time.Duration) error) error {
	return nil
}
