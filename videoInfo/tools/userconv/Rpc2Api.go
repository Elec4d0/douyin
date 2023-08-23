package userconv

import (
	"videoInfo/core/kitex_gen/api"
	userInfoApi "videoInfo/rpcApi/userInfoAPI/api"
)

func Rpc2Api(rpcUser *userInfoApi.FullUser) *api.User {
	if rpcUser == nil {
		return nil
	}
	return &api.User{
		Id:              rpcUser.Id,
		Name:            rpcUser.Name,
		FollowCount:     rpcUser.FollowCount,
		FollowerCount:   rpcUser.FollowerCount,
		IsFollow:        rpcUser.IsFollow,
		Avatar:          rpcUser.Avatar,
		BackgroundImage: rpcUser.BackgroundImage,
		Signature:       rpcUser.Signature,
		TotalFavorited:  rpcUser.TotalFavorited,
		WorkCount:       rpcUser.WorkCount,
		FavoriteCount:   rpcUser.FavoriteCount,
	}
}

func BatchRpc2Api(rpcUserList []*userInfoApi.FullUser) (apiUserList []*api.User) {
	apiUserList = make([]*api.User, len(rpcUserList))

	for i, rpcUser := range rpcUserList {
		apiUserList[i] = Rpc2Api(rpcUser)
	}
	return
}
