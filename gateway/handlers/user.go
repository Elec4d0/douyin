package handlers

import (
	"context"
	"fmt"
	"gateway/microService/user/api"
	"gateway/microService/user/api/userservice"
	"gateway/tools/jwt"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var userRpcClient userservice.Client

func InitUseruserRpcClient() userservice.Client {
	var err error
	userRpcClient, err = userservice.NewClient("user", client.WithHostPorts("0.0.0.0:8887"))
	if err != nil {
		fmt.Println("网关层Video 微服务初始化链接失败")
		return nil
	}
	return userRpcClient
}

func User(ginContext *gin.Context) {
	//———————网关统一鉴权————————
	Token := ginContext.Query("token")
	if jwt.ParseToken(Token) == -1 {
		str := "Token验证失败，请重新登录"
		ginContext.JSON(http.StatusOK, &api.DouyinUserResponse{
			StatusCode: -1,
			StatusMsg:  &str,
			User:       nil,
		})
	}

	userId, err := strconv.ParseInt(ginContext.Query("user_id"), 10, 64)
	if err != nil {
		fmt.Println("网关层解析userId失败")
	}

	rpcReq := &api.DouyinUserRequest{
		UserId: userId,
		Token:  Token,
	}

	resp, err := userRpcClient.UserInfo(context.Background(), rpcReq)

	if err != nil {
		errStr := "Publish Action接口 RPC调用失败"
		fmt.Println(errStr)
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
