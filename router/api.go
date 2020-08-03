package router

import (
	"gin-scaffold/controller"

	"github.com/gin-gonic/gin"
)

func apiRouter(r *gin.Engine) {
	// v1 版本API
	v1 := r.Group("/api/v1")
	{
		publicController := controller.PublicController{}
		publicGroup := v1.Group("/public")
		{
			publicGroup.POST("/login", publicController.Login)
		}

		indexController := controller.IndexController{}
		indexGroup := v1.Group("/index")
		{
			indexGroup.GET("", indexController.Get)
		}

		accountController := controller.AccountController{}
		accountGroup := v1.Group("/account")
		{
			accountGroup.POST("", accountController.Create)
		}
	}
}
