package e

var msgFlags = map[int]string{
	SUCCESS: "成功",

	ErrParamIsInvalid:     "无效参数",
	ErrParamTypeBindError: "参数类型错误",
	ErrInvalidAuth:        "token验证失败",
	ErrTokenIsNull:        "token不能为空",
	ErrTokenExpired:       "token已过期",

	ErrActionFail:   "操作失败",
	ErrDataNotFound: "数据未找到",

	ErrError:           "系统内部错误，请联系管理员",
	ErrCodeIsNotDefine: "状态码未定义",
}

// GetMsg .
func GetMsg(code int) string {
	if msg, ok := msgFlags[code]; ok {
		return msg
	}
	return msgFlags[ErrCodeIsNotDefine]
}
