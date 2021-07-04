package dao

import (
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/syncx"
	"sync"
	"time"
)

var (
	sharedCalls     = syncx.NewSharedCalls()
	errDataNotFound = errors.New("data not found")
	errSetData      = errors.New("set data failed")
	defaultExpire   = time.Hour * 24 * 2
)

type Shadow struct {
	pool *redis.Pool
	lock *sync.Mutex
}

func NewShadow(pool *redis.Pool) Shadow {
	return Shadow{
		pool: pool,
		lock: new(sync.Mutex),
	}
}

func (s Shadow) Del(keys ...string) error {
	c := s.pool.Get()
	err := c.Send("MULTI")
	if err != nil {
		logx.Error(err.Error())
		return err
	}
	for _, key := range keys {
		err = c.Send("del", key)
		if err != nil {
			return err
		}
	}
	_, err = c.Do("EXEC")
	return err
}

func (s Shadow) Get(key string, v interface{}) error {
	c := s.pool.Get()
	data, err := redis.String(c.Do("get", key))
	if err != nil {
		return errDataNotFound
	}
	err = json.Unmarshal([]byte(data), v)
	if err == nil {
		return nil
	}
	return errDataNotFound
}

// Set 默认7天
func (s Shadow) Set(key string, v interface{}) error {
	return s.SetWithExpire(key, v, defaultExpire)
}

// SetWithExpire expire unit is seconds
func (s Shadow) SetWithExpire(key string, v interface{}, expire time.Duration) error {
	data, err := json.Marshal(v)
	if err != nil {
		logx.Error(err.Error())
		return err
	}
	ok, err := redis.String(s.pool.Get().Do("SETEX", key, expire.Seconds(), string(data)))
	if err != nil {
		return err
	}
	if ok == "ok" {
		return errSetData
	}
	return nil
}

// Take  用key 去缓存获取数据绑定到v,如果不存在,则使用 query进行查询
func (s Shadow) Take(v interface{}, key string, query func(v interface{}) error) error {
	return s.TakeWithExpire(v, key, query, defaultExpire)
}

func (s Shadow) TakeWithExpire(v interface{}, key string, query func(v interface{}) error, expire time.Duration) error {
	val, fresh, err := sharedCalls.DoEx(key, func() (interface{}, error) {
		// 先从缓存获取
		err := s.Get(key, v)
		if err == nil {
			return nil, err
		}
		if err != errDataNotFound {
			return nil, err
		}
		// 没有再从query获取
		if err = query(v); err != nil {
			logx.Error(err.Error())
			return nil, err
		}
		// 获取到了再放缓存
		err = s.SetWithExpire(key, v, expire)
		if err != nil {
			return nil, err
		}
		return json.Marshal(v)
	})
	if err != nil {
		logx.Error(err.Error())
		return err
	}
	if fresh {
		return nil
	}
	return json.Unmarshal(val.([]byte), v)
}
