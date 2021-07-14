package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// GET:请求方式 /hello:请求的路径
	// 当客户端以GET方法请求/hello路径时,会执行后面的匿名函数
	// get 用来获取
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"method": "get",
		})
	})

	// post 用来新建
	r.POST("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "post",
		})
	})

	// put 用来更新
	r.PUT("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "put",
		})
	})

	// delete 用来更新
	r.DELETE("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "delete",
		})
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	_ = r.Run(":9090")
}
