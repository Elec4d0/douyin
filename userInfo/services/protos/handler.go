package protos

import (
	"context"
	"fmt"
	api "userInfo/services/protos/kitex_gen/api"
	"userInfo/utils/query"
)

// UserInfoServiceImpl implements the last service interface defined in the IDL.
type UserInfoServiceImpl struct{}

// GetFullUserInfo implements the UserInfoServiceImpl interface.
func (s *UserInfoServiceImpl) GetFullUserInfo(ctx context.Context, req *api.DouyinUserGetFullUserInfoRequest) (resp *api.DouyinUserGetFullUserInfoResponse, err error) {
	// TODO: Your code here...
	resp = new(api.DouyinUserGetFullUserInfoResponse)

	userId := req.UserId

	searchId := req.SearchId

	/*baseUser, _ := userModelServices.FindBaseUserById(searchId)

	if baseUser == nil {
		resp.StatusCode = 1
		msg := "查找失败！"
		resp.StatusMsg = &msg
		resp.User = nil
		return
	}

	followCount, followerCount, isFollow := GetRelationInfo(userId, searchId)

	totalFavorited, favoriteCount := GetFavoriteInfo(searchId)

	//获取视频点赞数
	workCount := GetWorkCount(searchId)

	id := baseUser.Id
	name := baseUser.Name
	avatar := baseUser.Avatar
	backgroundImage := baseUser.BackgroundImage
	signature := baseUser.Signature

	fullUser := &api.FullUser{
		Id:              id,
		Name:            name,
		Avatar:          avatar,
		BackgroundImage: backgroundImage,
		Signature:       signature,
		WorkCount:       &workCount,
		TotalFavorited:  &totalFavorited,
		FavoriteCount:   &favoriteCount,
		FollowerCount:   &followerCount,
		FollowCount:     &followCount,
		IsFollow:        isFollow,
	}
	resp.StatusCode = 0
	msg := "查找成功！"
	resp.StatusMsg = &msg
	resp.User = fullUser*/

	fullUser, err := query.QueryUserInfo(userId, searchId)
	if fullUser == nil {
		resp.StatusCode = 1
		msg := "查找失败！"
		resp.StatusMsg = &msg
		resp.User = fullUser
		return
	}
	resp.StatusCode = 0
	msg := "查找成功！"
	resp.StatusMsg = &msg
	resp.User = fullUser
	return
}

// GetFullUserInfoList implements the UserInfoServiceImpl interface.
func (s *UserInfoServiceImpl) GetFullUserInfoList(ctx context.Context, req *api.DouyinUserGetFullUserInfoListRequest) (resp *api.DouyinUserGetFullUserInfoListResponse, err error) {
	// TODO: Your code here...
	resp = new(api.DouyinUserGetFullUserInfoListResponse)

	userId := req.UserId
	fmt.Println("find list user_id:", userId)

	searchId := req.SearchId

	/*baseUser, _ := userModelServices.FindBaseUserList(searchId)

	// followCount, followerCount, isFollow := GetRelationInfoList(userId, searchId)
	followCount1 := int64(0)
	followerCount1 := int64(0)

	//获取用户与视频的喜好关系
	// totalFavorited, favoriteCount := GetFavoriteInfoList(searchId)
	totalFavorited1 := int64(0)
	favoriteCount1 := int64(0)

	//获取视频点赞数
	workCount := GetWorkCountList(searchId)

	var fullUser []*api.FullUser

	id := int64(0)
	name := ""
	var avatar *string
	avatar = nil
	var backgroundImage *string
	backgroundImage = nil
	var signature *string
	signature = nil

	for i := 0; i < len(searchId); i++ {
		if baseUser[i] != nil {
			id = baseUser[i].Id
			name = baseUser[i].Name
			avatar = baseUser[i].Avatar
			backgroundImage = baseUser[i].BackgroundImage
			signature = baseUser[i].Signature
		}
		fullUser = append(fullUser,
			&api.FullUser{
				Id:              id,
				Name:            name,
				Avatar:          avatar,
				BackgroundImage: backgroundImage,
				Signature:       signature,
				WorkCount:       &workCount[i],
				TotalFavorited:  &totalFavorited1,
				FavoriteCount:   &favoriteCount1,
				FollowerCount:   &followerCount1,
				FollowCount:     &followCount1,
				IsFollow:        false,
			})
	}
	resp.StatusCode = 0
	msg := "查找成功！"
	resp.StatusMsg = &msg
	resp.User = fullUser*/

	var userList []*api.FullUser

	userList, err = query.QueryUserListInfo(userId, searchId)
	if err != nil {
		resp.StatusCode = 1
		msg := "查找失败！"
		resp.StatusMsg = &msg
		resp.User = make([]*api.FullUser, len(searchId))
		return
	}
	resp.StatusCode = 0
	msg := "查找成功！"
	resp.StatusMsg = &msg
	resp.User = userList
	return
}
