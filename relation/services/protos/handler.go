package main

import (
	"context"
	"database/sql"
	"fmt"
	"relation/model"
	"relation/rpcApi"
	api "relation/services/protos/kitex_gen/api"
	"time"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// /douyin/relation/action/ - 关系操作
// 登录用户对其他用户进行关注或取消关注。
// 接口类型
// POST
// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *api.DouyinRelationActionRequest) (resp *api.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	t := time.Now()
	user_id := req.GetUserId()
	to_user_id := req.GetToUserId()
	isFollow := req.GetActionType()
	if to_user_id < 0 {
		statusMsg := "to_user_id is illegal"
		return &api.DouyinRelationActionResponse{
			StatusCode: 1,
			StatusMsg:  &statusMsg,
		}, fmt.Errorf("to_user_id is illegal")
	}
	if isFollow != 1 && isFollow != 2 {
		statusMsg := "ActionType is illegal"
		return &api.DouyinRelationActionResponse{
			StatusCode: 1,
			StatusMsg:  &statusMsg,
		}, fmt.Errorf("ActionType is illegal")
	}
	if isFollow == 1 {
		// 插入一条关注记录
		relation := &model.Relation{
			UserId:          user_id,
			FollowingUserId: to_user_id,
			IsFriend:        sql.NullBool{Valid: false},
			CreatedTime:     &t,
			UpdatedTime:     &t,
		}
		err := model.CreateRelation(relation)
		if err != nil {
			statusMsg := "Follow unsuccessful"
			resp = &api.DouyinRelationActionResponse{
				StatusCode: 1,
				StatusMsg:  &statusMsg,
			}
			return // 关注失败
		}
	} else if isFollow == 2 { // 取关操作
		err := model.CancelFollow(user_id, to_user_id)
		if err != nil {
			statusMsg := "Cancel Follow unsuccessful"
			resp = &api.DouyinRelationActionResponse{
				StatusCode: 1,
				StatusMsg:  &statusMsg,
			}
			return // 取消失败
		}
	}
	statusMsg := "Successful"
	resp = &api.DouyinRelationActionResponse{
		StatusCode: 0,
		StatusMsg:  &statusMsg,
	}
	return // 操作成功
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *api.DouyinRelationFollowListRequest) (resp *api.DouyinRelationFollowListResponse, err error) {
	// TODO: Your code here...
	userId := req.GetUserId()
	followIdList, err := model.FindFollowID(userId)
	if len(followIdList) == 0 {
		statusMsg := "No Followers"
		resp = &api.DouyinRelationFollowListResponse{
			StatusCode: 1,
			StatusMsg:  &statusMsg,
			UserList:   nil,
		}
		return resp, nil // 没有关注别人
	}
	var userList []*api.User
	for _, follow := range followIdList {
		user := rpcApi.GetUserById(userId, follow)
		userList = append(userList, user)
	}
	statusMsg := "Successful"
	resp = &api.DouyinRelationFollowListResponse{
		StatusCode: 0,
		StatusMsg:  &statusMsg,
		UserList:   userList,
	}
	return resp, nil
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *api.DouyinRelationFollowerListRequest) (resp *api.DouyinRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	userId := req.GetUserId()
	fansIdList, err := model.FindFansID(userId)
	if len(fansIdList) == 0 {
		statusMsg := "No Fans"
		resp = &api.DouyinRelationFollowerListResponse{
			StatusCode: 1,
			StatusMsg:  &statusMsg,
			UserList:   nil,
		}
		return resp, nil // 没有别人关注
	}
	var userList []*api.User
	for _, follow := range fansIdList {
		user := rpcApi.GetUserById(userId, follow)
		userList = append(userList, user)
	}
	statusMsg := "Successful"
	resp = &api.DouyinRelationFollowerListResponse{
		StatusCode: 0,
		StatusMsg:  &statusMsg,
		UserList:   userList,
	}
	return resp, nil
}

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *api.DouyinRelationFriendListRequest) (resp *api.DouyinRelationFriendListResponse, err error) {
	// TODO: Your code here...
	userId := req.GetUserId()
	friendIdList, err := model.FindFriendID(userId)
	if len(friendIdList) == 0 {
		statusMsg := "No Fans"
		resp = &api.DouyinRelationFriendListResponse{
			StatusCode: 1,
			StatusMsg:  &statusMsg,
			UserList:   nil,
		}
		return resp, nil // 没有别人关注
	}
	var friendUserList []*api.FriendUser
	for _, follow := range friendIdList {
		user := rpcApi.GetUserById(userId, follow)
		friendUserList = append(friendUserList, ConvertFriendUser(user))
	}
	statusMsg := "Successful"
	resp = &api.DouyinRelationFriendListResponse{
		StatusCode: 0,
		StatusMsg:  &statusMsg,
		UserList:   friendUserList,
	}
	return resp, nil
}

func ConvertFriendUser(user *api.User) (friend *api.FriendUser) {
	return &api.FriendUser{
		Id:              user.Id,
		Name:            user.Name,
		FollowCount:     user.FollowerCount,
		FollowerCount:   user.FollowCount,
		IsFollow:        user.IsFollow,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
		Message:         "",
		MsgType:         0,
	}
}

// GetOneRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetOneRelation(ctx context.Context, req *api.DouyinRelationSearchRequest) (resp *api.DouyinRelationSearchResponse, err error) {
	// TODO: Your code here...
	isfollow, follows, fans := model.FindRelationInfo(req.GetUserId(), req.GetSearchId())
	resp.IsFollow = isfollow
	resp.FollowCount = follows
	resp.FansCount = fans
	return
}

// GetListRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetListRelation(ctx context.Context, req *api.DouyinRelationSearchListRequest) (resp *api.DouyinRelationSearchListResponse, err error) {
	// TODO: Your code here...
	isfollow, follows, fans := model.FindAllRelationInfo(req.GetUserIdList(), req.GetSearchIdList())
	resp.IsFollowList = isfollow
	resp.FollowCountList = follows
	resp.FansCountList = fans
	return
}
