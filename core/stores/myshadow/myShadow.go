package myshadow

import (
	"database/sql"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/core/syncx"
	"gorm.io/gorm"
	"time"
)

// see doc/sql-cache.md
const cacheSafeGapBetweenIndexAndPrimary = time.Second * 5

var (
	sharedCalls = syncx.NewSharedCalls()
	stats       = cache.NewStat("myShadow")
)

type (

	// QueryFn defines the query method.
	QueryFn func(db *gorm.DB, v interface{}) error
	//IndexQueryFn 索引查询主键方法 返回主键
	IndexQueryFn func(db *gorm.DB, v interface{}) (interface{}, error)

	PrimaryQueryFn func(db *gorm.DB, v, primary interface{}) error
	Shadow         struct {
		DB    *gorm.DB
		R     *redis.Redis
		Cache cache.Cache
	}
)

// QueryRow unmarshals into v with given key and query func.
func (s Shadow) QueryRow(v interface{}, key string, query QueryFn) error {
	// 从缓存获取,如果没有再从query获取
	return s.Cache.Take(v, key, func(v interface{}) error {
		return query(s.DB, v)
	})
}

// QueryRowIndex
/*使用key去缓存获取值扫描到v,
如果缓存不存在,通过 indexQuery,查到到主键, keyer 用来制作主键对应的key
*/
func (s Shadow) QueryRowIndex(v interface{}, key string, keyer func(primary interface{}) string,
	indexQuery IndexQueryFn, primaryQuery PrimaryQueryFn) error {
	var primaryKey interface{}
	var found bool
	/* 这里的 TakeWithExpire 干的事情
	1. 通过 key 获取主键 绑定到 primaryKey
	2. 如果 缓存中没有找到则通过 indexQuery 方法获取. indexQuery 干的事情 1.查找 主键 ,2绑定值到 v
	3. 上面2步已经获取到主键了, 那么这里就直接通过 keyer 返回的主键缓存key 进行获取.
	*/
	if err := s.Cache.TakeWithExpire(&primaryKey, key, func(val interface{}, expire time.Duration) (err error) {
		primaryKey, err = indexQuery(s.DB, v)
		if err != nil {
			return
		}

		found = true
		return s.Cache.SetWithExpire(keyer(primaryKey), v, expire+cacheSafeGapBetweenIndexAndPrimary)
	}); err != nil {
		return err
	}
	if found {
		return nil
	}
	return s.Cache.Take(v, keyer(primaryKey), func(v interface{}) error {
		return primaryQuery(s.DB, v, primaryQuery)
	})
}

// DelCache 清除缓存
func (s Shadow) DelCache(key ...string) {
	_, err := s.R.Del(key...)
	if err != nil {
		logx.Error(err.Error())
	}
}

func NewNodeConn(db *gorm.DB, r *redis.Redis, opts ...cache.Option) Shadow {
	return Shadow{
		DB:    db,
		R:     r,
		Cache: cache.NewNode(r, sharedCalls, stats, sql.ErrNoRows, opts...),
	}
}
