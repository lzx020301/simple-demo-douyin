package main

import (
	"simple/controller"

	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine){
	r.Static("/static" ,"./message")

	apiRouter := r.Group("/douyin")

	apiRouter.GET("/feed/" ,controller.Feed)
}