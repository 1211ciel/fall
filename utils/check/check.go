package check

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	ErrEmptyPhoneNumber = errors.New("电话号码不能为空")
	ErrEmptyUsername    = errors.New("用户名不能为空")
	ErrEmptyPassword    = errors.New("密码不能为空")
	ErrEmptyEmail       = errors.New("邮箱不能为空")
	ErrFormatPhone      = errors.New("手机号码格式错误")
	ErrFormatEmail      = errors.New("邮箱格式不正确")
	ErrFormatPwd6       = errors.New("密码由6位数字组成")
)
var (
	errLevelC = errors.New("对特殊字符、大写字母、小写字母和数字至少存在1种 eg: csdn")
	errLevelB = errors.New("对特殊字符、大写字母、小写字母和数字至少存在2种 eg: csdn2020")
	errLevelA = errors.New("对特殊字符、大写字母、小写字母和数字至少存在3种 eg: csdn#2020")
	errLevelS = errors.New("密码中必须存在特殊字符、大小写字母和数字 eg:Csdn#2020")
	errLevel  = errors.New("不支持的Level等级")
)

// CheckUnameFormat 用户名由4-12位字母和数字组成
func CheckUnameFormat(u string) error {
	if u == "" {
		return ErrEmptyUsername
	}
	if res, _ := regexp.MatchString("^[a-zA-Z][a-zA-Z0-9_]{4,12}$", u); !res {
		return errors.New("用户名由4-12位字母和数字组成")
	}
	return nil
}

// CheckPhoneFormat 验证电话号码
func CheckPhoneFormat(p string) error {
	if p == "" {
		return ErrEmptyPhoneNumber
	}
	if res, _ := regexp.MatchString("^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|191|199|(147))\\d{8}$", p); !res {
		return ErrFormatPhone
	}
	return nil
}

func CheckPayPwdFormat(p string) error {
	if p == "" {
		return ErrEmptyPassword
	}
	if res, _ := regexp.MatchString("^\\d{6}$", p); !res {
		return ErrFormatPwd6
	}
	return nil
}

func CheckEmailFormat(p string) (err error) {
	if p == "" {
		return ErrEmptyEmail
	}
	if res, _ := regexp.MatchString("^([a-z0-9A-Z]+[-|_|\\.]?)+[a-z0-9A-Z]@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)?\\.)+[a-zA-Z]{2,}$", p); !res {
		return ErrFormatEmail
	}
	return
}

/*
密码强度	说明	示例
S	密码中必须存在特殊字符、大小写字母和数字	Csdn#2020
A	对特殊字符、大写字母、小写字母和数字至少存在3种 csdn#2020
B	对特殊字符、大写字母、小写字母和数字至少存在2种	csdn2020
C	对特殊字符、大写字母、小写字母和数字至少存在1种	csdn
D	不存在特殊字符、大小写字母和数字。	/、\
*/

type PwdLevel int

const (
	LevelD PwdLevel = iota // 不存在特殊字符、大小写字母和数字。	/、\
	LevelC                 // 对特殊字符、大写字母、小写字母和数字至少存在1种	csdn
	LevelB                 // 对特殊字符、大写字母、小写字母和数字至少存在2种	csdn2020
	LevelA                 // 对特殊字符、大写字母、小写字母和数字至少存在3种 csdn#2020
	LevelS                 // 密码中必须存在特殊字符、大小写字母和数字	Csdn#2020
)

/*
 *  minLength: 指定密码的最小长度
 *  maxLength：指定密码的最大长度
 *  minLevel：指定密码最低要求的强度等级
 *  pwd：明文密码
 */

// CheckPwd  6, 18 LevelB

func CheckPwd(minLength, maxLength int, minLevel PwdLevel, pwd string) (err error) {
	if pwd == "" {
		return ErrEmptyPassword
	}
	if len(pwd) < minLength {
		return fmt.Errorf("密码长度最少%d位", minLength)
	}
	if len(pwd) > maxLength {
		return fmt.Errorf("密码长度最多%d位", maxLength)
	}

	var level = LevelD
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[~!@#$%^&*?_-]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, pwd)
		if match {
			level++
		}
	}
	if level < minLevel {
		switch minLevel {
		case LevelC:
			err = errLevelC
		case LevelB:
			err = errLevelB
		case LevelA:
			err = errLevelA
		case LevelS:
			err = errLevelS
		default:
			err = errLevel
		}
	}
	return
}

func CheckDefaultPwd(pwd string) error {
	err := CheckPwd(6, 18, LevelB, pwd)
	return err
}
