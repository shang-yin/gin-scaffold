package controller

type accountCreateRequest struct {
	Username string `json:"username" binding:"required" comment:"用户名"`           //  用户名
	Password string `json:"password" binding:"required" comment:"密码"`            // 密码
	Mobile   string `json:"mobile" binding:"required,checkMobile" comment:"手机号"` //  手机号
}
