package handlers

import (
	"fmt"
	"gateway/rpcApi/favoriteInfo"
	favoriteInfoApi "gateway/rpcApi/favoriteInfo/api"
	"gateway/tools/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func InitFavoriteRpcClient() {
	favoriteInfo.InitFavoriteModelRpcClient()
}

func FavoriteAction(ginContext *gin.Context) {
	//网关统一鉴权
	Token := ginContext.Query("token")
	userId := jwt.ParseToken(Token)
	if userId == -1 {
		str := "Token验证失败，请重新登录"
		ginContext.JSON(http.StatusOK, &favoriteInfoApi.FavoriteInfoFavoriteActionResponse{
			StatusCode: -1,
			StatusMsg:  &str,
		})
		return
	}
	videoID := str2int64(ginContext.Query("video_id"))
	actionType := str2int32(ginContext.Query("action_type"))

	resp, err := favoriteInfo.FavoriteAction(userId, videoID, actionType)
	if err != nil {
		errStr := "Favorite Action接口 RPC调用失败"
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

func FavoriteList(ginContext *gin.Context) {
	//网关统一鉴权
	Token := ginContext.Query("token")
	userID := jwt.ParseToken(Token)
	if userID == -1 {
		str := "Token验证失败，请重新登录"
		ginContext.JSON(http.StatusOK, &favoriteInfoApi.FavoriteInfoQueryFavoriteListResponse{
			StatusCode: -1,
			StatusMsg:  &str,
		})
		return
	}

	searchID := str2int64(ginContext.Query("user_id"))
	log.Println("网关 sid：", searchID)
	resp, err := favoriteInfo.QueryFavoriteList(userID, searchID)
	if err != nil {
		errStr := "Favorite List接口 RPC调用失败"
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

func str2int64(str string) int64 {
	num, _ := strconv.ParseInt(str, 10, 64)
	return num
}

func str2int32(str string) int32 {
	tmp64, _ := strconv.ParseInt(str, 10, 32)
	num := int32(tmp64)
	return num
}
