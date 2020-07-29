package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.New()

	e.GET("/demo", func(c *gin.Context) {
		c.String(http.StatusOK, "demo")
	})
	e.GET("/index", func(c *gin.Context) {
		c.String(http.StatusOK, "index")
	})
	if err := e.Run(":8081"); err != nil {
		fmt.Println(err)
	}
}
