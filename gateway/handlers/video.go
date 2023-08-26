package handlers

import (
	"bytes"
	"fmt"
	"gateway/rpcApi/videoInfo"
	videoInfoApi "gateway/rpcApi/videoInfo/api"
	"gateway/rpcApi/videoPublish"
	videoPublishApi "gateway/rpcApi/videoPublish/api"
	"gateway/tools/jwt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func InitVideoRpcClient() {
	videoInfo.InitVideoInfoRpcClient()
	videoPublish.InitVideoPublishRpcClient()
}

func Action(ginContext *gin.Context) {
	//网关统一鉴权
	Token := ginContext.PostForm("token")
	userId := jwt.ParseToken(Token)
	if userId == -1 {
		str := "Token验证失败，请重新登录"
		ginContext.JSON(http.StatusOK, &videoPublishApi.VideoPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  &str,
		})
		return
	}
	fileHeader, _ := ginContext.FormFile("data")
	file, _ := fileHeader.Open()
	defer file.Close()

	buffer := bytes.NewBuffer(nil)
	_, err := io.Copy(buffer, file)
	if err != nil {
		errStr := "网关读取file对象过程中，buffer拷贝错误"
		log.Println(errStr)
		ginContext.JSON(http.StatusOK, &videoPublishApi.VideoPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  &errStr,
		})
		return
	}

	resp, err := videoPublish.PublishVideo(userId, buffer.Bytes(), ginContext.PostForm("title"))

	if err != nil {
		errStr := "Publish Action接口 RPC调用失败"
		fmt.Println(errStr)
		ginContext.JSON(http.StatusOK, &videoPublishApi.VideoPublishActionResponse{
			StatusCode: -1,
			StatusMsg:  &errStr,
		})
		return
	}

	ginContext.JSON(http.StatusOK, resp)

}

func List(ginContext *gin.Context) {
	Token := ginContext.Query("token")
	userID := jwt.ParseToken(Token)
	if userID == -1 {
		str := "Token验证失败，请重新登录"
		ginContext.JSON(http.StatusOK, &videoInfoApi.VideoInfoGetAuthorVideoInfoListResponse{
			StatusCode: -1,
			StatusMsg:  &str,
			VideoList:  nil,
		})
		return
	}

	authorID, err := strconv.ParseInt(ginContext.Query("user_id"), 10, 64)
	if err != nil {
		errStr := "Publish List接口 User_id 字符串解析int64失败"
		fmt.Println(errStr)
		ginContext.JSON(http.StatusOK, &videoInfoApi.VideoInfoGetAuthorVideoInfoListResponse{
			StatusCode: -1,
			StatusMsg:  &errStr,
			VideoList:  nil,
		})
		return
	}

	resp, err := videoInfo.GetAuthorVideoInfoList(authorID, userID)
	if err != nil {
		errStr := "Publish List接口 RPC调用失败"
		fmt.Println(errStr)
		ginContext.JSON(http.StatusOK, &videoInfoApi.VideoInfoGetAuthorVideoInfoListResponse{
			StatusCode: -2,
			StatusMsg:  &errStr,
			VideoList:  nil,
		})
		return
	}

	fmt.Println(resp)
	ginContext.JSON(http.StatusOK, resp)
	return
}

func Feed(ginContext *gin.Context) {
	//	httpReq
	token := ginContext.Query("token")
	userId := jwt.ParseToken(token)

	var latestTimeStr = ginContext.Query("latest_time")
	log.Println("网关获取前端提供的查询时间：", latestTimeStr)
	var lastTime int64
	if latestTimeStr == "" {
		lastTime = time.Now().Unix() * 1000
	} else {
		parseTime, err := strconv.ParseInt(latestTimeStr, 10, 64)
		if err != nil {
			fmt.Println("Feed接口 latestTime 字符串解析int64失败")
			lastTime = -1
		}
		lastTime = parseTime
	}
	log.Println("网关传递给微服务的查询时间：", lastTime)
	resp, err := videoInfo.GetFeed(userId, lastTime)

	if err != nil {
		errStr := "GetFeed Rpc接口调用失败"
		log.Println(errStr)
		ginContext.JSON(http.StatusOK, &videoInfoApi.VideoInfoGetFeedResponse{
			StatusCode: -2,
			StatusMsg:  &errStr,
			VideoList:  nil,
		})
	}

	//fmt.Println(resp)
	ginContext.JSON(http.StatusOK, resp)
}
