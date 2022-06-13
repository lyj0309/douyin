package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lyj0309/douyin/service"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList *[]service.Video `json:"video_list,omitempty"`
	NextTime  int64            `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	var latestTime time.Time
	timeStampStr := c.Query("latest_time")
	if timeStampStr == "" {
		latestTime = time.Now()
	} else {
		ts, err := strconv.Atoi(timeStampStr)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "时间戳解析错误",
			})
			return
		} else {
			latestTime = time.Unix(int64(ts), 0)
		}
	}
	res := service.Feed(latestTime, "")

	if res == nil || len(*res) == 0 {
		res = &DemoVideos
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: res,
		NextTime:  time.Now().Unix(),
	})
}
