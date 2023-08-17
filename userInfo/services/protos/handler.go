package protos

import (
	"context"
	api "userInfo/services/protos/kitex_gen/api"
	userModelServices "userInfo/userModelAPI"
)

// UserInfoServiceImpl implements the last service interface defined in the IDL.
type UserInfoServiceImpl struct{}

func GetWorkCount(user_id int64) int64 {
	return 0
}

func GetFavoriteInfo(user_id int64) (int64, int64) {
	return 0, 0
}

func GetRelationInfo(user_id int64, search_id int64) (int64, int64, bool) {
	return 0, 0, false
}

func GetWorkCountList(user_id []int64) []int64 {
	return nil
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

	fullUser := &api.FullUser{
		Id:              baseUser.Id,
		Name:            baseUser.Name,
		Avatar:          baseUser.Avatar,
		BackgroundImage: baseUser.BackgroundImage,
		Signature:       baseUser.Signature,
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

	searchId := req.SearchId

	baseUser, _ := userModelServices.FindBaseUserList(searchId)

	followCount, followerCount, isFollow := GetRelationInfoList(userId, searchId)

	//获取用户与视频的喜好关系
	totalFavorited, favoriteCount := GetFavoriteInfoList(searchId)

	//获取视频点赞数
	workCount := GetWorkCountList(searchId)

	var fullUser []*api.FullUser

	for i := 0; i < len(searchId); i++ {
		fullUser = append(fullUser,
			&api.FullUser{
				Id:              baseUser[i].Id,
				Name:            baseUser[i].Name,
				Avatar:          baseUser[i].Avatar,
				BackgroundImage: baseUser[i].BackgroundImage,
				Signature:       baseUser[i].Signature,
				WorkCount:       &workCount[i],
				TotalFavorited:  &totalFavorited[i],
				FavoriteCount:   &favoriteCount[i],
				FollowerCount:   &followerCount[i],
				FollowCount:     &followCount[i],
				IsFollow:        isFollow[i],
			})
	}
	resp.StatusCode = 0
	msg := "查找成功！"
	resp.StatusMsg = &msg
	resp.FullUser = fullUser
	return
}
