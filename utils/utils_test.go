package utils

import (
	"errors"
	"fmt"
	"github.com/1211ciel/fall/utils/dbutil"
	"github.com/tal-tech/go-zero/core/bloom"
	"github.com/tal-tech/go-zero/core/collection"
	"github.com/tal-tech/go-zero/core/fx"
	"github.com/tal-tech/go-zero/core/mr"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/core/stringx"
	"gorm.io/gorm"
	"strings"
	"testing"
	"time"
)

func TestInput(t *testing.T) {
}

func TestString(t *testing.T) {
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (receiver User) TableName() string {
	return "aoo"
}

var (
	tw    *collection.TimingWheel
	count = 0
	key   = "t1"
)

func TestBloom(t *testing.T) {
	store := redis.NewRedis("localhost:6379", redis.NodeType, "")
	bitSet := bloom.New(store, "test_key", 2)
	key = "dd"
	_ = bitSet.Add([]byte(key))
	_ = bitSet.Add([]byte("ok"))
	_ = bitSet.Add([]byte("ok3"))
	_ = bitSet.Add([]byte("ok5"))
	_ = bitSet.Add([]byte("ok6"))
	exists, _ := bitSet.Exists([]byte(key))
	fmt.Println(exists)

}

// 定时任务
func TestTimeWheel2(t *testing.T) {
	timingWheel, err := collection.NewTimingWheel(time.Second, 10, func(key, value interface{}) {
		count++
		fmt.Println(count, " ", time.Now())
		tw.SetTimer(key, "", time.Second*3)
		if count > 5 {
			tw.RemoveTimer(key) //
			tw.SetTimer(key, "", time.Second*1)
		}
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	tw = timingWheel
	tw.SetTimer(key, "", time.Second*1)
	select {}
}

// 进程内缓存使用
func TestCache(t *testing.T) {
	c, err := collection.NewCache(time.Minute, collection.WithLimit(10000))
	if err != nil {
		panic(err)
	}
	take, err := c.Take("key", func() (interface{}, error) {
		return User{Name: "ciel", Age: 18}, nil
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(take)
	c.Del("key")

}

// 查找敏感词
func TestStringX(t *testing.T) {
	filter := stringx.NewTrie([]string{
		"苍井空",
		"AV",
	})
	keywords := filter.FindKeywords("日本AV演员兼电视、电影演员。苍井空AV女优是xx出道, 日本AV女优们最精彩的表演是AV演员色情表演")
	fmt.Println(keywords)
}

// 敏感词过滤
func TestStringX2(t *testing.T) {
	filter := stringx.NewTrie([]string{
		"AV演员",
		"苍井空",
		"AV",
		"日本AV女优",
		"AV演员色情",
	}, stringx.WithMask('?')) // 默认替换为*
	safe, keywords, found := filter.Filter("日本AV演员兼电视、电影演员。苍井空AV女优是xx出道, 日本AV女优们最精彩的表演是AV演员色情表演")
	fmt.Println(safe)
	fmt.Println(keywords)
	fmt.Println(found)
}
func TestMapReduce(t *testing.T) {
	uids := []int64{1, 2, 3, 4, 5, 6, 7, 8}
	r, err := mr.MapReduce(func(source chan<- interface{}) {
		for _, item := range uids {
			source <- item
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		if item.(int64) > 10 {
			cancel(errors.New(fmt.Sprint(item, "大于10")))
		} else {
			writer.Write(item)
		}
	}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
		var newUids []int64
		for p := range pipe {
			newUids = append(newUids, p.(int64))
		}
		writer.Write(newUids)
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(r)

}
func TestMapReduceFinish(t *testing.T) {
	var u User
	err := mr.Finish(func() error {
		u.Name = "ciel"
		return nil
	}, func() error {
		u.Age = 17
		return errors.New("err")
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(u)
}
func TestFuc(t *testing.T) {
	var u User
	if err := GetValue(&u, func(v interface{}) error {
		// 模拟 query 查询
		a := User{Name: "john", Age: 18}
		// 请问这里如何将 a的数据 绑定到 v上面去呢? 最后返回给u
		*v.(*User) = a
		return nil
		//return nil
	}); err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(u)
}

// 模拟获取数据方法  先从缓存获取,如果没有再从 query 查询方法获取
func GetValue(v interface{}, query func(v interface{}) error) error {
	// 从缓存获取,如果没有从 query 方法查询获取
	return query(v)
}

// 流数据 相加
func TestFxSum(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	sum := 0
	fx.From(func(source chan<- interface{}) {
		for _, v := range s {
			source <- v
		}
	}).Filter(func(item interface{}) bool {
		if item.(int)%2 == 0 {
			return true
		}
		return false
	}).ForEach(func(item interface{}) {
		sum += item.(int)
	})
	fmt.Println(sum)
}

// 流数据分组
func TestFxGroup(t *testing.T) {
	ss := []string{"golang", "google", "php", "python", "java", "c++"}
	fx.From(func(source chan<- interface{}) {
		for _, s := range ss {
			source <- s
		}
	}).Group(func(item interface{}) interface{} {
		if strings.HasPrefix(item.(string), "g") {
			return "g"
		} else if strings.HasPrefix(item.(string), "p") {
			return "p"
		}
		return ""
	}).ForEach(func(item interface{}) {
		fmt.Printf("%T %v\n", item, item)
	})
}

// 流数据切分去重
func TestFxSplit(t *testing.T) {
	fx.Just(3, 4, 551, 23, 1, 3, 4, 5).
		Sort(func(a, b interface{}) bool {
			return a.(int)-b.(int) > 0
		}).
		// 去重
		Distinct(func(item interface{}) interface{} {
			return item
		}).
		Split(3). // 切分 每组3个
		ForEach(func(item interface{}) {
			fmt.Println(item)
		})
}

// 流数据并发处理
func TestFxWalk(t *testing.T) {
	// 例子
	var sum int
	fx.Just(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).Walk(func(item interface{}, pipe chan<- interface{}) {
		pipe <- item.(int) * 10
	}).ForEach(func(item interface{}) {
		sum += item.(int)
	})
	fmt.Println(sum)
}
func TestDB(t *testing.T) {
	var a User
	if err := getDB().Table("aoo").First(&a, "name = 'ciel'").Error; err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(a)
}
func getDB() *gorm.DB {
	return dbutil.GetDefaultMysql("root:123456@tcp(localhost:3306)/sky?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai", true)
}
