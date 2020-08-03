package main

import (
	"fmt"
	"gin-scaffold/migration"
	"gin-scaffold/router"

	//  这里要注意排序问题
	_ "gin-scaffold/pkg/conf"
	_ "gin-scaffold/pkg/database"
	_ "gin-scaffold/pkg/logging"
	_ "gin-scaffold/pkg/nosql"

	"github.com/gin-gonic/gin"
)

// @title gin-scaffold api docs
// @version 1.0
// @description 项目接口文档

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1

func init() {
	// 数据库自动创建
	migration.Migration()
}
func main() {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)

	router.InitRouter(r)
	if err := r.Run(":8081"); err != nil {
		panic(fmt.Sprintf("start server failed!! err = %v", err))
	}
}
