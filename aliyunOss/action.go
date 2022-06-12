package aliyunOss

import (
	//"Minio/aliyunOss"
	"fmt"
	"log"
	"net/http"

	//"github.com/RaymondCode/simple-demo/aliyunOss"
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/gin-gonic/gin"
)

//投稿接口
func Action(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	file, err := c.FormFile("data")
	if err != nil {
		log.Println(err)
	}
	oss := NewOSS("")
	err = oss.UploadFile(token, title, file)
	//url,_ := oss.GetFileURL(fileName)


	fmt.Println(filename)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, controller.Response{
		StatusCode: 0,
		StatusMsg:  " uploaded successfully",
	})
}
