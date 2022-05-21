package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lyj0309/douyin/db"
	"github.com/lyj0309/douyin/utils"
)

// code 0：成功， 1：失败， 2：用户或密码为空

func Register(c *gin.Context) {

	var user db.User
	user.Name = c.PostForm("username")
	user.Password = c.PostForm("password")
	if user.Name == "" || user.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 2,
			"msg":  "密码或用户名不能为空",
		})
		return
	}
	//将用户信息插入数据库，生成token
	db.Mysql.AutoMigrate(&user)

	//需要判断该用户名是否被占用
	res := db.Mysql.Where("username = ?", user.Name).Find(&user)
	if res.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "该用户名已被占用",
		})
		return
	}

	db.Mysql.Create(&user)
	token, err := utils.GenToken(user.Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "无效的token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": gin.H{"token": token},
	})
	return

}

func Login(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 2,
			"msg":  "用户账号或密码为空",
		})
		return
	}

	//先查看是否存在该用户
	var u db.User
	db.Mysql.AutoMigrate(&u)
	db.Mysql.Where("username = ? AND password = ?", username, password).Find(&u)

	// select * from user where
	if u.Name != username || u.Password != password {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "用户名或密码错误",
		})
		return
	}
	token, err := utils.GenToken(u.Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登录成功",
		"data": gin.H{"token": token},
	})
}

func GetUserInfo(c *gin.Context) {
	username, _ := c.Get("username")
	var user db.User

	result := db.Mysql.Where(" username = ?", username).Find(&user)

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
