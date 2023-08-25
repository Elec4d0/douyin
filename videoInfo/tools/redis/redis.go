package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"math/rand"
	"strconv"
	"time"
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

const expirationTime = time.Minute * 5
const existTTL = time.Second * 10

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
		int642string(video.VideoID),
		"id", video.VideoID,
		"author_id", video.AuthorID,
		"play_url", *video.PlayUrl,
		"cover_url", *video.CoverUrl,
		"favorite_count", video.FavoriteCount,
		"comment_count", video.CommentCount,
		"title", *video.Title,
	)
	redisClient.Expire(ctx, int642string(video.VideoID), getExpirationTime())
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

func QueryVideoList(videoIDList []int64) ([]*Video, bool) {
	var videoList = make([]*Video, len(videoIDList))
	for i, videoID := range videoIDList {
		videoList[i] = QueryVideoObjectCache(videoID)
	}
	return videoList, true
}

func CacheFeed(videoIDList []int64, createTimeList []int64) {
	var redisZGroup = make([]redis.Z, len(videoIDList))
	for i, videoID := range videoIDList {
		redisZGroup[i] = redis.Z{
			Score:  float64(createTimeList[i]),
			Member: videoID,
		}
	}
	cnt, err := redisClient2.ZAdd(ctx, "feed", redisZGroup...).Result()
	if err != nil {
		log.Println(cnt)
	}
	redisClient2.Expire(ctx, "feed", getExpirationTime())
}

func CacheAuthorVideoIDList(authorID int64, videoIDList []int64) {
	for _, videoID := range videoIDList {
		redisClient2.RPush(ctx, int642string(authorID), videoID)
		redisClient2.Expire(ctx, int642string(authorID), getExpirationTime())
	}
}

func QueryFeedIDList(timeUnix int64, limit int64) (videoIDList []int64, createTimeList []int64) {
	opt := &redis.ZRangeBy{
		Min:    "-inf",
		Max:    int642string(timeUnix - 1),
		Count:  limit,
		Offset: 0,
	}
	zGroup, _ := redisClient2.ZRevRangeByScoreWithScores(ctx, "feed", opt).Result()
	//redis 查询从小到大输出，而业务需求为从大到小
	n := len(zGroup)
	videoIDList = make([]int64, n)
	createTimeList = make([]int64, n)
	for i, z := range zGroup {
		videoID := str2ing64(z.Member.(string))
		log.Println(videoID)
		videoIDList[i] = videoID
		createTimeList[i] = int64(z.Score)
	}
	return videoIDList, createTimeList
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
	if count == 0 {
		return false
	}
	//检查过期时间
	ttl, err := redisClient2.TTL(ctx, "feed").Result()
	if err != nil {
		return false
	}
	if ttl < existTTL {
		return false
	}
	return true
}

func CheckFeedCountEnough(timeUnix int64, limit int64) bool {
	maxTime := int642string(timeUnix - 1)
	log.Println("check Feed timeUnix", maxTime)
	count, err := redisClient2.ZCount(ctx, "feed", "-inf", maxTime).Result()
	if err != nil {
		return false
	}
	log.Println("check Feed count:", count)
	return count >= limit
}

func CheckAuthorVideoIDListExists(authorID int64) bool {
	count, err := redisClient2.Exists(ctx, int642string(authorID)).Result()
	if err != nil {
		return false
	}
	if count == 0 {
		return false
	}
	//检查过期时间
	ttl, err := redisClient2.TTL(ctx, int642string(authorID)).Result()
	if err != nil {
		return false
	}
	if ttl < existTTL {
		return false
	}
	return true
}

func CheckVideoExists(videoID int64) bool {
	isVideoExist, err := redisClient.HExists(ctx, int642string(videoID), "id").Result()
	if err != nil {
		return false
	}
	if isVideoExist == false {
		return false
	}
	//检查过期时间
	ttl, err := redisClient.TTL(ctx, int642string(videoID)).Result()
	if err != nil {
		return false
	}
	if ttl < existTTL {
		return false
	}
	return true
}

func int642string(num int64) string {
	str := strconv.FormatInt(num, 10)
	return str
}

func str2ing64(str string) int64 {
	num, _ := strconv.ParseInt(str, 10, 64)
	return num
}

func getExpirationTime() time.Duration {
	return expirationTime + time.Second*time.Duration(rand.Int31n(120))
}
