package handlers

import (
	"context"
	"fmt"
	"gateway/microService/user/api"
	"gateway/microService/user/api/userservice"
	userInfo "gateway/rpcApi/userInfoAPI"
	userInfoApi "gateway/rpcApi/userInfoAPI/api"
	"gateway/tools/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var userRpcClient userservice.Client

func InitUseruserRpcClient() {
	userInfo.InitUserInfoRpcClient()
}

func User(ginContext *gin.Context) {
	//———————网关统一鉴权————————
	Token := ginContext.Query("token")
	userID := jwt.ParseToken(Token)
	if userID == -1 {
		str := "Token验证失败，请重新登录"
		ginContext.JSON(http.StatusOK, &userInfoApi.DouyinUserGetFullUserInfoResponse{
			StatusCode: -1,
			StatusMsg:  &str,
			FullUser:   nil,
		})
	}

	searchID, err := strconv.ParseInt(ginContext.Query("user_id"), 10, 64)
	if err != nil {
		errStr := "网关层解析userId失败"
		log.Println(errStr)

		ginContext.JSON(http.StatusOK, &userInfoApi.DouyinUserGetFullUserInfoResponse{
			StatusCode: -1,
			StatusMsg:  &errStr,
			FullUser:   nil,
		})
	}

	resp, err := userInfo.GetFullUserInfo(userID, searchID)
	if err != nil {
		errStr := "UserInfo微服务 FullUser接口 RPC调用失败"
		log.Println(errStr)
		ginContext.JSON(http.StatusOK, &userInfoApi.DouyinUserGetFullUserInfoResponse{
			StatusCode: -1,
			StatusMsg:  &errStr,
			FullUser:   nil,
		})
	}
	ginContext.JSON(http.StatusOK, resp)

}

func Login(ginContext *gin.Context) {
	rpcReq := &api.DouyinUserLoginRequest{
		Username: ginContext.Query("username"),
		Password: ginContext.Query("password"),
	}

	fmt.Println(rpcReq)
	resp, err := userRpcClient.UserLogin(context.Background(), rpcReq)
	if err != nil {
		errStr := "User Login接口 RPC调用失败"
		fmt.Println(errStr)
	}
	fmt.Println(resp)
	ginContext.JSON(http.StatusOK, resp)
	return
}

func Register(ginContext *gin.Context) {
	rpcReq := &api.DouyinUserRegisterRequest{
		Username: ginContext.Query("username"),
		Password: ginContext.Query("password"),
	}

	resp, err := userRpcClient.UserRegister(context.Background(), rpcReq)
	if err != nil {
		errStr := "User Register接口 RPC调用失败"
		fmt.Println(errStr)
	}
	ginContext.JSON(http.StatusOK, resp)
	return
}
