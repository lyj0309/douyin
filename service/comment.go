package service

import (
	"github.com/lyj0309/douyin/controller"
	"github.com/lyj0309/douyin/db"
	"time"
)

func CommentAdd(userId int64, videoId int64, commentText string) error {
	comment := db.Comment{
		UserID:     uint(userId),
		VideoID:    uint(videoId),
		Content:    commentText,
		CreateTime: time.Now(),
		Delete:     false,
	}
	if err := db.Mysql.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}
func CommentDelete(commentId int64) error {
	err := db.Mysql.Model(&db.Comment{}).Where("id = ?", commentId).Update("delete", true).Error
	if err != nil {
		return err
	}
	return nil
}

func CommentList(videoId int64, userIdStr string) ([]controller.Comment, error) {
	var comments []db.Comment //数据库查询的评论列表
	err := db.Mysql.Where("video_id = ?", videoId).Where("delete = ?", false).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	//->评论列表转换
	var commentList []controller.Comment //前端评论列表标准
	for _, comment := range comments {
		var user db.User
		db.Mysql.Where("id = ?", comment.UserID).Find(&user)
		userVo := controller.User{
			Id:            int64(user.ID),
			Name:          user.Name,
			FollowCount:   int64(user.FollowCount),
			FollowerCount: int64(user.FollowerCount),
			IsFollow:      IsFollow(userIdStr, Itoa(user.ID)), //关注功能逻辑
		}
		commentVo := controller.Comment{
			Id:         int64(comment.ID),
			User:       userVo,
			Content:    comment.Content,
			CreateDate: comment.CreateTime.Format("2006-01-02 15:04:05"),
		}
		commentList = append(commentList, commentVo)

	}
	return commentList, nil
}
