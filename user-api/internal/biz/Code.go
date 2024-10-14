package biz

import "github.com/zeromicro/x/errors"

const (
	Ok = 200
)

//var (
//	DBError          = NewError(10000, "数据库错误")
//	UserIsRegistered = NewError(10100, "用户已注册")
//)

var (
	DBError          = errors.New(10000, "数据库错误")
	UserIsRegistered = errors.New(10100, "用户已注册")
)
