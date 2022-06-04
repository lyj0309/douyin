package service

import (
	"github.com/lyj0309/douyin/db"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	Id         int64   `json:"id,omitempty"`
	User       UserRes `json:"user"`
	Content    string  `json:"content,omitempty"`
	CreateDate string  `json:"create_date,omitempty"`
}

func CommentAdd(userId int64, videoId int64, commentText string) (Comment, error) {
	comment := db.Comment{
		UserID:     uint(userId),
		VideoID:    uint(videoId),
		Content:    commentText,
		CreateTime: time.Now(),
		Delete:     false,
	}
	db.Mysql.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&comment).Error; err != nil {
			return err
		}
		err := tx.Model(&db.Video{}).Where("id = ?", comment.VideoID).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
		if err != nil {
			return err
		}
		return nil
	})
	var user db.User
	db.Mysql.Where("id = ?", userId).Find(&user)
	userVo := UserRes{
		subUser: &subUser{
			Id:            int64(user.ID),
			Name:          user.Name,
			FollowCount:   int64(user.FollowCount),
			FollowerCount: int64(user.FollowerCount),
		},
		IsFollow: false, //自己对自己关注与否？
	}
	commentVo := Comment{
		Id:         int64(comment.ID),
		User:       userVo,
		Content:    comment.Content,
		CreateDate: comment.CreateTime.Format("2006-01-02 15:04:05"),
	}
	return commentVo, nil
}
func CommentDelete(commentId, videoId int64) error {
	db.Mysql.Transaction(func(tx *gorm.DB) error {
		err := db.Mysql.Model(&db.Comment{}).Where("id = ?", commentId).Update("`delete`", true).Error
		if err != nil {
			return err
		}
		err = tx.Model(&db.Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}

func CommentList(videoId int64, userIdStr string) ([]Comment, error) {
	var comments []db.Comment //数据库查询的评论列表
	err := db.Mysql.Where("video_id = ?", videoId).Where("`delete` = ?", 0).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	var uid []uint
	for _, comment := range comments {
		uid = append(uid, comment.UserID)
	}
	//一次查询所有用户，减少查询次数
	var users []subUser
	db.Mysql.Model(&db.User{}).Find(&users, uid)

	m := make(map[int64]*subUser)
	for i, user := range users {
		m[user.Id] = &users[i]
	}
	//->评论列表转换
	var commentList []Comment //前端评论列表标准
	for _, comment := range comments {
		userVo := UserRes{
			subUser:  m[int64(comment.UserID)],
			IsFollow: IsFollow(userIdStr, Itoa(comment.UserID)),
		}
		commentVo := Comment{
			Id:         int64(comment.ID),
			User:       userVo,
			Content:    comment.Content,
			CreateDate: comment.CreateTime.Format("2006-01-02 15:04:05"),
		}
		commentList = append(commentList, commentVo)

	}
	return commentList, nil
}
