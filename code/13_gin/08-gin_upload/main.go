package main

import (
	"log"
	"net/http"
	"path"

	// "path"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// f1 对应html中input的name
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
			return
		}
		log.Println(file.Filename)

		// 保存到本地 SaveUploadedFile
		dst := path.Join("./", file.Filename)
		_ = c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// 多文件上传
	r.POST("/uploads", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		// files 对应html中input的name
		files := form.File["files"]
		for index, file := range files {
			log.Println(file.Filename, index)
			dst := path.Join("./", file.Filename)
			_ = c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	//加载html
	r.LoadHTMLFiles("./upload.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})

	r.Run(":9090")
}
