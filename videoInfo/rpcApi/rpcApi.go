package rpcApi

import (
	comment "videoInfo/rpcApi/comment"
	"videoInfo/rpcApi/favoriteModel"
	userInfo "videoInfo/rpcApi/userInfoAPI"
	userInfoApi "videoInfo/rpcApi/userInfoAPI/api"
	videoModel "videoInfo/rpcApi/videoModel"
	videoModelApi "videoInfo/rpcApi/videoModel/api"
)

func InitRpcClient() {
	videoModel.InitVideoModelRpcClient()
	userInfo.InitUserInfoRpcClient()
	favoriteModel.InitFavoriteModelRpcClient()
	comment.InitCommentInfoRpcClient()
}

func QueryVideoList(videoIDs []int64) []*videoModelApi.VideoModel {
	videoList, err := videoModel.QueryVideoList(videoIDs)
	if err != nil {
		n := len(videoIDs)
		return make([]*videoModelApi.VideoModel, n, n)
	}
	return videoList
}

func QueryVideo(videoID int64) *videoModelApi.VideoModel {
	video, _ := videoModel.QueryVideo(videoID)
	return video
}

func QueryFeedVideoIDList(nextTime int64, limit int64) (videoIDList []int64, createTimeList []int64) {
	videoIDList, createTimeList, _ = videoModel.QueryVideoFeed(nextTime, limit)
	return videoIDList, createTimeList
}

func QueryAuthorVideoIDList(AuthorId int64) []int64 {
	videoIDList, _ := videoModel.QueryAuthorVideoIDList(AuthorId)
	return videoIDList
}

func GetUserById(uid int64, aid int64) *userInfoApi.FullUser {
	userInfo, _ := userInfo.GetFullUserInfo(uid, aid)
	if userInfo == nil {
		return nil
	}
	return userInfo
}

func GetFavoriteCount(videoID int64) int64 {
	count, _ := favoriteModel.QueryVideoFavoriteCount(videoID)
	return count
}

func GetCommentCount(videoID int64) int64 {
	commentCount, _ := comment.GetCommentCount(videoID)
	return commentCount
}

func GetIsFavorite(userID int64, videoID int64) bool {
	isFavorite, _ := favoriteModel.QueryIsUserFavorite(userID, videoID)
	return isFavorite
}

func GetCommentCountList(videoIDList []int64) []int64 {
	commentCountList, _ := comment.GetCommentAllCount(videoIDList)
	return commentCountList
}

func GetFavouriteCountList(videoIDList []int64) []int64 {
	favoriteCountList, _ := favoriteModel.BatchQueryVideoFavoriteCount(videoIDList)
	return favoriteCountList
}

func GetIsFavoriteList(userID int64, videoIDList []int64) []bool {
	isFavoriteList, _ := favoriteModel.BatchQueryIsUserFavorite(userID, videoIDList)
	return isFavoriteList
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
