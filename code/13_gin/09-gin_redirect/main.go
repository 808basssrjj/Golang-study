package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//1. http重定向
	router.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	// 2.路由重定向
	router.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2" //修改请求的url
		router.HandleContext(c)       //继续后续处理
	})
	router.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "重定向成功",
		})
	})

	router.Run(":9090")
}
