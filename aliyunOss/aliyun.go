package aliyunOss

import (
	//"Minio/controller"
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/RaymondCode/simple-demo/controller"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
)

type aliyunOss struct {
	bucket *oss.Bucket
	bucketName string
	video_id int
}

type OSS interface {
	UploadFile(token string,title string,file *multipart.FileHeader) error
	GetFileURL(fileName string) (string,error)
}

var filename string

func NewOSS(BucketName string) OSS  {
	endpoint := "oss-cn-hangzhou.aliyuncs.com"
	accessKeyId := ""
	accessKeySercet := ""

	client,err := oss.New(endpoint ,accessKeyId ,accessKeySercet)

	if err != nil {
		fmt.Println("error : " ,err)
	}
	if BucketName == ""{
		BucketName = "example-bucket123"
	}


	bucket ,err := client.Bucket(BucketName)
	if err != nil {
		fmt.Println(err)
	}
	b:=&aliyunOss{bucketName: bucket.BucketName,
		bucket: bucket,
	video_id: 1}
	return b
}

func (a *aliyunOss)UploadFile(token string,title string,file *multipart.FileHeader) error  {

	var token2id controller.Token2ID
	var user controller.User
	var video controller.Video
	controller.GLOBAL_DB.Model(&controller.Token2ID{}).Where("token = ?" ,token).Find(&token2id)
	fmt.Println("userid:" ,token2id.ID)
	controller.GLOBAL_DB.Model(&controller.User{}).Where("user_id = ?" ,token2id.ID).Find(&user)
	controller.GLOBAL_DB.Model(&controller.Video{}).Last(&video)
	db, err := controller.GLOBAL_DB.DB()
	if err != nil {
	   return err
	}
	result, err := db.Query("select id from token2_id where token =  ?", token)
	if err != nil {
	   return err
	}
	result.Next()
	var userId int64
	err = result.Scan(&userId)
	name:=file.Filename
	split := strings.Split(name, ".")
	if err != nil {
	   return err
	}
	date:=time.Now().Format("20060102150405")
 
	fileName:="-"+date+"."+split[len(split)-1]
	
 
	fileReader, err := file.Open()
	if err != nil {
	   fmt.Println(err)
	   return err
	}
	a.video_id++
	err = a.bucket.PutObject(a.bucketName+"/"+fileName, fileReader)
	logrus.Info(a.bucketName+"/"+fileName)

	filename = "https://example-bucket123.oss-cn-hangzhou.aliyuncs.com/" + a.bucketName+"/"+fileName

	controller.GLOBAL_DB.Model(&controller.Video{}).Create(&controller.Video{VideoId: video.VideoId+1 ,User: user ,PlayUrl: filename ,CoverUrl: "" ,FavoriteCount: 0 ,CommentCount: 0,IsFavorite: false})
	if err != nil {
	   return err
	}
	return nil
 }

func (a *aliyunOss)GetFileURL(fileName string) (string,error)  {
	url, err := a.bucket.SignURL(a.bucketName+"/"+fileName, oss.HTTPGet, int64(time.Minute*3))

	return url,err
}