package dao

import (
	"fmt"
	model "github.com/1211ciel/fall/model/user"
	"testing"
	"time"
)

func TestShadow_TakeWithExpire(t *testing.T) {
	shadow := NewShadow(NewDefaultRedis(":6379"))
	for i := 0; i < 10; i++ {
		var u model.User
		err := shadow.TakeWithExpire(&u, "ciel", func(v interface{}) error {
			temp := model.User{
				Pid: 1, Uname: "ciel", Icon: "test222.png", Pwd: "123", Phone: "1222111",
			}
			*v.(*model.User) = temp
			return nil
		}, time.Second*10)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(u)
	}
}
func TestShadow_Take(t *testing.T) {
	shadow := NewShadow(NewDefaultRedis(":6379"))
	for i := 0; i < 10; i++ {
		var u model.User
		if err := shadow.Take(&u, "ciel", func(v interface{}) error {
			fmt.Println("i'm coming")
			temp := model.User{
				Pid: 1, Uname: "ciel", Icon: "test.png", Pwd: "123123", Phone: "15322311233",
			}
			*v.(*model.User) = temp
			return nil
		}); err != nil {
			t.Fatal(err.Error())
		}
		fmt.Println(u)
	}
}
func TestShadow_SetWithExpire(t *testing.T) {
	shadow := NewShadow(NewDefaultRedis(":6379"))
	err := shadow.SetWithExpire("s", "ciel", time.Second*10)
	if err != nil {
		t.Fatal(err.Error())
	}
}
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
func TestCache3(t *testing.T) {
	shadow := NewShadow(NewDefaultRedis(":6379"))
	for i := 0; i < 10; i++ {
		var u model.User
		err := shadow.Take(&u, "s", func(v interface{}) error {
			fmt.Println("进来啦")
			temp := model.User{Pid: 0, Uname: "s", Icon: "test.png", Pwd: "123", Phone: "15233112233"}
			*v.(*model.User) = temp
			return nil
		})
		if err != nil {
			t.Fatal(err.Error())
		}
		fmt.Println(u)
	}
}
func TestS75(t *testing.T) {
	shadow := NewShadow(NewDefaultRedis(":6379"))
	//for i := 0; i < 10; i++ {
	//	var u model.User
	//	err := shadow.Take(&u, "ciel", func(v interface{}) error {
	//		var temp model.User
	//		NewDefaultDB("root:123456@tcp(localhost:3306)/ciel202175?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai", true).
	//			Model(&temp).Where("id = 2").First(&temp)
	//		*v.(*model.User) = temp
	//		return nil
	//	})
	//	if err != nil {
	//		t.Fatal(err.Error())
	//	}
	//	fmt.Println(u)
	//}
	shadow.Del("ciel")
}
