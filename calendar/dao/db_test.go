package dao

import (
	"fmt"
	model "github.com/1211ciel/fall/model/user"
	"gorm.io/gorm"
	"testing"
)

func db() *gorm.DB {
	return NewDefaultDB("root:123456@tcp(localhost:3306)/ciel_21_7_1?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai", true)
}
func TestDBC(t *testing.T) {
	db := db()
	db.AutoMigrate(model.User{})
	db.Create(&model.User{
		Pid:   1,
		Uname: "ciel",
		Icon:  "1211.png",
		Pwd:   "123",
		Phone: "173622713362",
	})
}

func TestU(t *testing.T) {
	u := model.User{}
	db().Model(&u).Where("id = 1").Updates(map[string]interface{}{
		"pwd": "123456",
	})
	err := db().Model(&u).Where("id = 1").First(&u).Error
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("%+v", u)
}
func TestD(t *testing.T) {
	db().Where("id = 1").Unscoped().Delete(&model.User{})
}
func db2021_7_2() *gorm.DB {
	return NewDefaultDB("root:123456@tcp(localhost:3306)/ciel_21_7_2?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai", false)
}
func TestDBC202172(t *testing.T) {
	db := db2021_7_2()
	//db.AutoMigrate(model.User{})
	db.Create(&model.User{
		Pid:   1,
		Uname: "ciel",
		Icon:  "1211.png",
		Pwd:   "123",
		Phone: "1232832123",
	})
	var u model.User
	if err := db.Model(&u).Where("uname = ?", "ciel").First(&u).Error; err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(u)
}
func TestBCU202172(t *testing.T) {
	db := db2021_7_2()
	db.Model(&model.User{}).Where("uname = 'ciel'").Updates(map[string]interface{}{
		"icon": "1211sciel.png",
	})
	var u model.User
	if err := db.Model(&u).Where("id = 1").First(&u).Error; err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(u)
}
func TestDBD202172(t *testing.T) {
	//db2021_7_2().Unscoped().Delete(&model.User{}, "id = 1")
	var num int64
	db2021_7_2().Model(&model.User{}).Where("id = 1").Count(&num)
	fmt.Println(num)
}
func db2021_7_3() *gorm.DB {
	return NewDefaultDB("root:123456@tcp(localhost:3306)/ciel_21_7_3?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai", true)
}
func TestDB20210703(t *testing.T) {
	db := db2021_7_3()
	//db.AutoMigrate(model.User{})
	// c
	//db.Create(&model.User{
	//	Pid:   0,
	//	Uname: "ciel",
	//	Icon:  "test.png",
	//	Pwd:   "123",
	//	Phone: "13223335123",
	//})
	// u
	//if err := db.Model(&model.User{}).Where("id = 1").Update("icon", "test2.png").Error; err != nil {
	//	t.Fatal(err.Error())
	//}
	//// r
	//var u model.User
	//if err := db.Model(&u).Where("id = 1").First(&u).Error; err != nil {
	//	t.Fatal(err.Error())
	//}
	db.Unscoped().Delete(model.User{}, "id = 2")
}
func TestDB202174(t *testing.T) {
	db := NewDefaultDB("root:123456@tcp(localhost:3306)/ciel202174?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai", false)
	//db.AutoMigrate(model.User{})
	// create
	//affected := db.Create(&model.User{
	//	Pid: 0, Uname: "ciel", Icon: "test.png", Pwd: "123123", Phone: "1412131231",
	//}).RowsAffected
	//fmt.Println(affected)
	// u
	//db.Model(&model.User{}).Where("id = 1").Update("icon", "test1.png")
	//var u model.User
	//if err := db.Model(&u).Where("id = 1").First(&u).Error; err != nil {
	//	t.Fatal(err.Error())
	//}
	//fmt.Println(&u)
	affected := db.Model(model.User{}).Unscoped().Delete(model.User{}, "id = 1").RowsAffected
	fmt.Println(affected)
}
func TestDB202175(t *testing.T) {
	db := NewDefaultDB("root:123456@tcp(localhost:3306)/ciel202175?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai", false)
	// c
	db.Model(model.User{}).Create(&model.User{
		Pid: 0, Uname: "ciel", Icon: "text.png", Phone: "1522233111", Pwd: "123123",
	})
	// u
	//db.Model(model.User{}).Where("id = 1").Update("icon", "test2.png")
	// r
	//var u model.User
	//err := db.Model(&u).Where("id = 1").First(&u).Error
	//if err != nil {
	//	t.Fatal(err.Error())
	//}
	//fmt.Println(u)
	// d
	//db.Model(&model.User{}).Unscoped().Delete(&model.User{}, "id = 1")
}
