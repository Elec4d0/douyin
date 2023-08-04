package protos

import (
	"user/model"
	"user/services/protos/kitex_gen/api"
)

func convertUserTable(user model.User, isFollow bool) *api.User {
	return &api.User{
		Id:              user.ID,
		Name:            user.Name,
		FollowCount:     &user.FollowCount,
		FollowerCount:   &user.FollowerCount,
		IsFollow:        isFollow,
		Avatar:          &user.Avatar,
		BackgroundImage: &user.BackgroundImage,
		Signature:       &user.Signature,
		TotalFavorited:  &user.TotalFavorited,
		WorkCount:       &user.WorkCount,
		FavoriteCount:   &user.FavoriteCount,
	}
}
