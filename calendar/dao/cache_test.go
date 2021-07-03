package dao

import (
	"fmt"
	model "github.com/1211ciel/fall/model/user"
	"testing"
)

func TestShadow202172(t *testing.T) {
	db := NewDefaultDB("root:123456@tcp(localhost:3306)/ciel_21_7_2?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai", true)
	shadow := NewShadow(NewDefaultRedis("127.0.0.1:6379"))
	for i := 0; i < 10; i++ {
		var u model.User
		err := shadow.Take(&u, "ciel", func(v interface{}) error {
			var temp model.User
			err := db.Model(&u).Where("id = 2").First(&temp).Error
			if err != nil {
				return err
			}
			*v.(*model.User) = temp
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(u)
	}
}
func TestCache2(t *testing.T) {
	shadow := NewShadow(NewDefaultRedis(":6379"))
	for i := 0; i < 10; i++ {
		var u model.User
		err := shadow.Take(&u, "ciel", func(v interface{}) error {
			fmt.Println("进来了啊")
			temp := model.User{Pid: 1, Uname: "ciel", Icon: "icon", Phone: "1211"}
			*v.(*model.User) = temp
			return nil
		})
		if err != nil {
			return
		}
		fmt.Println(u)
	}
}
