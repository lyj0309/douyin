package service

import (
	"fmt"
	"testing"
)

func TestCommentAdd(t *testing.T) {
	//CommentAdd(1, 1, "Hello!")
	//CommentDelete(1, 1)
	//CommentAdd(1, 1, "点赞")
	//CommentAdd(2, 1, "评论")
	//CommentAdd(2, 1, "收藏")
	var comm []Comment
	comm, _ = CommentList(1, "1")
	fmt.Println(comm[1])
}
