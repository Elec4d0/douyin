package client

import (
	"comment/server/protos/kitex_gen/api"
	"comment/server/protos/kitex_gen/api/commentserver"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func InitCountClient() commentserver.Client {
	countClient, err := commentserver.NewClient("InitCountClient", client.WithHostPorts("localhost:8888"))
	if err != nil {
		log.Fatal(err)
	}
	return countClient
}

func CommentCount(c *gin.Context) {
	countClient := InitCountClient()
	video_id, _ := strconv.Atoi(c.Query("videoID"))
	videoID := int64(video_id)
	req := &api.DouyinCommentserverCommentcountRequest{
		VideoId: videoID,
	}
	resp, err := countClient.CommentCount(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, resp)
}
