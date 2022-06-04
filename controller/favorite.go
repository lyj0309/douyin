package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lyj0309/douyin/service"
	"net/http"
	"strconv"
)

type FavoriteListResponse struct {
	Response
	VideoList []service.FavoriteVideo `json:"video_list,omitempty"`
}

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	userIdStr := c.PostForm("user_id")
	// userIdStr := c.GetString(ctxUidKey)
	videoIdStr := c.PostForm("video_id")
	actionTypeStr := c.PostForm("action_type")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	actionType, _ := strconv.ParseInt(actionTypeStr, 10, 8)

	// 验证用户
	switch actionType {
	case 1:
		err := service.Like(userId, videoId)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		} else {
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Like success"})
		}
	case 2:
		err := service.Unlike(userId, videoId)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		} else {
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Like success"})
		}
	default:
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "parameter error"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	// userIdStr := c.GetString(ctxUidKey)
	userIdStr := c.Query("user_id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	videoList, _ := service.VideoList(userId)
	c.JSON(http.StatusOK, FavoriteListResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoList,
	})
}
