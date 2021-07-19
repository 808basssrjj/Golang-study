package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Userinfo 数据表
type Userinfo struct {
	ID     int
	Name   string
	Gender string
	Hobby  string
}

func main() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 自动迁移
	db.AutoMigrate(&Userinfo{})

	// 创建
	u1 := Userinfo{1, "路飞", "男", "吃肉"}
	u2 := Userinfo{2, "索隆", "男", "喝酒"}
	db.Create(u1)
	db.Create(&u2)

	// 查询
	var u Userinfo
	db.First(&u)
	fmt.Printf("%#v\n", u)

	// 条件查询
	var uu Userinfo
	db.Find(&uu, "hobby=?", "吃肉")
	fmt.Printf("%#v", uu)

	// 更新
	db.Model(&u).Update("hobby", "双色球")

	// 删除
	db.Delete(&u)
}
