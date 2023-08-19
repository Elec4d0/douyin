package client

import (
	"comment/services/protos/kitex_gen/api"
	"comment/services/protos/kitex_gen/api/commentserver"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func InitListClient() commentserver.Client {
	listClient, err := commentserver.NewClient("InitListClient", client.WithHostPorts("localhost:8888"))
	if err != nil {
		log.Fatal(err)
	}
	return listClient
}
func CommentList(c *gin.Context) {
	listClient := InitListClient()
	videoId, _ := strconv.Atoi(c.Query("video_id"))
	req := &api.DouyinCommentListRequest{
		VideoId: int64(videoId),
	}
	resq, err := listClient.CommentList(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, resq)
}
