package query

import (
	"log"
	"time"
	"videoInfo/core/kitex_gen/api"
	"videoInfo/rpcApi"
	"videoInfo/tools/redis"
)

func GetVideoInfoList(userID int64, videoIDList []int64) []*api.Video {
	//检查每一个视频的缓存情况
	isVideoListCache, cacheType := redis.CheckAllVideoCache(videoIDList)

	var queryOK bool
	var videoList []*api.Video

	if cacheType == 1 {
		//所有视频均在缓存，仅查redis
		videoList, queryOK = RedisQueryVideoList(userID, videoIDList)
	} else if cacheType == 0 {
		//部分视频缓存，混合查询
		videoList, queryOK = MixQueryVideoList(userID, videoIDList, isVideoListCache)
	}

	if queryOK {
		return videoList
	}

	//所有视频均未缓存, cacheType = -1
	//上面的两个if查询错误均回落于此
	videoList = RpcQueryVideoList(userID, videoIDList)
	return videoList
}

func GetAuthorVideoInfoList(userID int64, authorID int64) []*api.Video {
	//redis 是否缓存过用户作品列表的videoID
	var videoIDList []int64
	if redis.CheckAuthorVideoIDListExists(authorID) {
		//从redis中拿出缓存的用户作品列表，因Check过，故视频列表一定存在
		videoIDList = redis.QueryAuthorVideoIDList(authorID)
	} else {
		//从Model层拿出作者作品列表
		videoIDList = rpcApi.QueryAuthorVideoIDList(authorID)
		//Model层查询Author作品列表失败
		if videoIDList == nil {
			return nil
		}
		defer redis.CacheAuthorVideoIDList(authorID, videoIDList)
	}
	log.Println("需要查询的VideoID:", videoIDList)

	//获取作者UserInfo
	author := rpcApi.GetUserById(userID, authorID)
	//检查每一个视频的缓存情况
	isVideoListCache, cacheType := redis.CheckAllVideoCache(videoIDList)
	//log.Println(isVideoListCache)
	//log.Println(cacheType)
	var queryOK bool
	var videoList []*api.Video

	if cacheType == 1 {
		//所有视频均在缓存，仅查redis
		videoList, queryOK = CacheQueryAuthorVideoList(userID, author, videoIDList)
	} else if cacheType == 0 {
		//部分视频缓存，混合查询
		videoList, queryOK = MixQueryAuthorVideoList(userID, author, videoIDList, isVideoListCache)
	}

	if queryOK {
		return videoList
	}

	//所有视频均未缓存, cacheType = -1
	//上面的两个if查询错误均回落于此
	videoList = ModelQueryAuthorVideoList(userID, author, videoIDList)
	return videoList
}

func GetFeed(userID int64, queryMaxTime int64) (videoList []*api.Video, nextQueryTime int64) {
	log.Println("APP提供的查询时间：", queryMaxTime)
	var videoIDList []int64
	var createTimeList []int64
	var limit int64 = 2

	//Feed列表是否在缓存中, 且剩余Feed流足够
	if redis.CheckFeedExists() && redis.CheckFeedCountEnough(queryMaxTime, limit) {
		//通过redis获取Feed流ID列表
		videoIDList, createTimeList = redis.QueryFeedIDList(queryMaxTime, limit)
		if videoIDList == nil || createTimeList == nil {
			return nil, time.Now().Unix()
		}
		//Feed流此时一定有limit个， 直接查询返回即可
		log.Println("Redis后端返回的查询时间：", createTimeList[len(createTimeList)-1])
		return GetVideoInfoList(userID, videoIDList), createTimeList[len(createTimeList)-1]
	}

	//通过rpc获取Feed流ID列表
	videoIDList, createTimeList = rpcApi.QueryFeedVideoIDList(queryMaxTime, limit)

	//Feed流结束，往下刷不出视频了
	if videoIDList == nil || createTimeList == nil {
		return nil, time.Now().Unix() * 1000
	}
	//缓存
	defer redis.CacheFeed(videoIDList, createTimeList)
	//Model层正常查询返回
	log.Println("Model后端返回的查询时间：", createTimeList[len(createTimeList)-1])
	return GetVideoInfoList(userID, videoIDList), createTimeList[len(createTimeList)-1]
}
