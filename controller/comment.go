package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lyj0309/douyin/service"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	Response
	CommentList []service.Comment `json:"comment_list,omitempty"`
}
type CommentResponse struct {
	Response
	Comment service.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {

	userId, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType := c.Query("action_type")
	if actionType == "1" {
		commentText := c.Query("comment_text")
		if len(commentText) == 0 {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "评论内容不能为空"})
			return
		}
		commentVo, err := service.CommentAdd(userId, videoId, commentText)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		c.JSON(http.StatusOK, CommentResponse{
			Response: Response{StatusCode: 0, StatusMsg: "add Success"},
			Comment:  commentVo,
		})
	} else if actionType == "2" {
		commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err := service.CommentDelete(commentId, videoId); err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "delete Success"})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "参数有误"})
	}

}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	userIdStr := c.GetString("user_id")
	videoIdStr := c.Query("video_id")
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "参数有误"})
		return
	}
	var commentList []service.Comment
	commentList, err = service.CommentList(videoId, userIdStr)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: commentList,
	})
}
