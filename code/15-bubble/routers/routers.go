package routers

import (
	"bubble/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 加载路径
	r.LoadHTMLGlob("templates/*")
	// 静态资源
	r.Static("/static", "static")

	// 首页
	r.GET("/", controller.IndexHandler)
	//Api
	v1Group := r.Group("v1")
	{
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)

		// 查看所有
		v1Group.GET("/todo", controller.GetTodoList)

		// 修改某个
		v1Group.PUT("/todo/:id", controller.UpdateATodo)

		// 删除某个
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}

	return r
}
