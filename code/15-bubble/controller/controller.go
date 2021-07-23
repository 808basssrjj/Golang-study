package controller

import (
	"bubble/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
 url     --> controller  --> logic   -->    model
请求来了  -->  控制器      --> 业务逻辑  --> 模型层的增删改查
*/

// 首页
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// 添加
func CreateTodo(c *gin.Context) {
	// 1.从请求中获取数据
	var todo models.Todo
	c.BindJSON(&todo)

	// 2.加入数据库
	err := models.CreateATodo(&todo) //调用模型

	// 3.返回响应
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errOR": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// 查看全部
func GetTodoList(c *gin.Context) {
	todoList, err := models.GetAllTodo() //调用模型

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

// 修改某个
func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}

	todo, err := models.GetATode(id) //调用模型
	if err != nil {	//数据不存在
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.BindJSON(&todo)
	if err := models.UpdateATodo(todo); err != nil { //调用模型
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// 删除某个
func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}

	err := models.DeleteATodo(id) //调用模型
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "删除成功!"})
	}
}
