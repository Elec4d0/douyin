package handlers

import (
	"fmt"
	commentInfo "gateway/rpcApi/commentInfoAPI"
	"gateway/rpcApi/commentInfoAPI/api"
	favoriteInfoApi "gateway/rpcApi/favoriteInfo/api"
	"gateway/tools/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
	token := ginContext.Query("token")
	bl, userId := Authority(token)
	if bl == false {
		str := "Token验证失败，请重新登录"
		ginContext.JSON(http.StatusOK, &api.DouyinCommentActionResponse{
			StatusCode: -1,
			StatusMsg:  &str,
		})
	}
	videoId := str2int64(ginContext.Query("video_id"))
	actionType := str2int32(ginContext.Query("action_type"))
	commentText := ginContext.Query("comment_text")
	commentId := str2int64(ginContext.Query("comment_id"))

	resp, err := commentInfo.CommentAction(userId, videoId, actionType, commentText, commentId)

	if err != nil {
		errStr := "Comment Action接口 RPC调用失败"
		fmt.Println(errStr)
		ginContext.JSON(http.StatusOK, &favoriteInfoApi.FavoriteInfoFavoriteActionResponse{
			StatusCode: -1,
			StatusMsg:  &errStr,
		})
		return
	}

	ginContext.JSON(http.StatusOK, resp)
	return
}

func CommentList(ginContext *gin.Context) {
	token := ginContext.Query("token")
	bl, userId := Authority(token)
	if bl == false {
		str := "Token验证失败，请重新登录"
		ginContext.JSON(http.StatusOK, &api.DouyinCommentListResponse{
			StatusCode: -1,
			StatusMsg:  &str,
		})
		return
	}
	videoId := str2int64(ginContext.Query("video_id"))

	resp, err := commentInfo.CommentList(userId, videoId)
	if err != nil {
		errStr := "Comment List接口 RPC调用失败"
		fmt.Println(errStr)
		ginContext.JSON(http.StatusOK, &api.DouyinCommentListResponse{
			StatusCode: -1,
			StatusMsg:  &errStr,
		})
		return
	}

	ginContext.JSON(http.StatusOK, resp)
	return
}
