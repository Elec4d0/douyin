package query

import (
	"log"
	"videoInfo/core/kitex_gen/api"
	"videoInfo/rpcApi"
	"videoInfo/tools/redis"
	"videoInfo/tools/videoconv"
)

func MixQueryVideoList(userID int64, videoIDList []int64, isVideoListCache []bool) (videoList []*api.Video, ok bool) {
	//通过是否缓存信息，构造查询chache与 查询Model层的videoIDList
	var cacheVideoIDList, modelVideoIDList []int64
	for i, isCache := range isVideoListCache {
		if isCache {
			cacheVideoIDList = append(cacheVideoIDList, videoIDList[i])
		} else {
			modelVideoIDList = append(modelVideoIDList, videoIDList[i])
		}
	}

	//这里做异步并发
	//查redis
	cacheVideoList, ok := RedisQueryVideoList(userID, cacheVideoIDList)
	//查model层
	modelVideoList := RpcQueryVideoList(userID, modelVideoIDList)

	if !ok {
		return nil, false
	}
	videoList = make([]*api.Video, len(videoIDList))
	modelIdx := 0
	cacheIdx := 0
	for i, videoID := range videoIDList {
		if cacheIdx < len(cacheVideoList) && videoID == cacheVideoIDList[cacheIdx] {
			//来源于缓存
			videoList[i] = cacheVideoList[cacheIdx]
			cacheIdx++
		} else if modelIdx < len(modelVideoList) && videoID == modelVideoIDList[modelIdx] {
			//来源于Model层
			videoList[i] = modelVideoList[modelIdx]
			modelIdx++
		} else {
			//异常情况，建议改写map以增强健壮性
			log.Println("双指针错误")
			continue
		}
	}
	return
}

func CacheQueryVideoList(userID int64, videoIDList []int64) ([]*api.Video, bool) {
	return RedisQueryVideoList(userID, videoIDList)
}

func ModelQueryVideoList(userID int64, videoIDList []int64) (videoList []*api.Video) {
	return RpcQueryVideoList(userID, videoIDList)
}

func RedisQueryVideo(userID, videoID int64) (apiVideo *api.Video, ok bool) {
	if redis.CheckVideoExists(videoID) {
		//这里写异步并发
		redisVideo := redis.QueryVideo(videoID)
		rpcUser := rpcApi.GetUserById(userID, redisVideo.AuthorID)
		isFavorite := rpcApi.GetIsFavorite(userID, videoID)

		apiVideo = videoconv.Redis2Api(redisVideo, isFavorite, rpcUser)
		return apiVideo, true
	}

	return nil, false
}

func RedisQueryVideoList(userID int64, videoIDList []int64) (videoList []*api.Video, ok bool) {
	videoList = make([]*api.Video, len(videoIDList))
	for i, videoID := range videoIDList {
		videoList[i], ok = RedisQueryVideo(userID, videoID)
		if !ok {
			return nil, false
		}
	}
	return
}

func RpcQueryVideo(userID, videoID int64) *api.Video {
	//这里做异步并发
	rpcVideoModel := rpcApi.QueryVideo(videoID)
	favoriteCount := rpcApi.GetFavoriteCount(videoID)
	commentCount := rpcApi.GetCommentCount(videoID)

	//这里做异步并发
	rpcUser := rpcApi.GetUserById(userID, rpcVideoModel.AuthorId)
	isFavorite := rpcApi.GetIsFavorite(userID, videoID)

	video := videoconv.Rpc2Api(rpcVideoModel, rpcUser, favoriteCount, commentCount, isFavorite)

	//CacheVideoToRedis，这里做defer
	defer redis.CacheVideo(videoconv.Api2Redis(video))
	return video
}

func RpcQueryVideoList(userID int64, videoIDList []int64) (videoList []*api.Video) {
	//这里做异步并发
	//获取Video基本模型
	rpcVideoModel := rpcApi.QueryVideoList(videoIDList)
	//获取用户与视频的喜好关系
	isFavoriteList := rpcApi.GetIsFavoriteList(userID, videoIDList)
	//获取视频点赞数
	FavoriteCountList := rpcApi.GetFavouriteCountList(videoIDList)
	//获取视频评论数
	CommentCountList := rpcApi.GetCommentCountList(videoIDList)

	//非异步项，异步后同步项：获取video作者UserInfo
	var authorIDList []int64
	for _, video := range rpcVideoModel {
		authorIDList = append(authorIDList, video.AuthorId)
	}
	authorList := rpcApi.GetAuthorList(userID, authorIDList)

	videoList = videoconv.BatchRpc2Api(rpcVideoModel, authorList, FavoriteCountList, CommentCountList, isFavoriteList)

	//结构体转化并缓存
	defer redis.CacheVideoList(videoconv.BatchApi2Redis(videoList))
	return
}
