package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type Video struct {
	VideoID       int64   `json:"id"`
	AuthorID      int64   `json:"author_id"`
	PlayUrl       *string `json:"play_url"`
	CoverUrl      *string `json:"cover_url"`
	FavoriteCount int64   `json:"favorite_count"`
	CommentCount  int64   `json:"comment_count"`
	Title         *string `json:"title"`
}

var redisClient *redis.Client
var redisClient2 *redis.Client
var ctx context.Context

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "192.168.1.41:6379",
		Password: "",
		DB:       1,
	})

	redisClient2 = redis.NewClient(&redis.Options{
		Addr:     "192.168.1.41:6379",
		Password: "",
		DB:       2,
	})
	ctx = context.Background()
}

/*
	func CreateVideoCache(video *Video) {
		str, _ := json.Marshal(video)
		redisClient.HSet(ctx,
			strconv.FormatInt(video.VideoID, 10),
			"json", str)
	}

	func QueryVideoCache(videoID int64) *Video {
		jsonStr, _ := redisClient.HGet(ctx,
			strconv.FormatInt(videoID, 10), "json").Result()
		var video *Video
		json.Unmarshal([]byte(jsonStr), &video)
		return video
	}
*/
func CreateVideoObjectCache(video *Video) {
	redisClient.HSet(
		ctx,
		strconv.FormatInt(video.VideoID, 10),
		"id", video.VideoID,
		"author_id", video.AuthorID,
		"play_url", *video.PlayUrl,
		"cover_url", *video.CoverUrl,
		"favorite_count", video.FavoriteCount,
		"comment_count", video.CommentCount,
		"title", *video.Title,
	)
}

func QueryVideoObjectCache(videoID int64) *Video {
	mp, _ := redisClient.HGetAll(ctx,
		strconv.FormatInt(videoID, 10)).Result()
	playUrl := mp["play_url"]
	CoverUrl := mp["cover_url"]
	Title := mp["title"]

	video := &Video{
		VideoID:       str2ing64(mp["id"]),
		AuthorID:      str2ing64(mp["author_id"]),
		PlayUrl:       &playUrl,
		CoverUrl:      &CoverUrl,
		FavoriteCount: str2ing64(mp["favorite_count"]),
		CommentCount:  str2ing64(mp["comment_count"]),
		Title:         &Title,
	}
	return video
}

func CacheFeed(videoIDList []int64) {
	for _, videoID := range videoIDList {
		redisClient2.RPush(ctx, "feed", videoID)
	}
}

func CacheAuthorVideoIDList(authorID int64, videoIDList []int64) {
	for _, videoID := range videoIDList {
		redisClient2.RPush(ctx, int642string(authorID), videoID)
	}
}

func QueryFeedIDList() []int64 {
	redidRes := redisClient2.LRange(ctx, "feed", 0, -1).Val()

	feedIDList := make([]int64, len(redidRes))
	for i, str := range redidRes {
		feedIDList[i] = str2ing64(str)
	}
	return feedIDList
}

func QueryAuthorVideoIDList(authorID int64) []int64 {
	redidRes := redisClient2.LRange(ctx, int642string(authorID), 0, -1).Val()

	videoIDList := make([]int64, len(redidRes))
	for i, str := range redidRes {
		videoIDList[i] = str2ing64(str)
	}
	return videoIDList
}

func CheckAllVideoCache(videoIDList []int64) (isVideoListCache []bool, cacheType int64) {
	allCache := true
	noneCache := true
	isVideoListCache = make([]bool, len(videoIDList))
	for i, videoID := range videoIDList {
		isVideoListCache[i] = CheckVideoExists(videoID)
		if isVideoListCache[i] {
			noneCache = false
		} else {
			allCache = false
		}
	}

	//缓存情况
	if allCache {
		cacheType = 1
	} else if noneCache {
		cacheType = -1
	} else {
		cacheType = 0
	}

	return
}

func CheckFeedExists() bool {
	count, err := redisClient2.Exists(ctx, "feed").Result()
	if err != nil {
		return false
	}
	return count > 0
}

func CheckAuthorVideoIDListExists(authorID int64) bool {
	count, err := redisClient2.Exists(ctx, int642string(authorID)).Result()
	if err != nil {
		return false
	}
	return count > 0
}

func CheckVideoExists(videoID int64) bool {
	isVideoExist, err := redisClient.HExists(ctx, int642string(videoID), "id").Result()
	if err != nil {
		return false
	}
	return isVideoExist
}

func int642string(num int64) string {
	str := strconv.FormatInt(num, 10)
	return str
}

func str2ing64(str string) int64 {
	num, _ := strconv.ParseInt(str, 10, 64)
	return num
}
