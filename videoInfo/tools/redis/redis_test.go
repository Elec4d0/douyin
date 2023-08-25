package redis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreateVideoCache(t *testing.T) {
	InitRedis()
	playUrl := "http://douyin.g324.asia:9000/video/c4486d7a-66ac-4b6c-9ba8-af3eb9a80181.mp4"
	CoverUrl := "http://douyin.g324.asia:9000/jepg/c4486d7a-66ac-4b6c-9ba8-af3eb9a80181.jpg"
	Title := "电动自行车"
	video := &Video{
		VideoID:       90854,
		AuthorID:      1000004,
		PlayUrl:       &playUrl,
		CoverUrl:      &CoverUrl,
		FavoriteCount: 10788000,
		CommentCount:  3457,
		Title:         &Title,
	}
	CreateVideoObjectCache(video)
}

func TestQueryVideoCache(t *testing.T) {
	var videoID int64 = 90854
	var video *Video
	for i := 0; i < 1; i++ {
		video = QueryVideoObjectCache(videoID)
	}
	fmt.Println(video)
	fmt.Println(video.VideoID)
	printVideo(*video)
}

func TestCheckVideoEists(t *testing.T) {
	fmt.Println("————————————开始测试——————TestCheckVideoEists————————")
	var videoID int64 = 908354
	res := CheckVideoExists(videoID)
	fmt.Println(res)
	fmt.Println("————————————结束测试——————TestCheckVideoEists————————")
}

func printVideo(video Video) {
	bs, _ := json.Marshal(video)
	var out bytes.Buffer
	json.Indent(&out, bs, "", "\t")
	fmt.Printf("video=%v\n", out.String())
}
