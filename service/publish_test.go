package service

import (
	"time"

	"github.com/lyj0309/douyin/db"
)

// type User struct {
// 	Id            int64  `json:"id,omitempty"`
// 	Name          string `json:"name,omitempty"`
// 	FollowCount   int64  `json:"follow_count,omitempty"`
// 	FollowerCount int64  `json:"follower_count,omitempty"`
// 	IsFollow      bool   `json:"is_follow,omitempty"`
// }
// type Comment struct {
// 	Id         int64  `json:"id,omitempty"`
// 	User       User   `json:"user"`
// 	Content    string `json:"content,omitempty"`
// 	CreateDate string `json:"create_date,omitempty"`
// }
// type Comment struct {
// 	ID         uint      `gorm:"primaryKey;autoIncrement'"`
// 	UserID     uint      //评论者
// 	VideoID    uint      //评论视频
// 	Content    string    //内容
// 	CreateTime time.Time //评论时间
// 	Delete     bool      //是否删除
// }
// type Video struct {
// 	ID         uint   `gorm:"primaryKey;autoIncrement'"`
// 	Title      string //视频标题
// 	UserID     uint
// 	PlayUrl    string    //视频播放地址
// 	CoverUrl   string    //视频封面地址
// 	CreateTime time.Time //创建时间
// }

func VideoAdd(userId int64, play_url string, video_title string, cover_url string) error {
	// comment := db.Comment{
	// 	UserID:     uint(userId),
	// 	VideoID:    uint(videoId),
	// 	Content:    commentText,
	// 	CreateTime: time.Now(),
	// 	Delete:     false,
	// }
	video := db.Video{
		Title:      video_title, //视频标题
		UserID:     uint(userId),
		PlayUrl:    play_url,   //视频播放地址
		CoverUrl:   cover_url,  //视频封面地址
		CreateTime: time.Now(), //创建时间
	}
	if err := db.Mysql.Create(&video).Error; err != nil {
		return err
	}
	return nil
}

// func CommentDelete(commentId int64) error {
// 	err := db.Mysql.Model(&db.Comment{}).Where("id = ?", commentId).Update("delete", true).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func VideoList(userIdStr string) ([]Video, error) {
	var videos []db.Video //数据库查询的列表
	err := db.Mysql.Where("user_id = ?", userIdStr).Where("delete = ?", false).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	//->评论列表转换
	var videoList []Video //前端评论列表标准
	for _, video := range videos {
		var user db.User
		db.Mysql.Where("id = ?", video.UserID).Find(&user)
		// userVo := User{
		// 	Id:            int64(user.ID),
		// 	Name:          user.Name,
		// 	FollowCount:   int64(user.FollowCount),
		// 	FollowerCount: int64(user.FollowerCount),
		// 	IsFollow:      IsFollow(userIdStr, Itoa(user.ID)), //关注功能逻辑
		// }
		videoVo := Video{
			Id: int64(video.ID),
			// User:       userVo,
			// Content:    video.Content,
			// CreateDate: video.CreateTime.Format("2006-01-02 15:04:05"),
		}
		videoList = append(videoList, videoVo)

	}
	return videoList, nil
}
