package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//StatCost 定义一个中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "aa") //通过 c.Set在请求的上下文中设置值，后续函数可以拿到

		c.Next() // 调用该请求的剩余处理程序
		// c.Abort()// 不调用该请求的剩余处理程序
		cost := time.Since(start)
		fmt.Println(cost)
	}

}

func main() {
	r := gin.New()

	// 注册一个全局中间件
	r.Use(StatCost())

	// 为某个路由单独注册中间件
	// r.GET("/index", StatCost, func(c *gin.Context) {
	r.GET("/index", func(c *gin.Context) {
		// name ,ok := c.Get("name").(string) //从上下文取值
		name := c.MustGet("name").(string)
		c.JSON(http.StatusOK, gin.H{"msg": name})
	})

	// 路由组注册中间件1
	userGroup := r.Group("/user", StatCost())
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "userIndex"})
		})
		userGroup.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "userLogin"})
		})
	}
	// 路由组注册中间件2
	user2Group := r.Group("/user2")
	user2Group.Use(StatCost())
	{
		user2Group.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "userIndex"})
		})
	}

	r.Run(":8081")

	// gin.Default()默认使用了Logger和Recovery中间件，其中：
	// Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。
	// Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
	// 如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由。

	// 当在中间件或handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）。
}
