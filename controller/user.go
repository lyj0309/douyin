package controller

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lyj0309/douyin/db"
	"github.com/lyj0309/douyin/utils"
)

// code 0：成功， 1：失败， 2：用户或密码为空

func Register(c *gin.Context) {

	var user db.User
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")

	if user.Name == "" || user.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 2,
			"msg":  "密码或用户名不能为空",
		})
		return
	}

	//需要判断该用户名是否被占用
	res := db.Mysql.Where("name = ?", user.Name).Find(&user)
	if res.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "该用户名已被占用",
		})
		return
	}

	//将用户信息插入数据库
	db.Mysql.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
	return

}

func Login(c *gin.Context) {

	username := c.PostForm("name")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 2,
			"msg":  "用户账号或密码为空",
		})
		return
	}

	//先查看是否存在该用户
	var user db.User
	//db.Mysql.AutoMigrate(&user)
	res := db.Mysql.Find(&user, "name = ? AND password = ?", username, password)

	// select * from user where
	if res.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "用户名或密码错误",
		})
		return
	}
	token, err := utils.GenToken(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "生成token失败",
		})
		return
	}

	//将jwt存储在redis中
	db.Rdb.Set(c, "token:"+token, username, time.Hour*24*15)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登录成功",
		"data": gin.H{"token": token},
	})
}

func Logout(c *gin.Context) {
	jwt := c.Request.Header.Get("Authorization")
	token := strings.SplitN(jwt, " ", 2)[1]
	if token == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "token为空",
		})
	}
	//从redis中删除token
	db.Rdb.Del(c, "token:"+token)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "成功退出登录！",
	})
}

func GetUserInfo(c *gin.Context) {
	username, _ := c.Get("name")
	var user db.User

	result := db.Mysql.Where(" name = ?", username).Find(&user)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 2,
			"msg":  "该用户不存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": user,
	})
	return
}
