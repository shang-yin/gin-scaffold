package router

import (
	_ "gin-scaffold/docs"
	"gin-scaffold/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.Trans())
	// swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// api router
	apiRouter(r)

	// backend router
	backendRouter(r)

}
