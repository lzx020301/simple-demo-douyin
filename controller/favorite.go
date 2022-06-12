package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserFavoriteVideos struct {
	User  User  `gorm:"embedded"`
	Video Video `gorm:"embedded"`
}

// FavoriteAction no practical effect, just check if token is valid
type FavoriteActionRequest struct {
	UserID     uint `form:"user_id" binding:"required"`
	VideoID    uint `form:"video_id" binding:"required"`
	ActionType uint `form:"action_type" binding:"required,min=1,max=2"`
}

///favorite/action/
func FavoriteAction(c *gin.Context) {

	GLOBAL_DB.AutoMigrate(&UserFavoriteVideos{})

	var token2id Token2ID
	var user User
	var userfavor UserFavoriteVideos
	token := c.Query("token")
	actionType := c.Query("action_type")
	GLOBAL_DB.Model(&Token2ID{}).Where("token = ?", token).First(&token2id)
	GLOBAL_DB.Model(&User{}).Where("user_id = ?", token2id.ID).First(&user)
	videoid := c.Query("video_id")
	//uid := user.UserId
	vid, _ := strconv.ParseInt(videoid, 10, 64)
	GLOBAL_DB.Model(&UserFavoriteVideos{}).Where("user_id = ? AND video_id = ?", user.UserId, vid).Find(&userfavor)
	actiontype, _ := strconv.ParseInt(actionType, 10, 64)
	if actiontype == 2 {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "取消点赞",
		})
		GLOBAL_DB.Where("user_id = ? AND video_id = ?", user.UserId, vid).Delete(&UserFavoriteVideos{})
	} else {

		var video Video
		GLOBAL_DB.Model(&Video{}).Where("video_id = ?", vid).First(&video)

		GLOBAL_DB.Model(&UserFavoriteVideos{}).Create(&UserFavoriteVideos{User: user, Video: video})

		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "点赞成功",
		})
	}

}

// FavoriteList all users have same favorite video list
///favorite/list/
func FavoriteList(c *gin.Context) {
	var token2id Token2ID

	// uid, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	GLOBAL_DB.Model(&Token2ID{}).Where("token = ?", token).Find(&token2id)

	var video []Video

	GLOBAL_DB.Model(&UserFavoriteVideos{}).Where("user_id = ?", token2id.ID).Find(&video)

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: video,
	})
}
