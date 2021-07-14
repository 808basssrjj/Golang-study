package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	// 1.获取querystring参数   /user/search?username=小王子&address=沙河
	r.GET("/query", func(c *gin.Context) {
		username := c.DefaultQuery("username", "迪迦") //设置默认值
		// username:= c.Query("username")
		addrss := c.Query("address")
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"usernmae": username,
			"address":  addrss,
		})
	})

	// 2.获取form参数
	r.POST("/postform", func(c *gin.Context) {
		username := c.PostForm("username")
		address := c.DefaultPostForm("address", "火星") //默认值
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"address":  address,
		})
	})

	// 3.获取json参数
	r.POST("/json", func(c *gin.Context) {
		b, _ := c.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(b, &m)
		c.JSON(http.StatusOK, m)
	})

	// 4.获取path参数  /user/search/小王子/沙河
	r.GET("/param/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"address":  address,
		})
	})

	// 5.参数绑定
	// 为了能够更方便的获取请求相关参数，提高开发效率，
	// 我们可以基于请求的Content-Type识别请求数据类型并利用反射机制自动提取请求中QueryString、form表单、JSON、XML等参数到结构体中
	// ShouldBind会按照下面的顺序解析请求中的数据完成绑定：
	// 如果是 GET 请求，只使用 Form 绑定引擎（query）。
	// 如果是 POST 请求，首先检查 content-type 是否为 JSON 或 XML，然后再使用 Form（form-data）

	// 5.1// 绑定JSON的示例 ({"user": "q1mi", "password": "123456"})
	r.POST("/loginJSON", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}
	})
	// 5.2绑定form表单示例 (user=q1mi&password=123456)
	r.POST("/loginForm", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}
	})
	// 5.3绑定QueryString示例 (/loginQuery?user=q1mi&password=123456)
	r.GET("/loginQuery", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}
	})

	r.Run(":7070")
}
