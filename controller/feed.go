package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

func Feed(c *gin.Context){
	c.JSON(http.StatusOK ,FeedResponse{
		Response: Response{StatusCode: 0 ,StatusMsg: "加载成功"},
		VideoList: DemoVideos,
		NextTime: time.Now().Unix(),
	})
}