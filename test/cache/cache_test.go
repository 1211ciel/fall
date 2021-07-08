package cache

import (
	"fmt"
	model "github.com/1211ciel/fall/model/user"
	"github.com/1211ciel/word-of-wind/calendar/dao"
	"testing"
)

func TestCache77(t *testing.T) {
	cache := dao.NewCache(dao.NewDefaultRedis(":6379"))
	for i := 0; i < 10; i++ {
		var u model.User
		err := cache.Take(&u, "ciel", func(v interface{}) error {
			fmt.Println("123")
			temp := model.NewUser(2, "ciel", "icon.png", "123", "1234556")
			*v.(*model.User) = temp
			return nil
		})
		if err != nil {
			t.Fatal(err.Error())
		}
		fmt.Println(u)
	}
}
func TestCache78(t *testing.T) {
	cache := dao.NewCache(dao.NewDefaultRedis(":6379"))
	for i := 0; i < 10; i++ {
		var u model.User
		err := cache.Take(&u, "ciel", func(v interface{}) error {
			fmt.Println("came in")
			temp := model.NewUser(2, "ciel", "icon.png", "123", "12233322111")
			*v.(*model.User) = temp
			return nil
		})
		if err != nil {
			t.Fatal(err.Error())
		}
		fmt.Println(u)
	}
}
