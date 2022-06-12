package controller

import (
	"net/http"
	"strconv"

	//"strconv"

	"github.com/RaymondCode/simple-demo/models"
	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	// var user *models.User
	var request struct {
		VideoID     uint   `form:"video_id" binding:"required"`
		ActionType  uint   `form:"action_type" binding:"required,min=1,max=2"`
		CommentText string `form:"comment_text" binding:"omitempty"`
		CommentID   uint   `form:"comment_id" binding:"omitempty"`
	}

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	var token2id Token2ID
	token := c.Query("token")
	GLOBAL_DB.Model(&Token2ID{}).Where("token = ?", token).First(&token2id)

	//写评论
	if request.ActionType == 1 {
		comment := CreateComment(&models.Comment2{
			UserID:  uint(token2id.ID),
			VideoID: request.VideoID,
			Content: request.CommentText,
		})
		c.JSON(http.StatusOK, CommentResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "添加评论成功",
			},
			Comment: CommentChange(*comment),
		})
	} else if request.ActionType == 2 {
		//删评论
		DeleteComment(request.CommentID)
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "评论删除成功",
		})
	}
}

type Commentlist struct {
	User    User    `gorm:"embedded"`
	Comment Comment `grom:"embeded"`
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	var comment []models.Comment2
	videoid := c.Query("video_id")
	video_id, _ := strconv.ParseInt(videoid, 10, 64)
	GLOBAL_DB.Model(&models.Comment2{}).Where("video_id = ?", video_id).Find(&comment)
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: CommentsChange(comment),
	})
}
