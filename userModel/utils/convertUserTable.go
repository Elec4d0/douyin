package utils

import (
	"userModel/model"
	"userModel/services/protos/kitex_gen/api"
)

func ConvertUserTableToBaseUser(user model.User) *api.BaseUser {
	return &api.BaseUser{
		Id:              user.ID,
		Name:            user.Name,
		Avatar:          &user.Avatar,
		BackgroundImage: &user.BackgroundImage,
		Signature:       &user.Signature,
	}
}
