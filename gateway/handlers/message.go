package handlers

import (
	"fmt"
	messageService "gateway/rpcApi/messageAPI"
	messageServiceApi "gateway/rpcApi/messageAPI/api"
	"gateway/tools/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func InitMessageRpcClient() {
	messageService.InitMessageRpcClient()
}

func MessageAction(ginContext *gin.Context) {
	//———————网关统一鉴权————————
	Token := ginContext.Query("token")
	userID := jwt.ParseToken(Token)
	if userID == -1 {
		str := "Token验证失败，无法发送消息"
		ginContext.JSON(http.StatusOK, &messageServiceApi.DouyinRelationActionResponse{
			StatusCode: -1,
			StatusMsg:  &str,
		})
		return
	}

	toUserID, err := strconv.ParseInt(ginContext.Query("to_user_id"), 10, 64)
	if err != nil {
		errStr := "网关层解析to_userId失败"
		log.Println(errStr)

		ginContext.JSON(http.StatusOK, &messageServiceApi.DouyinRelationActionResponse{
			StatusCode: -1,
			StatusMsg:  &errStr,
		})
		return
	}

	actionType, err := strconv.ParseInt(ginContext.Query("action_type"), 10, 32)
	if err != nil || actionType != 1 {
		errStr := "action_type错误"
		log.Println(errStr)

		ginContext.JSON(http.StatusOK, &messageServiceApi.DouyinRelationActionResponse{
			StatusCode: -1,
			StatusMsg:  &errStr,
		})
		return
	}

	content := ginContext.Query("content")

	resp, err := messageService.RelationAction(userID, toUserID, int32(actionType), content)
	if err != nil {
		errStr := "Message微服务 MessageAction接口 RPC调用失败"
		log.Println(errStr)
		ginContext.JSON(http.StatusOK, &messageServiceApi.DouyinRelationActionResponse{
			StatusCode: -1,
			StatusMsg:  &errStr,
		})
		return
	}
	log.Println(resp)
	ginContext.JSON(http.StatusOK, resp)
	return
}

func MessageChat(ginContext *gin.Context) {
	Token := ginContext.Query("token")
	userID := jwt.ParseToken(Token)
	if userID == -1 {
		str := "Token验证失败，无法查看消息"
		ginContext.JSON(http.StatusOK, &messageServiceApi.DouyinMessageChatResponse{
			StatusCode:  -1,
			StatusMsg:   &str,
			MessageList: nil,
		})
		return
	}

	toUserID, err := strconv.ParseInt(ginContext.Query("to_user_id"), 10, 64)
	if err != nil {
		errStr := "网关层解析to_userId失败"
		log.Println(errStr)

		ginContext.JSON(http.StatusOK, messageServiceApi.DouyinMessageChatResponse{
			StatusCode:  -1,
			StatusMsg:   &errStr,
			MessageList: nil,
		})
		return
	}

	preMsgTime, err := strconv.ParseInt(ginContext.Query("pre_msg_time"), 10, 64)
	if err != nil {
		errStr := "网关层解析pre_msg_time失败"
		log.Println(errStr)

		ginContext.JSON(http.StatusOK, messageServiceApi.DouyinMessageChatResponse{
			StatusCode:  -1,
			StatusMsg:   &errStr,
			MessageList: nil,
		})
		return
	}

	resp, err := messageService.MessageChat(userID, toUserID, preMsgTime)
	if err != nil {
		errStr := "Message微服务 MessageChat接口 RPC调用失败"
		log.Println(errStr)
		ginContext.JSON(http.StatusOK, messageServiceApi.DouyinMessageChatResponse{
			StatusCode:  -1,
			StatusMsg:   &errStr,
			MessageList: nil,
		})
		return
	}
	fmt.Println(resp)
	ginContext.JSON(http.StatusOK, resp)
	return
}
