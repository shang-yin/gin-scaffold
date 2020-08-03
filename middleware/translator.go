package middleware

import (
	"fmt"
	"gin-scaffold/pkg/logging"
	"gin-scaffold/pkg/translation"

	"github.com/gin-gonic/gin"
)

func Trans() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := translation.InitTrans("zh"); err != nil {
			logging.Error(fmt.Sprintf("translation.InitTrans error = %v", err))
		}
	}
}
