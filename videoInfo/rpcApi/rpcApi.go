package rpcApi

import (
	"videoInfo/core/kitex_gen/api"
	userInfo "videoInfo/rpcApi/userInfoAPI"
	userInfoApi "videoInfo/rpcApi/userInfoAPI/api"
	videoModel "videoInfo/rpcApi/videoModel"
	videoModelApi "videoInfo/rpcApi/videoModel/api"
)

func InitRpcClient() {
	videoModel.InitVideoModelRpcClient()
	userInfo.InitUserInfoRpcClient()
}

func QueryVideoList(videoIDs []int64) []*videoModelApi.VideoBaseInfo {
	videoList, err := videoModel.QueryVideoList(videoIDs)
	if err != nil {
		n := len(videoIDs)
		return make([]*videoModelApi.VideoBaseInfo, n, n)
	}
	return videoList
}

func QueryVideo(videoID int64) *videoModelApi.VideoBaseInfo {
	video, _ := videoModel.QueryVideo(videoID)
	return video
}

func QueryVideoFeed(nextTime int64) []*videoModelApi.VideoBaseInfo {
	videos, _ := videoModel.QueryVideoFeed(nextTime)
	return videos
}

func QueryAuthorVideoList(AuthorId int64) []*videoModelApi.VideoBaseInfo {
	videos, _ := videoModel.QueryAuthorVideoList(AuthorId)
	return videos
}

func GetUserById(uid int64, aid int64) *userInfoApi.FullUser {
	userInfo, _ := userInfo.GetFullUserInfo(uid, aid)
	if userInfo == nil {
		return nil
	}
	return userInfo
}

func GetFavoriteCount(videoID int64) int64 {
	return 22
}
func GetCommentCount(videoID int64) int64 {
	return 33
}
func GetIsFavortite(userID int64, authorID int64) bool {
	return true
}

func GetCommentCountList(videoIDList []int64) []int64 {
	n := len(videoIDList)
	List := make([]int64, n, n)
	return List
}

func GetFavouriteCountList(videoIDList []int64) []int64 {
	n := len(videoIDList)
	List := make([]int64, n, n)
	return List
}

func GetIsFavoriteList(userID int64, videoIDList []int64) []bool {
	n := len(videoIDList)
	List := make([]bool, n, n)
	return List
}

func RpcUserList2ApiUserList(rpcUserList []*userInfoApi.FullUser) []*api.User {
	if rpcUserList == nil {
		return nil
	}

	n := len(rpcUserList)
	var ApiUserList = make([]*api.User, n, n)

	for i, rpcUser := range rpcUserList {
		ApiUserList[i] = RpcUser2ApiUser(rpcUser)
	}
	return ApiUserList
}

func RpcUser2ApiUser(userInfo *userInfoApi.FullUser) *api.User {
	if userInfo == nil {
		return nil
	}

	user := &api.User{
		Id:              userInfo.Id,
		Name:            userInfo.Name,
		FollowCount:     userInfo.FollowCount,
		FollowerCount:   userInfo.FollowerCount,
		IsFollow:        userInfo.IsFollow,
		Avatar:          userInfo.Avatar,
		BackgroundImage: userInfo.BackgroundImage,
		Signature:       userInfo.Signature,
		TotalFavorited:  userInfo.TotalFavorited,
		WorkCount:       userInfo.WorkCount,
		FavoriteCount:   userInfo.FavoriteCount,
	}
	return user
}

func GetAuthorList(userID int64, authorIDList []int64) []*userInfoApi.FullUser {
	n := len(authorIDList)
	rpcAuthorList, err := userInfo.GetFullUserInfoList(userID, authorIDList)
	if err != nil {
		return make([]*userInfoApi.FullUser, n, n)
	}
	//防止从其他rpc拿到author 数组为 nil
	if rpcAuthorList == nil {
		return make([]*userInfoApi.FullUser, n, n)
	}

	return rpcAuthorList
}

func BuildVideoList(videoBaseInfoList []*videoModelApi.VideoBaseInfo, AuthorList []*api.User, isFavoriteList []bool, favoriteCountList []int64, commentCountList []int64) []*api.Video {
	var videoList []*api.Video

	for i, videoBaseInfo := range videoBaseInfoList {
		var video *api.Video
		//如果基本信息没查到，那么视频也就播放不了，其他信息就没有意义，直接返回nil
		if videoBaseInfo != nil {
			video = &api.Video{
				Id:            videoBaseInfo.VideoId,
				Author:        AuthorList[i],
				PlayUrl:       videoBaseInfo.PlayUrl,
				CoverUrl:      videoBaseInfo.CoverUrl,
				Title:         videoBaseInfo.Title,
				FavoriteCount: favoriteCountList[i],
				CommentCount:  commentCountList[i],
				IsFavorite:    isFavoriteList[i],
			}
		} else {
			video = nil
		}

		videoList = append(videoList, video)
	}
	return videoList
}
