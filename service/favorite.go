package service

import (
	"github.com/lyj0309/douyin/db"
)

func Like(userId int64, videoId int64) error {
	like := db.Like{
		UserID:  uint(userId),
		VideoID: uint(videoId),
	}
	if err := db.Mysql.Create(&like).Error; err != nil {
		return err
	}
	return nil
}
func Unlike(userId int64, videoId int64) error {
	err := db.Mysql.Where("user_id = ? AND video_id = ?", uint(userId), uint(videoId)).Delete(&db.Like{}).Error
	if err != nil {
		return err
	}
	return nil
}

func VideoList(userId int64) ([]Video, error) {
	var likes []db.Like //数据库查询点赞列表
	err := db.Mysql.Where("user_id = ?", userId).Find(&likes).Error
	if err != nil {
		return nil, err
	}

	var videoList []Video
	for _, comment := range comments {
		var user db.User
		db.Mysql.Where("id = ?", comment.UserID).Find(&user)
		userVo := User{
			Id:            int64(user.ID),
			Name:          user.Name,
			FollowCount:   int64(user.FollowCount),
			FollowerCount: int64(user.FollowerCount),
			IsFollow:      IsFollow(userIdStr, Itoa(user.ID)), //关注功能逻辑
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
