package main

import "github.com/gin-gonic/gin"

// 1.初始化go.mod
// go mod init 项目名
// 2. 安装依赖
// go get -u 地址
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hhhhh")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
