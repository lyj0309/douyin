package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lyj0309/douyin/service"
	"net/http"
)

type UserListResponse struct {
	Response
	UserList *[]service.UserRes `json:"user_list"`
}

// RelationAction 关注
func RelationAction(c *gin.Context) {
	// userID 是在jwt中写入上下文的
	//userID := c.GetString("user_id")
	userID := c.Query("user_id")
	//获取参数
	toUserID := c.Query("to_user_id")
	actionType := c.Query("action_type")

	//参数校验
	if toUserID == "" {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "参数错误"})
		return
	}
	if actionType != "1" && actionType != "2" {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "参数错误"})
		return
	}

	err := service.RelationAction(userID, toUserID, actionType)

	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{StatusCode: 0})

}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {

	//userList := service.FollowList(c.GetString("user_id"))
	userList := service.FollowList(c.Query("user_id"))

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: userList,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	userList := service.FollowerList(c.GetString("user_id"))

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: userList,
	})
}
