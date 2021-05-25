package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tal-tech/go-zero/core/jsonx"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/core/syncx"
	"time"
)

const (
	Second = 1
	Minute = Second * 60
	Hour   = Minute * 60
	Day    = Hour * 24
)

var (
	ErrRedisUnmarshal     = errors.New("解析缓存数据出错")
	ErrRedisMarshal       = errors.New("json编码失败")
	ErrRedisPut           = errors.New("缓存写入失败")
	ErrRedisDel           = errors.New("缓存清除失败")
	ErrRedisLock          = errors.New("缓存加锁失败")
	ErrRedisUnLock        = errors.New("缓存解锁失败")
	ErrRedisRelease       = errors.New("缓存释放锁失败")
	ErrRedisOperationFast = errors.New("操作太快请稍后重试")
	ErrNotFound           = errors.New("缓存数据为空")
	ErrDelCache           = errors.New("删除缓存失败")
	ErrGetCache           = errors.New("获取缓存数据出错")
	ErrPutCache           = errors.New("写入缓存失败 ")
	cacheLock             = "cache#lock#key#%v"
	LockKey               = "lock"
)

func Set(r *redis.Redis, key string, data string, expiration int) error {
	if err := r.Setex(key, data, expiration); err != nil {
		return ErrRedisPut
	}
	return nil
}

func SetMarshal(r *redis.Redis, key string, data interface{}, seconds int) error {
	syncx.NewSharedCalls()
	marshal, err := json.Marshal(data)
	if err != nil {
		return ErrRedisMarshal
	}
	if err = r.Setex(key, string(marshal), seconds); err != nil {
		return ErrRedisPut
	}
	return nil
}

func GetUnMarshal(r *redis.Redis, key string, box interface{}) error {
	str, err := r.Get(key)
	if err != nil {
		return ErrRedisUnmarshal
	}
	if str == "" {
		return ErrNotFound
	}
	return json.Unmarshal([]byte(str), box)
}
func Lock(r *redis.Redis, key string, seconds int) (*redis.RedisLock, error) {
	lock := redis.NewRedisLock(r, key)
	lock.SetExpire(seconds)
	acquire, err := lock.Acquire()
	if err != nil {
		return nil, ErrRedisLock
	}
	if !acquire {
		return nil, ErrRedisOperationFast
	}
	return lock, nil
}
func LockDefault(r *redis.Redis, key string) (*redis.RedisLock, error) {
	lock := redis.NewRedisLock(r, fmt.Sprint(key, "#lock"))
	lock.SetExpire(10)
	acquire, err := lock.Acquire()
	if err != nil {
		return nil, ErrRedisLock
	}
	if !acquire {
		return nil, ErrRedisOperationFast
	}
	return lock, nil
}
func UnLock(lock *redis.RedisLock) error {
	release, err := lock.Release()
	if err != nil {
		logx.Error(err)
		return ErrRedisUnLock
	}
	if !release {
		return ErrRedisUnLock
	}
	return nil
}

var (
	sharedCalls = syncx.NewSharedCalls()
)

// Take 从缓存获取key数据绑定到v, 如果没有则从query查询并写入缓存
// 本方法使用了sharedCalls确保只有一个协程去数据库查询
func Take(r *redis.Redis, key string, v interface{}, query func(v interface{}) error, expire time.Duration) error {
	// 用barrier来防止缓存击穿，确保一个进程内只有一个请求去加载key对应的数据
	val, fresh, err := sharedCalls.DoEx(key, func() (interface{}, error) {
		if err := GetCache(r, key, v); err != nil {
			if err := query(v); err != nil {
				return nil, err
			}
			// 加入缓存
			if err = SetCacheWithExpire(r, key, v, expire); err != nil {
				logx.Error(err)
			}
		}
		return jsonx.Marshal(v)
	})
	if err != nil {
		return err
	}

	if fresh {
		return nil
	}
	return jsonx.Unmarshal(val.([]byte), v)
}

// SetCacheWithExpire  写入缓存
func SetCacheWithExpire(r *redis.Redis, key string, v interface{}, expire time.Duration) error {
	data, err := jsonx.Marshal(v)
	if err != nil {
		return err
	}
	return r.Setex(key, string(data), int(expire.Seconds()))
}

// GetCache 从缓存获取数据
func GetCache(r *redis.Redis, key string, v interface{}) error {
	data, err := r.Get(key)
	if err != nil {
		logx.Error(err.Error())
		return err
	}
	if len(data) == 0 {
		return ErrNotFound
	}
	if err = jsonx.Unmarshal([]byte(data), v); err == nil {
		return nil
	}
	if _, e := r.Del(key); e != nil {
		logx.Errorf("delete invalid cache,key:%s, value:%s,error:%v", key, data, e)
	}
	return ErrNotFound
}
