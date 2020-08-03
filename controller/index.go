package controller

import (
	"gin-scaffold/pkg/comm"
	"gin-scaffold/pkg/e"

	"github.com/gin-gonic/gin"
)

// IndexController .
type IndexController struct {
}

// Get .
// @Summary index page
// @Description 测试接口
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} comm.Resp{data=model.Account}
// @Router /index [get]
func (ctl *IndexController) Get(c *gin.Context) {
	comm.ReturnJSON(c, e.SUCCESS)
}
