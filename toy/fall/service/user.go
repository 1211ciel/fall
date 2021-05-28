package service

import (
	"github.com/1211ciel/fall/common/consts"
	model "github.com/1211ciel/fall/model/user"
	"github.com/1211ciel/fall/toy/fall/svc"
	"github.com/1211ciel/fall/utils/pwdutil"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

func Register(c *cli.Context) error {
	uname := c.String("uname")
	pwd := c.String("pwd")
	if err := checkUnamePwd(uname, pwd); err != nil {
		return nil
	}
	if err := svc.UserModel.Register(uname, pwd); err != nil {
		color.Red(err.Error())
		return nil
	}
	color.Green("创建成功,去登录吧")
	return nil
}
func Login(c *cli.Context) error {
	uname := c.String("uname")
	pwd := c.String("pwd")
	if err := checkUnamePwd(uname, pwd); err != nil {
		return nil
	}
	if err := svc.DB.Transaction(func(tx *gorm.DB) error {
		var u model.User
		if err := tx.Model(&u).Where("uname = ?", uname).First(&u).Error; err != nil {
			color.Red("用户不存在")
			return nil
		}
		if err := pwdutil.ComparePwd(u.Pwd, pwd); err != nil {
			color.Red("密码错误")
			return nil
		}
		if u.LoginStatus == 1 {
			color.Red("该用户已经,请先退出")
			return nil
		}
		tx.Model(&u).Where("id = ?", u.ID).Update("login_status", "1")
		color.Green("登录成功")
		return nil
	}); err != nil {
		color.Red(err.Error())
		return nil
	}
	return nil
}

func Logout(c *cli.Context) error {
	var u model.User
	db := svc.DB.Model(&u).Where("login_status = 1")
	err := db.First(&u).Error
	if err != nil {
		color.Red("请登录后操纵")
		return nil
	}
	if u.LoginStatus == 0 {
		color.Yellow("用户%v已退出请不要重复操作", u.Uname)
		return nil
	}
	db.Updates(map[string]interface{}{"login_status": 0})
	color.Green("%v已退出", u.Uname)
	return nil
}
func checkUnamePwd(uname, pwd string) error {
	if uname == "" {
		color.Red("用户名不能为空")
		return consts.ErrParam
	}
	if pwd == "" {
		color.Red("密码不能为空")
		return consts.ErrParam
	}
	return nil
}

func currentUser() (*model.User, error) {
	var u model.User
	err := svc.DB.Model(&u).Where("login_status = 1").First(&u).Error
	if err != nil {
		color.Red("请先登录")
		return nil, err
	}
	return &u, nil
}
