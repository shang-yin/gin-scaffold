package controller

type publicLogin struct {
	Username string `json:"username" binding:"required,checkMobile" comment:"用户名"` //  用户名
	Password string `json:"password" binding:"required" comment:"密码"`              // 密码
}
