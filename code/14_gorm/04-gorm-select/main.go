package main

import (
	"database/sql"
	"fmt"

	// "fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID     int64
	Name   *string        `gorm:"default:'娜美'"`
	Gender sql.NullString `gorm:"default:'男'"`
	Age    int64
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})

	var user User
	// 根据主键查询第一条记录
	// db.First(&user)
	//// SELECT * FROM users ORDER BY id LIMIT 1;

	// 随机获取一条记录
	// db.Take(&user)
	//// SELECT * FROM users LIMIT 1;

	// 根据主键查询最后一条记录
	// db.Last(&user)
	//// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	// 查询所有的记录}
	// var users []User
	// db.Debug().Find(&users)
	//// SELECT * FROM users;
	// fmt.Printf("%v\n", users)

	// 查询指定的某条记录(仅当主键为整型时可用)
	db.First(&user, 10)
	//// SELECT * FROM users WHERE id = 10;
	fmt.Printf("%v\n", user)

	// Get first matched record
	// db.Where("name = ?", "jinzhu").First(&user)
	db.Where("age = ?", 18).First(&user)
	//// SELECT * FROM users WHERE name = 'jinzhu' limit 1;

	// Get all matched records
	// db.Where("name = ?", "jinzhu").Find(&users)
	//// SELECT * FROM users WHERE name = 'jinzhu';

	// <>
	// db.Where("name <> ?", "jinzhu").Find(&users)
	//// SELECT * FROM users WHERE name <> 'jinzhu';

	// IN
	// db.Where("name IN (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	//// SELECT * FROM users WHERE name in ('jinzhu','jinzhu 2');

	// LIKE
	// db.Where("name LIKE ?", "%jin%").Find(&users)
	//// SELECT * FROM users WHERE name LIKE '%jin%';

	// AND
	// db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	//// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

	// Time
	// db.Where("updated_at > ?", lastWeek).Find(&users)
	//// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

	// BETWEEN
	// db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
	//// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
	fmt.Printf("%v\n", user)

}
