package consts

import "errors"

var (
	ErrService = errors.New("服务器繁忙,请稍后再试")
	ErrIllegal = errors.New("非法操作")
	ErrUnKnow  = errors.New("出现未知错误")
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

// sys
var (
	ErrParam           = errors.New("参数解析失败")
	ErrType            = errors.New("类型错误")
	ErrRepeatOperation = errors.New("重复的操作")
)
