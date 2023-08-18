package main

import (
	"comment/client"
	"github.com/gin-gonic/gin"
)

func main() {

	ginRoute := gin.Default()

	ginRoute.POST("/douyin/comment/action/", client.CommentAction)
	ginRoute.GET("/douyin/comment/list/", client.CommentList)
	ginRoute.GET("/douyin/commentserver/CommentCount/:videoID", client.CommentCount)
	ginRoute.POST("/douyin/commentserver/CommentAllCount/", client.CommentAllCount)

	ginRoute.Run(":8089")
}
