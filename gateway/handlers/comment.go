package handlers

import (
	"context"
	commentInfo "gateway/rpcApi/commentInfoAPI"
	"gateway/rpcApi/commentInfoAPI/api"
	"gateway/rpcApi/commentInfoAPI/api/commentserver"
	"gateway/tools/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var commentClient commentserver.Client

func InitCommentClient() {
	commentInfo.InitCommentInfoRpcClient()
}

func Authority(Token string) (bool, int64) {
	//———————网关统一鉴权————————
	userID := jwt.ParseToken(Token)
	if userID == -1 {
		return false, 0
	}
	return true, userID
}

func CommentAction(ginContext *gin.Context) {
	token := ginContext.PostForm("token")
	bl, userId := Authority(token)
	if bl == false {
		str := "Token验证失败，请重新登录"
		ginContext.JSON(http.StatusOK, &api.DouyinCommentActionResponse{
			StatusCode: -1,
			StatusMsg:  &str,
		})
		return
	}
	videoId, _ := strconv.Atoi(ginContext.PostForm("video_id"))
	actionType, _ := strconv.Atoi(ginContext.PostForm("action_type"))
	if actionType == 1 {
		commentText := ginContext.PostForm("content")
		req := &api.DouyinCommentActionRequest{
			UserId:      userId,
			ActionType:  int32(actionType),
			VideoId:     int64(videoId),
			CommentText: &commentText,
		}
		resp, err := commentClient.CommentAction(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		ginContext.JSON(http.StatusOK, resp)
	} else if actionType == 2 {
		comment_id, _ := strconv.Atoi(ginContext.PostForm("comment_id"))
		commentId := int64(comment_id)
		req := &api.DouyinCommentActionRequest{
			UserId:     userId,
			ActionType: int32(actionType),
			VideoId:    int64(videoId),
			CommentId:  &commentId,
		}
		resp, err := commentClient.CommentAction(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		ginContext.JSON(http.StatusOK, resp)
	}
}

func CommentList(ginContext *gin.Context) {
	token := ginContext.Query("token")
	if token == "" { //用户是否登录
		videoId, _ := strconv.Atoi(ginContext.Query("video_id"))
		req := &api.DouyinCommentListRequest{
			UserId:  -1,
			VideoId: int64(videoId),
		}
		resp, err := commentClient.CommentList(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		ginContext.JSON(http.StatusOK, resp)
	} else {
		bl, userId := Authority(token)
		if bl == false {
			str := "Token验证失败，请重新登录"
			ginContext.JSON(http.StatusOK, &api.DouyinCommentListResponse{
				StatusCode: -1,
				StatusMsg:  &str,
			})
			return
		}
		videoId, _ := strconv.Atoi(ginContext.Query("video_id"))
		req := &api.DouyinCommentListRequest{
			UserId:  userId,
			VideoId: int64(videoId),
		}
		resp, err := commentClient.CommentList(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		ginContext.JSON(http.StatusOK, resp)
	}

}
