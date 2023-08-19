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

type videoIDs struct {
	videoIds []string `json:"videoIds"`
}

func InitAllCountClient() commentserver.Client {
	allcountClient, err := commentserver.NewClient("InitAllCountClient", client.WithHostPorts("localhost:8888"))
	if err != nil {
		log.Fatal(err)
	}
	return allcountClient
}

func CommentAllCount(c *gin.Context) {
	allcountClient := InitAllCountClient()
	var VideoIds videoIDs
	if err := c.ShouldBindJSON(&VideoIds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var videoIDs []int64
	for num, video_id := range VideoIds.videoIds {
		videoId, _ := strconv.Atoi(video_id)
		videoID := int64(videoId)
		videoIDs[num] = videoID
	}
	req := &api.DouyinCommentserverCommentallcountRequest{
		VideoIds: videoIDs,
	}
	resp, err := allcountClient.CommentAllCount(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, resp)
}
