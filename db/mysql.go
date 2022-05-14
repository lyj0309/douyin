package db

import "gorm.io/gorm"
import "gorm.io/driver/mysql"

var Mysql *gorm.DB

func init() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	Mysql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
