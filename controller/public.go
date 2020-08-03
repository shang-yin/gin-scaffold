package controller

import (
	"gin-scaffold/pkg/comm"
	"gin-scaffold/pkg/e"
	"gin-scaffold/pkg/request"

	"github.com/gin-gonic/gin"
)

type PublicController struct {
}

// Login .
// @Summary login page
// @Description login api
// @Accept  json
// @Produce  json
// @Param body body publicLogin true "登陆参数"
// @Success 200 {object} comm.Resp{data=model.Account}
// @Router /login [get]
func (ctl *PublicController) Login(c *gin.Context) {
	var (
		err error
		req publicLogin
	)
	if err = request.ParseRequest(c, &req); err != nil {
		return
	}
	comm.ReturnJSON(c, e.SUCCESS, req)
}
