package service

import (
	"github.com/lyj0309/douyin/db"
	"time"
)

type Video struct {
	Id            int64    `json:"id,omitempty"`
	Author        *UserRes `json:"author"`
	PlayUrl       string   `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string   `json:"cover_url,omitempty"`
	FavoriteCount int64    `json:"favorite_count,omitempty"`
	CommentCount  int64    `json:"comment_count,omitempty"`
	IsFavorite    bool     `json:"is_favorite,omitempty"`
}

func Feed(latestTime time.Time, userID string) *[]Video {
	var videos []db.Video
	//db.Mysql.Where(`create_time < ?`, latestTime).Limit(30).Find(&videos)
	db.Mysql.Limit(30).Find(&videos)

	//fmt.Println(res, len(res))

	var uids []uint
	for _, re := range videos {
		uids = append(uids, re.UserID)
	}

	//fmt.Println(len(uids))
	authors := GetUserRes(userID, uids)

	//fmt.Println("authors", authors, len(*authors))
	var videores []Video
	for i, re := range videos {
		var count int64

		if userID != "" {
			db.Mysql.Where(`user_id = ? AND video_id = ?`, userID, re.ID).Count(&count)
		}
		isfav := false
		if count == 1 {
			isfav = true
		}

		videores = append(videores, Video{
			Id:            int64(re.ID),
			Author:        &(*authors)[i],
			PlayUrl:       re.PlayUrl,
			CoverUrl:      re.CoverUrl,
			FavoriteCount: int64(re.FavoriteCount),
			CommentCount:  int64(re.CommentCount),
			IsFavorite:    isfav,
		})
	}

	return &videores
}
