package biz

var (
	DBError          = NewError(10000, "数据库错误")
	UserIsRegistered = NewError(10100, "用户已注册")
	NameOrPwdError   = NewError(10101, "用户名或密码错误")
	TokenError       = NewError(10102, "token错误")
	InvalidToken     = NewError(10103, "token无效")
	RedisError       = NewError(10104, "redis错误")
)

//var (
//	DBError          = errors.New(10000, "数据库错误")
//	UserIsRegistered = errors.New(10100, "用户已注册")
//	NameOrPwdError   = errors.New(10101, "用户名或密码错误")
//	TokenError       = errors.New(10102, "token错误")
//)
