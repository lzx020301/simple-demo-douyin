package controller

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var Bucket *oss.Bucket

func Ali_lianjie(){
	endpoint := "oss-cn-hangzhou.aliyuncs.com"
	accessKeyId := "LTAI5tRY94HvZCVqsFAi6T3q"
	accessKeySercet := "av6Um2M6pPPlbOs4umYZG1EltcSL78"

	client,err := oss.New(endpoint ,accessKeyId ,accessKeySercet)

	if err != nil {
		fmt.Println("error : " ,err)
	}

	bucket ,err := client.Bucket("example-bucket123")

	if err != nil {
		fmt.Println("bucket error : " ,err)
	}

	Bucket = bucket
}