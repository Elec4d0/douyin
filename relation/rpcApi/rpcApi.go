package rpcApi

import (
	userInfo "relation/rpcApi/userInfoAPI"
	"relation/services/protos/kitex_gen/api"
	// userInfoApi "relation/rpcApi/userInfoAPI/api"
)

func InitRpcClient() {
	userInfo.InitUserInfoRpcClient()
}

func GetUserById(uid int64, aid int64) *api.User {
	userinfo, _ := userInfo.GetFullUserInfo(uid, aid)
	if userinfo == nil {
		return nil
	}

	user := &api.User{
		Id:              userinfo.Id,
		Name:            userinfo.Name,
		FollowCount:     userinfo.FollowCount,
		FollowerCount:   userinfo.FollowerCount,
		IsFollow:        userinfo.IsFollow,
		Avatar:          userinfo.Avatar,
		BackgroundImage: userinfo.BackgroundImage,
		Signature:       userinfo.Signature,
		TotalFavorited:  userinfo.TotalFavorited,
		WorkCount:       userinfo.WorkCount,
		FavoriteCount:   userinfo.FavoriteCount,
	}
	return user
}
