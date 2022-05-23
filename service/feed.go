package service

import (
	"fmt"
	"github.com/lyj0309/douyin/db"
	"time"
)

type feedSqlRes struct {
	ID             uint
	UserID         uint
	PlayUrl        string
	CreateTime     time.Time
	CoverUrl       string
	CommentCount   int
	FavouriteCount int
}

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
	var res []feedSqlRes
	db.Mysql.Raw(`SELECT videos.id ,videos.user_id,videos.play_url,videos.create_time,videos.cover_url
,comm_tmp.comment_count,like_tmp.favorite_count

FROM videos

LEFT JOIN (
	SELECT count(*) AS comment_count,video_id
	FROM comments
	GROUP BY comments.video_id
) AS comm_tmp
ON videos.id = comm_tmp.video_id

LEFT JOIN (
	SELECT COUNT(*) AS favorite_count,video_id
	FROM likes
	GROUP BY likes.video_id
) AS like_tmp
ON videos.id = like_tmp.video_id
ORDER BY videos.create_time
LIMIT 30
`).Scan(&res)

	fmt.Println(res, len(res))

	var uids []uint
	for _, re := range res {
		uids = append(uids, re.UserID)
	}

	fmt.Println(len(uids))
	authors := GetUserRes(userID, uids)

	fmt.Println("authors", authors, len(*authors))
	var videos []Video
	for i, re := range res {
		fmt.Println(i)
		var count int64

		if userID != "" {
			db.Mysql.Where(`user_id = ? AND video_id = ?`, userID, re.ID).Count(&count)
		}
		isfav := false
		if count == 1 {
			isfav = true
		}

		videos = append(videos, Video{
			Id:            int64(re.ID),
			Author:        &(*authors)[i],
			PlayUrl:       re.PlayUrl,
			CoverUrl:      re.CoverUrl,
			FavoriteCount: int64(re.FavouriteCount),
			CommentCount:  int64(re.CommentCount),
			IsFavorite:    isfav,
		})
	}

	return &videos
}
