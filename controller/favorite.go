package controller

import (
	"fmt"
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
	userId, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	fmt.Printf("%d\n", userId)
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 8)

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
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Unlike success"})
		}
	default:
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "parameter error"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	favoriteList, _ := service.FavoriteList(userId)
	c.JSON(http.StatusOK, FavoriteListResponse{
		Response:  Response{StatusCode: 0},
		VideoList: favoriteList,
	})
}
