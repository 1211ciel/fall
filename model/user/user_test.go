package model

import (
	"fmt"
	"github.com/1211ciel/fall/utils/dbutil"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"testing"
)

func TestDefaultUserModel_Register(t *testing.T) {
	err := NewUserModel(getDBR()).Register("ciel", "111")
	if err != nil {
		t.Fatal(err)
	}
}
func TestUserFindByUname(t *testing.T) {
	m := NewUserModel(getDBR())
	err := m.DelUserById(1)
	if err != nil {
		t.Fatal(err.Error())
	}
	//data.Nickname = "22"
	//err = m.UpdateUser(data)
	//if err != nil {
	//	t.Fatal(err.Error())
	//}

}
func TestNewLoginLogModel(t *testing.T) {
	m := NewLoginLogModel(getDBR())
	log := LoginLog{
		Uname: "ddd",
		Uid:   1,
		Ip:    "ddd",
	}
	err := m.CreateLoginLog(&log)
	if err != nil {
		t.Fatal(err)
	}
	data, err := m.FindLoginLogById(log.ID)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("%+v", data)
}

// initData
func TestInitUser(t *testing.T) {
	db, _ := getDBR()
	err := db.AutoMigrate(&User{}, &LoginLog{})
	if err != nil {
		t.Fatal(err.Error())
	}
}
func TestNewUser(t *testing.T) {
	err := NewUserModel(getDBR()).CreateUser(&User{
		Uname:   "ciel",
		Pwd:     "",
		SysCode: "",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}
func TestFindById(t *testing.T) {
	u, err := NewUserModel(getDBR()).FindUserById(1)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(u)
}
func TestUpdate(t *testing.T) {
	m := NewUserModel(getDBR())
	u, _ := m.FindUserById(1)
	fmt.Println(u)
	u.Pwd = "222"
	u.ID = 1
	err := m.UpdateUser(u)
	if err != nil {
		return
	}
	u, err = m.FindUserById(1)
	if err != nil {
		return
	}
	fmt.Println(u)
}
func TestDefaultUserModel_DeleteUserById(t *testing.T) {
	err := NewUserModel(getDBR()).DelUserById(1)
	if err != nil {
		return
	}
}
func getDBR() (*gorm.DB, *redis.Redis) {
	r := redis.NewRedis("localhost:6379", redis.NodeType, "")
	d := dbutil.GetDefaultMysql("root:123456@tcp(localhost:3306)/ciel?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai", true)
	return d, r
}
