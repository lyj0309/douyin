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
type CommentAct struct {
	UserId      string
	VideoId     string
	ActionType  string
	CommentText string
	CommentId   string
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	var commentAct CommentAct
	if err := c.ShouldBind(&commentAct); err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	userId, _ := strconv.ParseInt(commentAct.UserId, 10, 64)
	videoId, _ := strconv.ParseInt(commentAct.VideoId, 10, 64)
	commentId, _ := strconv.ParseInt(commentAct.CommentId, 10, 64)

	if commentAct.ActionType == "1" {
		commentVo, err := service.CommentAdd(userId, videoId, commentAct.CommentText)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		c.JSON(http.StatusOK, CommentResponse{
			Response: Response{StatusCode: 0, StatusMsg: "add Success"},
			Comment:  commentVo,
		})
	} else if commentAct.ActionType == "2" {
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
	userIdStr := c.GetString(ctxUidKey)
	videoIdStr := c.Query("video_id")
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	commentList, _ := service.CommentList(videoId, userIdStr)
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: commentList,
	})
}
