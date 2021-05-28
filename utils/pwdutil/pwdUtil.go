package pwdutil

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotMatch = errors.New("用户名不存在,或密码错误")
)

func GenPwd(pwd string) string {
	password, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(password)
}

// ComparePwd pwd 数据库里面的密码 rePwd 用户输入的密码
func ComparePwd(localPwd, inputPwd string) error {
	if e := bcrypt.CompareHashAndPassword([]byte(localPwd), []byte(inputPwd)); e != nil {
		return ErrUserNotMatch
	}
	return nil
}
