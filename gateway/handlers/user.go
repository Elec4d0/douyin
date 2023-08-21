package handlers

import (
	"fmt"
	userService "gateway/rpcApi/userAPI"
	userServiceApi "gateway/rpcApi/userAPI/api"
	userInfo "gateway/rpcApi/userInfoAPI"
	userInfoApi "gateway/rpcApi/userInfoAPI/api"
	"gateway/tools/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func InitUseruserRpcClient() {
	userInfo.InitUserInfoRpcClient()
	userService.InitUserRpcClient()

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
			User:       nil,
		})
		return
	}

	searchID, err := strconv.ParseInt(ginContext.Query("user_id"), 10, 64)
	if err != nil {
		errStr := "网关层解析userId失败"
		log.Println(errStr)

		ginContext.JSON(http.StatusOK, &userInfoApi.DouyinUserGetFullUserInfoResponse{
			StatusCode: -1,
			StatusMsg:  &errStr,
			User:       nil,
		})
		return
	}

	resp, err := userInfo.GetFullUserInfo(userID, searchID)
	if err != nil {
		errStr := "UserInfo微服务 FullUser接口 RPC调用失败"
		log.Println(errStr)
		ginContext.JSON(http.StatusOK, &userInfoApi.DouyinUserGetFullUserInfoResponse{
			StatusCode: -1,
			StatusMsg:  &errStr,
			User:       nil,
		})
		return
	}
	log.Println(resp)
	ginContext.JSON(http.StatusOK, resp)

}

func Login(ginContext *gin.Context) {
	username := ginContext.Query("username")
	password := ginContext.Query("password")

	resp, err := userService.UserLogin(username, password)
	if err != nil {
		errStr := "User Login接口 RPC调用失败"
		fmt.Println(errStr)
		ginContext.JSON(http.StatusOK, &userServiceApi.DouyinUserLoginResponse{
			StatusCode: -1,
			StatusMsg:  &errStr,
			UserId:     0,
			Token:      "",
		})
		return
	}
	fmt.Println(resp)
	ginContext.JSON(http.StatusOK, resp)
	return
}

func Register(ginContext *gin.Context) {
	userName := ginContext.Query("username")
	passWord := ginContext.Query("password")

	resp, err := userService.UserRegister(userName, passWord)
	if err != nil {
		errStr := "User Register接口 RPC调用失败"
		fmt.Println(errStr)
		ginContext.JSON(http.StatusOK, &userServiceApi.DouyinUserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  &errStr,
			UserId:     0,
			Token:      "",
		})
		return
	}
	ginContext.JSON(http.StatusOK, resp)
	return
}
