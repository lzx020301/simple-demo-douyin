package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"sync/atomic"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"yfldouyin": {
		UserId:            1,
		Name:          "游飞龙",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var UserIDinfo = make(map[int64]User)

var UserIDsequence int64


//var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	var user User
	GLOBAL_DB.Model(&User{}).Last(&user)
	UserIDsequence = user.UserId + 1
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		//atomic.AddInt64(&userIdSequence, 1)
		newUser := User{
			UserId:   UserIDsequence,
			Name: username,
		}
		newToken2ID := Token2ID{
			Token: token,
			ID: UserIDsequence,
		}
		GLOBAL_DB.Create(&newUser)
		GLOBAL_DB.Create(&newToken2ID)
		usersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   UserIDsequence,
			Token:    username + password,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.UserId,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

type Token2ID struct{
	Token string
	ID int64
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	var id Token2ID
	GLOBAL_DB.Where("token = ?" ,token).Find(&id)

	var users User

	GLOBAL_DB.First(&users ,id.ID)

	//user, exist := usersLoginInfo[token]

	// if  exist {
	// 	c.JSON(http.StatusOK, UserResponse{
	// 		Response: Response{StatusCode: 0},
	// 		User:     users,
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, UserResponse{
	// 		Response: Response{StatusCode: 1, StatusMsg: "请先登录"},
	// 	})
	// }

	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User:     users,
	})
}