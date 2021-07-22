package main

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 定义模型
type User struct {
	ID     int64
	Name   *string        `gorm:"default:'娜美'"` // 默认值    使用指针来存零值
	Gender sql.NullString `gorm:"default:'男'"`  // 使用sql.NullString来存零值   sql.NullString实现了Scanner/Valuer接口
	Age    int64
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("mysql connect failed err:", err)
		return
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	// u1 := User{Age: 18}
	// fmt.Println(db.NewRecord(u1)) // 主键为空返回`true`
	// db.Create(&u1)                // 创建user
	// fmt.Println(db.NewRecord(u1)) // 创建`user`后返回`false`

	// 所有字段的零值, 比如0, "",false或者其它零值，都不会保存到数据库内，但会使用他们的默认值。
	// 如果你想避免这种情况，可以考虑使用指针或实现 Scanner/Valuer接口，
	u2 := User{Name: new(string), Gender: sql.NullString{String: "", Valid: true}, Age: 18}
	fmt.Println(db.NewRecord(u2))
	db.Debug().Create(&u2) //Debug()打印sql
	fmt.Println(db.NewRecord(u2))

}
