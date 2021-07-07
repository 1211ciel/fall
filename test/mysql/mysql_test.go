package mysql

import (
	"fmt"
	model "github.com/1211ciel/fall/model/user"
	"github.com/1211ciel/word-of-wind/calendar/dao"
	"testing"
)

func Test77(t *testing.T) {
	db := dao.NewDefaultDB("root:123456@tcp(localhost:3306)/ciel77?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai", true)
	//err := db.AutoMigrate(model.User{})
	//if err != nil {
	//	t.Fatal(err.Error())
	//}
	// c
	//user := model.NewUser(1, "ciel", "test.png", "123", "123455")
	//err = db.Create(&user).Error
	//if err != nil {
	//	t.Fatal(err.Error())
	//}
	// u
	//var u model.User
	//db.Model(&u).Where("id = 2").Update("icon", "icon2.png")
	//err := db.Model(&u).Where("id = 2").First(&u).Error
	//if err != nil {
	//	t.Fatal(err.Error())
	//}
	//fmt.Println(&u)
	// d
	affected := db.Model(&model.User{}).Unscoped().Delete(model.User{}, "id = 2").RowsAffected
	fmt.Println(affected)
}
