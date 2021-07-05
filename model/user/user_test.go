package model

import (
	"fmt"
	"github.com/1211ciel/word-of-wind/calendar/dao"
	"testing"
)

func TestNewUserModel(t *testing.T) {
	db := dao.NewDefaultDB("root:123456@tcp(localhost:3306)/ciel202175?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai", true)
	redis := dao.NewDefaultRedis(":6379")
	m := NewUserModel(db, redis)
	u, err := m.FindUserById(5)
	if err != nil {
		t.Fatal(err.Error())
	}
	u.Icon = "888.png"
	m.UpdateUser(u)
	fmt.Println(u)
	//err := m.CreateUser(&User{
	//	Pid:   0,
	//	Uname: "ciel2",
	//	Icon:  "test.png",
	//	Pwd:   "123456",
	//	Phone: "123123123",
	//})
	//if err != nil {
	//	t.Fatal(err)
	//}
}
