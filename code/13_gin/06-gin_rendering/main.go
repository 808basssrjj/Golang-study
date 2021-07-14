package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 2. 自定义函数 定义一个不转义的函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	// 3.静态文件处理  使用时名字 , 根路径
	r.Static("/xxx", "./statics")

	//模板解析	 使用LoadHTMLGlob()或者LoadHTMLFiles()方法进行HTML模板解析
	// r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/posts/index", func(c *gin.Context) {
		// 模板渲染
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
			"a":     "<a href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		// 模板渲染
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "users/index",
			"a":     "<a href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})

	// 4.json渲染
	var msg struct {
		Name    string `json:"user"`
		Message string
		Age     int
	}
	msg.Name = "迪迦"
	msg.Message = "hello"
	msg.Age = 10000
	r.GET("/search", func(c *gin.Context) {
		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "hello",
		// })
		c.JSON(http.StatusOK, msg) //利用结构体
	})
	_ = r.Run(":9090")
}
