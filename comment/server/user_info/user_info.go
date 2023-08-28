package user_info

import (
	"comment/comment_deploy/commentsql"
	"comment/server/protos/kitex_gen/api"
	rpcClient "comment/userInfoAPI"
	api2 "comment/userInfoAPI/api"
	"comment/userInfoAPI/api/userinfoservice"
	"log"
)

type User struct {
	Id               int64
	Name             string
	Follow_count     int64
	Follower_count   int64
	Is_follow        bool
	Avatar           string
	Background_image string
	Signature        string
	Total_favorited  int64
	Work_count       int64
	Favorite_count   int64
}

var userInfoRpcClient userinfoservice.Client

func UserInfo(userId int64, searchId int64) (*api.User, error) {
	rpcClient.InitUserInfoRpcClient()
	userInfo, err := rpcClient.GetFullUserInfo(userId, searchId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	user := ToApiFullUser(userInfo)
	return user, nil
}
func UserInfoList(userId int64, comment []*commentsql.Comment) ([]*api.Comment, error) {
	searchId := make([]int64, len(comment))
	for i, commentInfo := range comment {
		searchId[i] = commentInfo.User_id
	}
	rpcClient.InitUserInfoRpcClient()
	userInfoList, err := rpcClient.GetFullUserInfoList(userId, searchId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var CommentList = make([]*api.Comment, len(comment))
	for i, rpcUser := range userInfoList {
		apiUser := ToApiFullUser(rpcUser)
		CommentList[i] = ToApiComment(comment[i], apiUser)
	}
	return CommentList, nil
}
func ToApiFullUser(user *api2.FullUser) *api.User {
	return &api.User{
		Id:              user.Id,
		Name:            user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
	}
}

func ToApiComment(comment *commentsql.Comment, userInfo *api.User) *api.Comment {
	return &api.Comment{
		Id:         comment.Id,
		User:       userInfo,
		Content:    comment.Content,
		CreateDate: comment.Create_date,
	}
}
