package consts

import "errors"

var (
	ErrService    = errors.New("服务器繁忙,请稍后再试")
	ErrIllegal    = errors.New("非法操作")
	ErrUnDontKnow = errors.New("出现未知错误")
)

// data
var (
	ErrData             = errors.New("数据库操作异常")
	ErrDataNotFound     = errors.New("数据不存在")
	ErrDataAlreadyExist = errors.New("数据已存在")
	ErrDataCommitTx     = errors.New("事务提交失败")
	ErrDataBeginTx      = errors.New("事务开启失败")
	ErrDataUpdate       = errors.New("数据更新失败")
	ErrDelFailed        = errors.New("数据不存在或已被删除")
	ErrParse            = errors.New("数据解析失败")
)

var (
	ErrRedisUnmarshal     = errors.New("解析缓存数据出错")
	ErrRedisMarshal       = errors.New("json编码失败")
	ErrRedisPut           = errors.New("缓存写入失败")
	ErrRedisDel           = errors.New("缓存清除失败")
	ErrRedisLock          = errors.New("缓存加锁失败")
	ErrRedisUnLock        = errors.New("缓存解锁失败")
	ErrRedisRelease       = errors.New("缓存释放锁失败")
	ErrRedisOperationFast = errors.New("操作太快请稍后重试")
	ErrRedisDataEmpty     = errors.New("缓存数据为空")
)

// sys
var (
	ErrParam           = errors.New("参数解析失败")
	ErrType            = errors.New("类型错误")
	ErrRepeatOperation = errors.New("重复的操作")
)

// User
var (
	ErrCode                   = errors.New("验证码不能为空")
	ErrPhone                  = errors.New("手机号不能为空")
	ErrCodeExpire             = errors.New("验证码已过期")
	ErrCodeNotMatch           = errors.New("验证码错误")
	ErrSendCode               = errors.New("请一分钟后重试")
	ErrUsernameOrPhoneExisted = errors.New("用户名或电话号码已被占用")
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

// file
var (
	ErrGetFile    = errors.New("获取文件失败")
	ErrUploadFile = errors.New("上传文件失败")
)

//  im

var (
	ErrImUserNotOnline = errors.New("该用户暂时不在线,消息发送失败")
)
