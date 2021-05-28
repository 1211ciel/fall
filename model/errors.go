package model

import "errors"

// User
var (
	ErrCode                   = errors.New("验证码不能为空")
	ErrPhone                  = errors.New("手机号不能为空")
	ErrCodeExpire             = errors.New("验证码已过期")
	ErrCodeNotMatch           = errors.New("验证码错误")
	ErrSendCode               = errors.New("请一分钟后重试")
	ErrUnameOrPhoneExisted    = errors.New("用户名或电话号码已被占用")
	ErrUnameExisted           = errors.New("用户名已被占用")
	ErrInvitationCodeNotExist = errors.New("邀请码不存在")
	ErrUserNotMatch           = errors.New("用户名不存在,或密码错误")
	ErrUserNotFond            = errors.New("用户不存在,或已被禁用")
	ErrUserOldPwd             = errors.New("旧密码不匹配")
	ErrUserPwdSame            = errors.New("两次密码相同")
	ErrUserAuthAlready        = errors.New("用户已实名认证，请不要重复操作")
	ErrOnlyHaveOneBankCard    = errors.New("只能绑定一张银行卡")
	ErrBankCardAlreadyBound   = errors.New("银行卡已被绑定")
	ErrOldPwdNotMatch         = errors.New("旧密码不匹配")
)
