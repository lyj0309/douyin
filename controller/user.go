package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lyj0309/douyin/db"
	"github.com/lyj0309/douyin/utils"
	"net/http"
)

// code 0：成功， 1：失败， 2：用户或密码为空

func Register(c *gin.Context) {

	var user db.User
	user.Name = c.Query("username")
	user.Password = c.Query("password")

	if user.Name == "" || user.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 2,
			"status_msg":  "密码或用户名不能为空",
		})
		return
	}

	//需要判断该用户名是否被占用
	res := db.Mysql.Where("name = ? AND password = ?", user.Name, user.Password).Find(&user)
	if res.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "该用户已存在",
		})
		return
	}

	//将用户信息插入数据库
	db.Mysql.Save(&user)

	token, err := utils.GenToken(user.ID, user.Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "register fail",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "success",
		"user_id":     user.ID,
		"token":       token,
	})

}

func Login(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 2,
			"status_msg":  "用户账号或密码为空",
		})
		return
	}

	//先查看是否存在该用户
	var user db.User

	res := db.Mysql.Where(" name = ? AND password = ?", username, password).Find(&user)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "用户名或密码错误",
		})
		return
	}

	token, err := utils.GenToken(user.ID, user.Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "登录成功",
		"user_id":     user.ID,
		"token":       token,
	})
}

func GetUserInfo(c *gin.Context) {
	uid, _ := c.Get("user_id")
	
	var user db.User

	result := db.Mysql.Where(" id = ?", uid).Find(&user)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 2,
			"status_msg":  "该用户不存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "success",
		"user":        user,
	})

}
