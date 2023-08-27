package userConv

import (
	"favoriteInfo/core/kitex_gen/api"
	videoInfo "favoriteInfo/rpcApi/videoInfo/api"
)

func Rpc2Api(user *videoInfo.User) *api.FavoriteUser {
	return &api.FavoriteUser{
		Id:              user.Id,
		Name:            user.Name,
		FavoriteCount:   user.FavoriteCount,
		FollowCount:     user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       user.WorkCount,
		FollowerCount:   user.FavoriteCount,
	}
}
