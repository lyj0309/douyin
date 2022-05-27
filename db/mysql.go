package db

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Mysql *gorm.DB

func init() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local&interpolateParams=true"

	logrus.Info("初始化数据库···")

	Mysql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = Mysql.AutoMigrate(User{}, Comment{}, Like{}, Video{})
	if err != nil {
		logrus.Errorln("表生成出错", err)
		return
	}
}
