package db

import "time"

// User 用户表
type User struct {
	Id            int64 `gorm:"primaryKey;autoIncrement'"`
	Name          string
	Password      string
	FollowCount   int64     //关注总数
	FollowerCount int64     //粉丝总数
	RegisterTime  time.Time //注册时间
}

// Comment 评论表
type Comment struct {
	Id         int64     `gorm:"primaryKey;autoIncrement'"`
	UserId     int64     //评论者
	VideoId    int64     //评论视频
	Content    string    //内容
	CreateTime time.Time //评论时间
	Delete     bool      //是否删除
}

// Like 点赞表(待讨论)
type Like struct {
	UserId  uint `gorm:"primaryKey"`
	VideoId uint `gorm:"primaryKey"`
}

// Video 视频表
type Video struct {
	Id         int64 `gorm:"primaryKey;autoIncrement'"`
	UserId     int64
	PlayUrl    string    //视频播放地址
	CoverUrl   string    //视频封面地址
	CreateTime time.Time //创建时间
}
