package query

import (
	"log"
	"videoInfo/core/kitex_gen/api"
	"videoInfo/rpcApi"
	userApi "videoInfo/rpcApi/userInfoAPI/api"
	"videoInfo/tools/redis"
	"videoInfo/tools/videoconv"
)

/*
func CacheQueryAuthorVideoIDList() {

}

func ModelQueryAuthorVideoIDList() {

}

*/

func CacheQueryAuthorVideoList(userID int64, author *userApi.FullUser, videoIDList []int64) (videoList []*api.Video, ok bool) {
	log.Println("Cache层查询的videoID：", videoIDList)
	return RedisQueryAuthorVideoList(userID, author, videoIDList)
}

func ModelQueryAuthorVideoList(userID int64, author *userApi.FullUser, videoIDList []int64) []*api.Video {
	log.Println("Model层查询的videoID：", videoIDList)
	return RpcQueryAuthorVideoList(userID, author, videoIDList)
}

func MixQueryAuthorVideoList(userID int64, author *userApi.FullUser, videoIDList []int64, isVideoListCache []bool) (videoList []*api.Video, ok bool) {
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
	cacheVideoList, ok := CacheQueryAuthorVideoList(userID, author, cacheVideoIDList)
	//查model层
	modelVideoList := ModelQueryAuthorVideoList(userID, author, modelVideoIDList)

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

func RedisQueryAuthorVideoList(userID int64, author *userApi.FullUser, videoIDList []int64) (videoList []*api.Video, ok bool) {
	redisVideoList, ok := redis.QueryVideoList(videoIDList)
	if !ok {
		return nil, false
	}
	//获取用户与视频的喜好关系
	isFavoriteList := rpcApi.GetIsFavoriteList(userID, videoIDList)
	videoList = videoconv.BuildBatchApiWithAuthor(redisVideoList, isFavoriteList, author)
	return videoList, ok
}

func RpcQueryAuthorVideoList(userID int64, author *userApi.FullUser, videoIDList []int64) []*api.Video {
	//获取视频基本模型
	videoModelList := rpcApi.QueryVideoList(videoIDList)

	//获取用户与视频的喜好关系
	isFavoriteList := rpcApi.GetIsFavoriteList(userID, videoIDList)
	//获取视频点赞数
	FavoriteCountList := rpcApi.GetFavouriteCountList(videoIDList)
	//获取视频评论数
	CommentCountList := rpcApi.GetCommentCountList(videoIDList)

	videoList := videoconv.BatchWithAuthorRpc2Api(videoModelList, author, FavoriteCountList, CommentCountList, isFavoriteList)
	CacheVideoList(videoList)
	return videoList
}
