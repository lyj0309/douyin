package controller

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/lyj0309/douyin/db"
)

type VideoListResponse struct {
	Response
	VideoList []db.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {

	userIdStr := c.GetString(ctxUidKey)

	// NEW
	var user db.User

	res := db.Mysql.Where("id = ?", userIdStr).Find(&user)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "User doesn't exist",
		})
		return
	}

	// NEW END

	// if _, exist := usersLoginInfo[token]; !exist {
	// 	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	// 	return
	// }
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	// user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.ID, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// service.videoAdd(userId, videoId, commentText)

	// newVideo := db.Video{

	// 	Title:      title, //视频标题
	// 	UserID:     user.ID,
	// 	PlayUrl:    "http://10.113.166.143:8080/static/" + finalName,                  //视频播放地址
	// 	CoverUrl:   "http://10.113.166.143:8080/static/20190330000110_360x480_55.jpg", //视频封面地址
	// 	CreateTime: time.Now(),                                                        //创建时间
	// 	// FavoriteCount: 0,
	// 	// CommentCount:  0,
	// }
	// db.Mysql.Save(&newVideo)

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	var video []db.Video
	db.Mysql.Find(&video)
	// fmt.Print(video)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: video,
	})
}
