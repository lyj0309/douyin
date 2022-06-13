package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lyj0309/douyin/db"
	// "github.com/lyj0309/douyin/service"
)

type VideoListResponse struct {
	Response
	VideoList []db.Video `json:"video_list"`
}

// var usersLoginInfo = map[string]User{
// 	"zhangleidouyin": {
// 		Id:            1,
// 		Name:          "zhanglei",
// 		FollowCount:   10,
// 		FollowerCount: 5,
// 		IsFollow:      true,
// 	},
// }

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	// 判断 token 中的 user 是否存在
	username, _ := c.Get("username")
	var user db.User
	var video db.Video

	result := db.Mysql.Where(" name = ?", username).Find(&user)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 2,
			"msg":  "该用户不存在",
		})
		return
	}

	// 读取视频数据 data
	data, err := c.FormFile("data")

	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 加工保存视频数据
	filename := filepath.Base(data.Filename)

	finalName := fmt.Sprintf("%d_%s", user.ID, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 读取视频标题title
	video.Title = c.PostForm("title")
	video.UserID = user.ID
	video.PlayUrl = "http://49.233.7.107:8081/static/" + finalName
	video.CoverUrl = "https://wx1.sinaimg.cn/large/007WELPTly1h371ow92iqj30bf0gnjyd.jpg"
	video.CreateTime = time.Now()

	// fmt.Print("title = ", video.Title)
	db.Mysql.Save(&video)

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	var user db.User

	user_name, _ := c.Get("username")

	// fmt.Print("&&&", user_name, "&&&&")

	result := db.Mysql.Where(" Name = ?", user_name).Find(&user)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 2,
			"msg":  "该用户不存在",
		})
		return
	}

	var video []db.Video //数据库查询的评论列表
	if err := db.Mysql.Where("user_ID = ?", user.ID).Find(&video).Error; err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: video,
	})
}
