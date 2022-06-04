package service

import (
	"fmt"
	"testing"
)

func TestFavorite(t *testing.T) {
	Unlike(1, 2)
	Unlike(1, 3)
	Unlike(2, 3)
	Like(1, 2)
	Like(1, 3)
	Like(2, 3)
	vid, _ := VideoList(1)
	fmt.Println(len(vid))
	for _, v := range vid {
		fmt.Println(v.Id)
	}
	//CommentDelete(1, 1)
	//CommentAdd(1, 1, "点赞")
	//CommentAdd(2, 1, "评论")
	//CommentAdd(2, 1, "收藏")
}
