package controller

import (
	"fmt"
	"gin-scaffold/model"
	"gin-scaffold/pkg/comm"
	"gin-scaffold/pkg/e"
	"gin-scaffold/pkg/encrypt"
	"gin-scaffold/pkg/request"

	"github.com/gin-gonic/gin"
)

// AccountController .
type AccountController struct {
}

// Create .
func (ctl *AccountController) Create(c *gin.Context) {
	var (
		err     error
		req     accountCreateRequest
		account = &model.Account{}
	)
	if err = request.ParseRequest(c, &req); err != nil {
		return
	}
	account.Password, err = encrypt.GeneratePassword(req.Password)
	if err != nil {
		comm.ReturnJSON(c, e.ErrError)
		return
	}
	fmt.Println(req)
	// if err := account.Insert(); err != nil {
	// 	comm.ReturnJSON(c, e.ErrCodeIsNotDefine)
	// }
}
