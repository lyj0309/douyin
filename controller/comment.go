package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lyj0309/douyin/service"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	userIdStr := c.PostForm("user_id")
	videoIdStr := c.PostForm("video_id")
	actionTypeStr := c.PostForm("action_type")
	commentText := c.PostForm("comment_text")
	commentIdStr := c.PostForm("comment_id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	actionType, _ := strconv.ParseInt(actionTypeStr, 10, 8)
	commentId, _ := strconv.ParseInt(commentIdStr, 10, 64)
	if actionType == 1 {
		service.CommentAdd(userId, videoId, commentText)
		c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "add Success"})
		return
	} else if actionType == 2 {
		service.CommentDelete(commentId)
		c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "delete Success"})
		return
	}

}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	userIdStr := c.GetString(ctxUidKey)
	videoIdStr := c.PostForm("video_id")
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	commentList, _ := service.CommentList(videoId, userIdStr)
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: commentList,
	})
}
