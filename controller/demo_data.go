package controller

import "github.com/lyj0309/douyin/service"

var DemoVideos = []service.Video{
	{
		Id:            1,
		Author:        &DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = service.UserRes{
	IsFollow: false,
}

func init() {
	DemoUser.Id = 1
	DemoUser.Name = "TestUser"
	DemoUser.FollowCount = 0
	DemoUser.FollowerCount = 0
}
