package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	userIdStr := c.PostForm("user_id")
	token := c.PostForm("token")
	videoIdStr := c.PostForm("video_id")
	actionTypeStr := c.PostForm("action_type")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	actionType, _ := strconv.ParseInt(actionTypeStr, 10, 8)

	// 验证用户
	if uid, err := JwtParseUser(token); err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	} else {
		if uid != userIdStr {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't approved"})
			return
		}
	}
	switch actionType {
	case 1:

	case 2:

	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
