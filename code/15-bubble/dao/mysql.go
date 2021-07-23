package dao

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

// 连接数据库
func InitMysql() (err error) {
	dsn := "root:123456@(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return DB.DB().Ping()
}

// 关闭数据库
func Close() {
	DB.Close()
}
