package e

const (
	// SUCCESS .
	SUCCESS = 200

	// 1 开头 参数错误（参数无效、参数类型错误等）

	ErrParamIsInvalid     = 10001
	ErrParamTypeBindError = 10002
	ErrInvalidAuth        = 10003
	ErrTokenIsNull        = 10004
	ErrTokenExpired       = 10005

	// 2 开头 用户错误（登陆失败、密码错误等）

	// 3 开头 业务错误 （创建失败，更新失败等）

	ErrActionFail = 30000

	// 4 开头 （数据未找到，数据有误等）

	ErrDataNotFound = 40000

	// 5 开头 系统内部错误 （接口禁止访问，接口负载过高，系统内部异常等）

	ErrError           = 50000
	ErrCodeIsNotDefine = 50001
)
