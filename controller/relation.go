package controller

import (
	//"fmt"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

type FromFollowList struct {
	User        User `gorm:"embedded"`
	From_userid int64
}

type Followlist struct {
	User      User `gorm:"embedded"`
	To_userid int64
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	touserid := c.Query("to_user_id")
	var touser User
	var user User
	var token2id Token2ID
	var fromfollowlist FromFollowList
	to_userid, _ := strconv.ParseInt(touserid, 10, 64)
	GLOBAL_DB.Model(&Token2ID{}).Where("token = ?", token).Find(&token2id)
	GLOBAL_DB.Where("user_id = ?", to_userid).Find(&touser)
	GLOBAL_DB.Where("user_id = ?", token2id.ID).Find(&user)
	GLOBAL_DB.Where("from_userid = ?", token2id.ID).Find(&fromfollowlist)
	if fromfollowlist.User.UserId == to_userid {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "取消关注",
		})
		GLOBAL_DB.Where("to_userid = ?", to_userid).Delete(&Followlist{})
		GLOBAL_DB.Where("from_userid = ?", user.UserId).Delete(&FromFollowList{})
		GLOBAL_DB.Model(&User{}).Where("user_id = ?", to_userid).Update("is_follow", 0)
		GLOBAL_DB.Model(&User{}).Where("user_id = ?", token2id.ID).Update("follow_count", gorm.Expr("follow_count - 1"))
	} else {
		GLOBAL_DB.Model(&User{}).Where("user_id = ?", to_userid).Update("is_follow", 1)
		GLOBAL_DB.Create(&Followlist{User: user, To_userid: to_userid})
		GLOBAL_DB.Model(&User{}).Where("user_id = ?", token2id.ID).Update("follow_count", gorm.Expr("follow_count + 1"))
		GLOBAL_DB.Create(&FromFollowList{User: touser, From_userid: user.UserId})
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  "关注成功",
		})
	}

	// if _, exist := usersLoginInfo[token]; exist {
	// 	c.JSON(http.StatusOK, Response{StatusCode: 0})
	// } else {
	// 	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	// }
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	GLOBAL_DB.AutoMigrate(&Followlist{})
	GLOBAL_DB.AutoMigrate(&FromFollowList{})
	token := c.Query("token")
	var token2id Token2ID
	GLOBAL_DB.Where("token = ?", token).Find(&token2id)
	var userlist User
	var followuserlist []Followlist
	var count int64
	//GLOBAL_DB.Where("user_id = ?" ,userid).Find(&followuserlist).Count(&count)
	fmt.Println(token2id.ID)
	GLOBAL_DB.Model(&Followlist{}).Where("user_id = ?", token2id.ID).Find(&followuserlist).Distinct("to_userid").Count(&count).Group("user_id")
	userlist2 := make([]User, count)
	var i int
	for i = 0; i < int(count); i++ {
		GLOBAL_DB.Model(&User{}).Where("user_id = ?", followuserlist[i].To_userid).Find(&userlist)
		userlist2[i] = userlist
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		// UserList: []User{DemoUser},
		UserList: userlist2,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	var user []User

	GLOBAL_DB.Find(&user)
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: user,
	})
}
