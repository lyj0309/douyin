package db

import "time"

// User 用户表
type User struct {
	ID            uint `gorm:"primaryKey;autoIncrement'"`
	Name          string
	Password      string
	FollowCount   int       //关注总数
	FollowerCount int       //粉丝总数
	RegisterTime  time.Time //注册时间
}

// Comment 评论表
type Comment struct {
	ID         uint      `gorm:"primaryKey;autoIncrement'"`
	UserID     uint      //评论者
	VideoID    uint      //评论视频
	Content    string    //内容
	CreateTime time.Time //评论时间
	Delete     bool      //是否删除
}

// Like 点赞表(待讨论)
type Like struct {
	UserID  uint `gorm:"primaryKey"`
	VideoID uint `gorm:"primaryKey"`
}

// Video 视频表
type Video struct {
	ID            uint   `gorm:"primaryKey;autoIncrement'"`
	Title         string //视频标题
	UserID        uint
	PlayUrl       string    //视频播放地址
	CoverUrl      string    //视频封面地址
	FavoriteCount uint      //点赞总数
	CommentCount  uint      //评论总数
	CreateTime    time.Time //创建时间
}

//Relation 关注表,使用redis实现
//type Relation struct {
//	UserID   uint `gorm:"primaryKey"`
//	ToUserID uint `gorm:"primaryKey"`
//}
