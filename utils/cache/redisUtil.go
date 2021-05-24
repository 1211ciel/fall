package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/redis"
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
	ErrRedisDataEmpty     = errors.New("缓存数据为空")
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
	marshal, err := json.Marshal(data)
	if err != nil {
		return ErrRedisMarshal
	}
	if err = r.Setex(key, string(marshal), int(seconds)); err != nil {
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
		return ErrRedisDataEmpty
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
