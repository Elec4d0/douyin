package rpcClient

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"strconv"
	"video/api"
	"video/api/userservice"
	video "video/services/protos/kitex_gen/api"
)

var userRpcClient userservice.Client

func InitUserRpcClient() userservice.Client {
	var err error
	userRpcClient, err = userservice.NewClient("user", client.WithHostPorts("0.0.0.0:8887"))
	if err != nil {
		fmt.Println("网关层Video 微服务初始化链接失败")
		return nil
	}
	fmt.Println("Video 微服务：初始化链接User微服务成功")
	return userRpcClient
}

func GetUserInFo(userId uint64) (*video.User, error) {
	uid, _ := strconv.ParseInt(strconv.FormatUint(userId, 10), 10, 64)
	//fmt.Println(uid)
	rpcReq := &api.DouyinUserRequest{
		UserId: uid,
	}

	resp, err := userRpcClient.UserInfo(context.Background(), rpcReq)

	if err != nil {
		errStr := "Video:::User_RPC调用失败, userId不存在"
		//fmt.Println(errStr)
		return nil, errors.New(errStr)
	}
	usr := resp.User
	return &video.User{
		Id:              usr.Id,
		Name:            usr.Name,
		FollowCount:     usr.FollowCount,
		FollowerCount:   usr.FollowerCount,
		IsFollow:        usr.IsFollow,
		Avatar:          usr.Avatar,
		BackgroundImage: usr.BackgroundImage,
		Signature:       usr.Signature,
		TotalFavorited:  usr.TotalFavorited,
		WorkCount:       usr.WorkCount,
		FavoriteCount:   usr.FavoriteCount,
	}, nil
}
