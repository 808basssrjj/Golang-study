package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 1.普通路由
	// 为没有配置处理函数的路由添加处理程序
	// r.GET("/index", func(c *gin.Context) {...})
	// r.GET("/login", func(c *gin.Context) {...})
	// r.POST("/login", func(c *gin.Context) {...})
	//匹配所有请求方法
	r.Any("/test", func(c *gin.Context) {
		var msg string
		switch c.Request.Method {
		case http.MethodGet:
			msg = "GET"
		case http.MethodPost:
			msg = "POST"
		}
		c.JSON(http.StatusOK, gin.H{
			"method": msg,
		})
	})
	// 为没有配置处理函数的路由添加处理程序
	r.LoadHTMLFiles("views/404.html")
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "404 not found",
		})
	})

	//2. 路由组
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "userIndex"})
		})
		userGroup.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "userLogin"})
		})
	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "shopIndex"})
		})
		shopGroup.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "shopLogin"})
		})
	}

	r.Run(":8081")
}
