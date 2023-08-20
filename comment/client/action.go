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

func InitActionClient() commentserver.Client {
	actionClient, err := commentserver.NewClient("InitActionClient", client.WithHostPorts("localhost:8888"))
	if err != nil {
		log.Fatal(err)
	}
	return actionClient
}
func CommentAction(c *gin.Context) {
	actionClient := InitActionClient()
	videoId, _ := strconv.Atoi(c.PostForm("video_id")) //Query获取信息，strconv.Atoi-->string转为int
	actionType, _ := strconv.Atoi(c.PostForm("action_type"))
	if actionType == 1 {
		commentText := c.PostForm("content")
		req := &api.DouyinCommentActionRequest{
			ActionType:  int32(actionType),
			VideoId:     int64(videoId),
			CommentText: &commentText,
		}
		resp, err := actionClient.CommentAction(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, resp)
	} else if actionType == 2 {
		comment_id, _ := strconv.Atoi(c.PostForm("comment_id"))
		commentId := int64(comment_id)
		req := &api.DouyinCommentActionRequest{
			ActionType: int32(actionType),
			VideoId:    int64(videoId),
			CommentId:  &commentId,
		}
		resp, err := actionClient.CommentAction(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, resp)
	}
}
