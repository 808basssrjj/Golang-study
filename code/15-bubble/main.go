package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	// 注册路由
	r := routers.SetupRouter()
	r.Run(":7070")

}
