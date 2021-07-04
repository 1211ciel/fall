package dao

import "time"

type (
	Cache interface {
		Del(keys ...string) error
		Get(key string, v interface{}) error
		Set(key string, v interface{}) error
		SetWithExpire(key string, v interface{}, expire time.Duration) error
		Take(v interface{}, key string, query func(v interface{}) error) error
		TakeWithExpire(v interface{}, key string, query func(v interface{}), expire time.Duration) error
	}
)
