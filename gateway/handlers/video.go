package handlers

import (
	"bytes"
	"context"
	"fmt"
	"gateway/microService/feed/api"
	"gateway/microService/feed/api/feedprotobuf"
	"gateway/tools/jwt"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

var rpcClient feedprotobuf.Client

func InitVideoRpcClient() feedprotobuf.Client {
	var err error
	rpcClient, err = feedprotobuf.NewClient("video", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		fmt.Println("网关层Video 微服务初始化链接失败")
		return nil
	}
	return rpcClient
}

func Action(ginContext *gin.Context) {
	//Token := ginContext.Query("token")
	Token := ginContext.PostForm("token")
	userId := jwt.ParseToken(Token)
	if userId == -1 {
		str := "Token验证失败，请重新登录"
		ginContext.JSON(http.StatusOK, &api.DouyinPublishListResponse{
			StatusCode: -1,
			StatusMsg:  &str,
			VideoList:  nil,
		})
	}

	fileHeader, _ := ginContext.FormFile("data")
	file, _ := fileHeader.Open()
	defer file.Close()

	buffer := bytes.NewBuffer(nil)
	_, err := io.Copy(buffer, file)
	if err != nil {
		err_str := "网关读取file对象过程中，buffer拷贝错误"
		fmt.Println(err_str)
		return
	}

	rpcReq := &api.DouyinPublishActionRequest{
		UserId: userId,
		Title:  ginContext.PostForm("title"),
		Data:   buffer.Bytes(),
	}

	resp, err := rpcClient.PublishVideo(context.Background(), rpcReq)

	if err != nil {
		errStr := "Publish Action接口 RPC调用失败"
		fmt.Println(errStr)
	}

	ginContext.JSON(http.StatusOK, resp)

}

func List(ginContext *gin.Context) {
	Token := ginContext.Query("token")
	if jwt.ParseToken(Token) == -1 {
		str := "Token验证失败，请重新登录"
		ginContext.JSON(http.StatusOK, &api.DouyinPublishListResponse{
			StatusCode: -1,
			StatusMsg:  &str,
			VideoList:  nil,
		})
	}

	userId, err := strconv.ParseInt(ginContext.Query("user_id"), 10, 64)
	if err != nil {
		errStr := "Publish List接口 User_id 字符串解析int64失败"
		fmt.Println(errStr)
		ginContext.JSON(http.StatusOK, &api.DouyinPublishListResponse{
			StatusCode: -1,
			StatusMsg:  &errStr,
			VideoList:  nil,
		})
		return
	}
	rpcReq := &api.DouyinPublishListRequest{
		UserId: userId,
	}

	resp, err := rpcClient.GetAuthorVideoList(context.Background(), rpcReq)

	if err != nil {
		errStr := "Publish List接口 RPC调用失败"
		fmt.Println(errStr)
		ginContext.JSON(http.StatusOK, &api.DouyinPublishListResponse{
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
	/*
		if userId < 0 {
			str := "token 过期或不正确"
			ginContext.JSON(http.StatusOK, api.DouyinFeedResponse{
				StatusCode: -1,
				StatusMsg:  &str,
				VideoList:  nil,
				NextTime:   nil,
			})
			return
		}
	*/
	var lateestTimeStr string = ginContext.Query("latest_time")
	var lastTime int64
	if lateestTimeStr == "" {
		lastTime = -1
	} else {
		parseTime, err := strconv.ParseInt(lateestTimeStr, 10, 64)
		if err != nil {
			fmt.Println("Feed接口 lastestTime 字符串解析int64失败")
			lastTime = -1
		}
		lastTime = parseTime
	}

	rpcReq := &api.DouyinFeedRequest{
		LatestTime: &lastTime,
		UserId:     userId,
	}

	resp, err := rpcClient.GetFeed(context.Background(), rpcReq)

	if err != nil {
		fmt.Println("Rpc接口调用失败")
	}
	/*
		httpResp := &api.DouyinFeedResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
			VideoList:  resp.VideoList,
			NextTime:   resp.NextTime,
		}
	*/

	fmt.Println(resp)

	ginContext.JSON(http.StatusOK, resp)
}
