package comm

import (
	"net/http"

	"gin-scaffold/pkg/e"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Code int         `json:"code"` // code 码
	Msg  string      `json:"msg"`  // 具体的错误消息
	Data interface{} `json:"data"` // 返回的数据
}

// ReturnJSON  返回json
func ReturnJSON(c *gin.Context, code int, data ...interface{}) {
	if len(data) == 0 {
		c.JSON(http.StatusOK, Resp{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: make(map[string]string),
		})
	} else if len(data) == 1 {
		c.JSON(http.StatusOK, Resp{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: data[0],
		})
	} else {
		c.JSON(http.StatusOK, Resp{
			Code: code,
			Msg:  e.GetMsg(code),
			Data: data,
		})
	}
}
