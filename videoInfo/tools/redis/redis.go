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

func CacheVideo(video *Video) {
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

func QueryVideo(videoID int64) *Video {
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
	//var videoList = make([]*Video, len(videoIDList))
	//构造管道并送入命令
	pipeline := redisClient.Pipeline()
	for _, videoID := range videoIDList {
		redisClient.HGetAll(ctx, int642string(videoID)).Result()
	}

	//执行命令
	cmders, err := pipeline.Exec(ctx)
	if err != nil {
		log.Println("管道执行错误", err)
		return nil, false
	}

	var videoList []*Video
	//批量查询执行成功，查询到系列值
	for _, cmder := range cmders {
		cmd := cmder.(*redis.MapStringStringCmd)
		//解析数据
		mp, err := cmd.Result()
		if err != nil {
			log.Println()
		}
		if len(mp) <= 1 {
			videoList = append(videoList, nil)
		} else {
			playUrl := mp["play_url"]
			coverUrl := mp["cover_url"]
			title := mp["title"]
			video := &Video{
				VideoID:       str2ing64(mp["id"]),
				AuthorID:      str2ing64(mp["author_id"]),
				PlayUrl:       &playUrl,
				CoverUrl:      &coverUrl,
				FavoriteCount: str2ing64(mp["favorite_count"]),
				CommentCount:  str2ing64(mp["comment_count"]),
				Title:         &title,
			}
			videoList = append(videoList, video)
		}
	}

	return videoList, true
}

func CacheVideoList(videoList []*Video) {
	//构造管道并送入命令
	pipeline := redisClient.Pipeline()
	for _, video := range videoList {
		pipeline.HSet(
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
		pipeline.Expire(ctx, int642string(video.VideoID), getExpirationTime())
	}

	_, err := pipeline.Exec(ctx)
	if err != nil {
		log.Println("管道执行错误", err)
	}
	return
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
	//构造命令送入管道
	pipeline := redisClient.Pipeline()
	for _, videoID := range videoIDList {
		pipeline.HExists(ctx, int642string(videoID), "id")
		pipeline.TTL(ctx, int642string(videoID))
	}

	//管道执行与执行异常
	cmders, err := pipeline.Exec(ctx)
	if err != nil {
		log.Println("管道执行错误", err)
		return nil, -1
	}

	//对查询结果进行cmd解析
	allCache := true
	noneCache := true
	isVideoListCache = make([]bool, len(videoIDList))
	for i, cmder := range cmders {
		if i%2 == 1 {
			//跳过，两条cmd 读取一次
			continue
		}
		index := i / 2
		cmd := cmder.(*redis.BoolCmd)

		isVideoExist, err := cmd.Result()
		if err != nil || isVideoExist == false {
			allCache = false
			isVideoListCache[index] = false
			continue
		}

		ttl, err := cmders[i+1].(*redis.DurationCmd).Result()
		if ttl < existTTL || err != nil {
			allCache = false
			isVideoListCache[index] = false
			continue
		}

		//上述检查均无异常
		isVideoListCache[index] = isVideoExist
		if isVideoExist {
			noneCache = false
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
