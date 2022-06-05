package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	//"strconv"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

var VideoLists = make(map[User][]Video)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.UserId, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	Bucket.PutObjectFromFile("public/" + filename ,saveFile)

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}


// PublishList all users have same publish video list
func PublishList(c *gin.Context) {

	GLOBAL_DB.AutoMigrate(&Video{})
	token := c.Query("token")
	//userId := c.Query("user_id")
	var video []Video
	

	//_ ,exist := UsersLoginInfo[token]

	//userid, _ := strconv.ParseInt(userId, 10, 64)

	var userid Token2ID
	GLOBAL_DB.Where("token = ?" ,token).Find(&userid)

	GLOBAL_DB.Where("user_id = ?" ,userid.ID).Find(&video)

	//VideoLists[UserIDinfo[userid]] = DemoVideos

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{StatusCode: 0},
		//videoList: VideoLists[UserIDinfo[userid]],
		 //VideoList: VideoLists[UserIDinfo[userid]],
		//VideoList: Videolist,
		VideoList: video,
	})
}