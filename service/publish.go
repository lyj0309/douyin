package service

import (
	"time"

	"github.com/lyj0309/douyin/db"
)

// type User struct {
// 	ID            int64  `json:"id,omitempty"`
// 	Name          string `json:"name,omitempty"`
// 	FollowCount   int64  `json:"follow_count,omitempty"`
// 	FollowerCount int64  `json:"follower_count,omitempty"`
// 	IsFollow      bool   `json:"is_follow,omitempty"`
// }
type Video_list struct {
	Id int64 `json:"id,omitempty"`
	//Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

func UserExit(username string) error {
	var user db.User
	if err := db.Mysql.Where("user_name = ?", username).Find(&user).Error; err != nil {
		return err
	}
	return nil
}

func VideoAdd(username string, videotitle string, play_url string, cover_url string) error {

	var user db.User
	if err := db.Mysql.Where("user_name = ?", username).Find(&user).Error; err != nil {
		return err
	}

	video := db.Video{
		Title:      videotitle,
		UserID:     user.ID,
		PlayUrl:    play_url,
		CoverUrl:   cover_url,
		CreateTime: time.Now(),
	}
	if err := db.Mysql.Create(&video).Error; err != nil {
		return err
	}
	return nil
}

func VideoList(username string, videoId int64, userIdStr string) ([]Video_list, error) {
	var user db.User
	if err := db.Mysql.Where("user_name = ?", username).Find(&user).Error; err != nil {
		return nil, err
	}

	var videos []db.Video
	err := db.Mysql.Where("user_id = ?", user.ID).Find(&videos).Error
	if err != nil {
		return nil, err
	}

	var videoList []Video_list
	for _, video := range videos {
		var user db.User

		db.Mysql.Where("id = ?", video.UserID).Find(&user)
		userVo := User{
			Id:            int64(user.ID),
			Name:          user.Name,
			FollowCount:   int64(user.FollowCount),
			FollowerCount: int64(user.FollowerCount),
			IsFollow:      IsFollow(userIdStr, Itoa(user.ID)),
		}
		videoVo := Video_list{
			Id:       int64(video.ID),
			Author:   userVo,
			PlayUrl:  video.PlayUrl,
			CoverUrl: video.CoverUrl,
		}
		videoList = append(videoList, videoVo)

	}
	return videoList, nil
}
