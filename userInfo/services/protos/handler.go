package protos

import (
	"context"
	"fmt"
	api "userInfo/services/protos/kitex_gen/api"
	userModelServices "userInfo/userModelAPI"
	videoModelServices "userInfo/videoModel"
)

// UserInfoServiceImpl implements the last service interface defined in the IDL.
type UserInfoServiceImpl struct{}

func GetWorkCount(user_id int64) int64 {
	count, err := videoModelServices.QueryAuthorWorkCount(user_id)
	if err != nil {

		return 0
	}
	return int64(count)
}

func GetFavoriteInfo(user_id int64) (int64, int64) {
	return 0, 0
}

func GetRelationInfo(user_id int64, search_id int64) (int64, int64, bool) {
	return 0, 0, false
}

func GetWorkCountList(user_id []int64) []int64 {
	count, _ := videoModelServices.QueryAuthorWorkCountList(user_id)
	//var workCount = make([]int64, len(user_id))
	//if count != nil {
	//	for i := 0; i < len(count); i++ {
	//		workCount = append(workCount, int64(count[i]))
	//	}
	//}
	var workCount []int64
	for i := 0; i < len(count); i++ {
		workCount = append(workCount, int64(count[i]))
	}
	return workCount
}

func GetFavoriteInfoList(user_id []int64) ([]int64, []int64) {
	return nil, nil
}

func GetRelationInfoList(user_id int64, search_id []int64) ([]int64, []int64, []bool) {
	return nil, nil, nil
}

// GetFullUserInfo implements the UserInfoServiceImpl interface.
func (s *UserInfoServiceImpl) GetFullUserInfo(ctx context.Context, req *api.DouyinUserGetFullUserInfoRequest) (resp *api.DouyinUserGetFullUserInfoResponse, err error) {
	// TODO: Your code here...
	resp = new(api.DouyinUserGetFullUserInfoResponse)

	userId := req.UserId

	searchId := req.SearchId

	baseUser, _ := userModelServices.FindBaseUserById(searchId)

	followCount, followerCount, isFollow := GetRelationInfo(userId, searchId)

	totalFavorited, favoriteCount := GetFavoriteInfo(searchId)

	//获取视频点赞数
	workCount := GetWorkCount(searchId)

	id := int64(0)
	name := ""
	var avatar *string
	avatar = nil
	var backgroundImage *string
	backgroundImage = nil
	var signature *string
	signature = nil

	if baseUser != nil {
		id = baseUser.Id
		name = baseUser.Name
		avatar = baseUser.Avatar
		backgroundImage = baseUser.BackgroundImage
		signature = baseUser.Signature
	}

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
	resp.FullUser = fullUser
	return
}

// GetFullUserInfoList implements the UserInfoServiceImpl interface.
func (s *UserInfoServiceImpl) GetFullUserInfoList(ctx context.Context, req *api.DouyinUserGetFullUserInfoListRequest) (resp *api.DouyinUserGetFullUserInfoListResponse, err error) {
	// TODO: Your code here...
	resp = new(api.DouyinUserGetFullUserInfoListResponse)

	userId := req.UserId
	fmt.Println("find list user_id:", userId)

	searchId := req.SearchId

	baseUser, _ := userModelServices.FindBaseUserList(searchId)

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
	resp.FullUser = fullUser
	return
}
