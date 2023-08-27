package main

import (
	"gateway/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	r := InitRoute()
	r.Run()

}

func InitRoute() *gin.Engine {

	//gin引擎
	ginRouter := gin.Default()

	//初始化微服务RPC客户端链接
	handlers.InitVideoRpcClient()
	handlers.InitUseruserRpcClient()
	handlers.InitCommentClient()

	//抖音路由组
	v1 := ginRouter.Group("/douyin")
	{
		v1.GET("/feed/", handlers.Feed)
		v1.POST("/publish/action/", handlers.Action)
		v1.GET("/publish/list/", handlers.List)

		user := v1.Group("/user")
		{
			user.GET("/", handlers.User)
			user.POST("/register/", handlers.Register)
			user.POST("/login/", handlers.Login)
		}

		comment := v1.Group("/comment")
		{
			comment.POST("/action/", handlers.CommentAction)
			comment.GET("/list/", handlers.CommentList)
		}
	}
	return ginRouter
}
