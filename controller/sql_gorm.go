package controller

import (
	"fmt"

	"gorm.io/driver/mysql"
	"github.com/RaymondCode/simple-demo/models"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var GLOBAL_DB *gorm.DB

func SqlLianjie() {
	//连接数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:Yfl20020301@tcp(rm-bp1bux16i07sc08d50o.mysql.rds.aliyuncs.com:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 171,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		SkipDefaultTransaction: false,
	})

	GLOBAL_DB = db

	db.AutoMigrate(&models.User2{}, &models.Comment2{})

	if err != nil {
		fmt.Println(err)
	}
}

func UserLikeVedio(uid int64, vid int64) error {

	var user *models.User2
	var video *models.Video2

	//检查用户和视频是否合法
	GLOBAL_DB.First(&user, uid)
	GLOBAL_DB.First(&video, vid)

	if int64(user.ID) != uid || int64(video.ID) != vid {
		return fmt.Errorf("不存在该用户或视频")

	}
	// 用户点赞列表添加视频id
	GLOBAL_DB.Model(user).Association("FavoriteVideos").Append(video)
	return nil
}

func UserDislikeVedio(uid int64, vid int64) error {

	var user *models.User2
	var video *models.Video2

	GLOBAL_DB.First(&user, uid)
	GLOBAL_DB.First(&video, vid)
	if int64(user.ID) != uid || int64(video.ID) != vid {
		return fmt.Errorf("不存在该用户或视频")

	}
	GLOBAL_DB.Model(&user).Association("FavoriteVideos").Delete(&video)
	return nil
}

func GetFavoriteVideosByUserID(uid int64) []models.Video2 {
	var user *models.User2
	GLOBAL_DB.Preload("FavoriteVideos").Find(&user, uid)
	return user.FavoriteVideos
}

//创建评论
func CreateComment(comment *models.Comment2) *models.Comment2 {
	GLOBAL_DB.Create(&comment).Commit()
	return comment
}

//删除评论
func DeleteComment(commentID uint) {
	var comment *models.Comment2
	GLOBAL_DB.Model(&comment).Delete("id = ?", commentID).Commit()
}