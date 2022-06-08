package service

import (
	"github.com/lyj0309/douyin/db"
	"gorm.io/gorm"
)

type FavoriteVideo struct {
	Id            int64    `json:"id,omitempty"`
	Author        *UserRes `json:"author"`
	PlayUrl       string   `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string   `json:"cover_url,omitempty"`
	FavoriteCount int64    `json:"favorite_count,omitempty"`
	CommentCount  int64    `json:"comment_count,omitempty"`
	IsFavorite    bool     `json:"is_favorite,omitempty"`
	Title         string   `json:"title,omitempty"`
}

func Like(userId int64, videoId int64) error {
	like := db.Like{
		UserID:  uint(userId),
		VideoID: uint(videoId),
	}
	// 检查是不是点过了
	var count int64
	if db.Mysql.Where(&like).Count(&count); count != 0 {
		return nil
	}
	db.Mysql.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&like).Error; err != nil {
			return err
		}
		err := tx.Model(&db.Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}
func Unlike(userId int64, videoId int64) error {
	db.Mysql.Transaction(func(tx *gorm.DB) error {
		err := db.Mysql.Where("user_id = ? AND video_id = ?", uint(userId), uint(videoId)).Delete(&db.Like{}).Error
		if err != nil {
			return err
		}
		err = tx.Model(&db.Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}

func FavoriteList(userId int64) ([]FavoriteVideo, error) {
	var likes []db.Like //数据库查询点赞列表
	err := db.Mysql.Where("user_id = ?", userId).Find(&likes).Error
	if err != nil {
		return nil, err
	}

	// 查询对应用户的所有点赞视频id
	var vid []uint
	for _, like := range likes {
		vid = append(vid, like.VideoID)
	}

	// 查询点赞视频表
	var videos []db.Video
	db.Mysql.Where("ID IN ?", vid).Find(&videos)

	// 确定视频的作者ID 可能会有重复ID
	var authorsId []uint
	for _, video := range videos {
		authorsId = append(authorsId, video.UserID)
	}
	// 查询这些视频作者的信息 并建立映射关系
	var authorRes []subUser
	// 会去重
	db.Mysql.Model(&db.User{}).Find(&authorRes, authorsId)
	m := make(map[int64]*subUser)
	// 无重复的一个映射 作者名 到作者信息
	for i, author := range authorRes {
		m[author.Id] = &authorRes[i]
	}

	var favoriteVideos []FavoriteVideo
	for _, video := range videos {
		userVo := UserRes{
			subUser:  m[int64(video.UserID)],
			IsFollow: IsFollow(Itoa(userId), Itoa(video.ID)),
		}
		videoVo := FavoriteVideo{
			Id:            int64(video.ID),
			Author:        &userVo,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: int64(video.FavoriteCount),
			CommentCount:  int64(video.CommentCount),
			IsFavorite:    true,
			Title:         video.Title,
		}
		favoriteVideos = append(favoriteVideos, videoVo)
	}
	return favoriteVideos, nil
}
