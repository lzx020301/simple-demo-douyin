package controller

import (
	"fmt"

	"github.com/RaymondCode/simple-demo/models"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	VideoId       int64  `json:"id,omitempty"`
	User          User   `json:"author" gorm:"embedded"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	UserId        int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

func CommentsChange(commentList []models.Comment2) []Comment {
	var comments []Comment
	for _, comment := range commentList {
		comments = append(comments, CommentChange(comment))
	}
	return comments
}

func CommentChange(comment models.Comment2) Comment {
	user, _ := GetUserByID(comment.UserID)
	return Comment{
		Id:         int64(comment.ID),
		Content:    comment.Content,
		CreateDate: comment.CreatedAt.Format("01-02"),
		User:       *user,
	}
}

func GetUserByID(userID uint) (*User, error) {
	var user *User
	GLOBAL_DB.Model(&User{}).Where("user_id = ?", userID).First(&user)
	if user == nil {
		return nil, fmt.Errorf("未找到用户")
	}
	return user, nil
}
