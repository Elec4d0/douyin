package model

import (
	"fmt"
	"testing"
	"time"
)

/*
func TestCreateVideo(t *testing.T) {
	tVideo := &Video{
		AuthorID:      1,
		PlayUrl:       "http://bilibili.tv",
		CoverUrl:      "http://bilibili.tv",
		Title:         "测试标题",
		CreatedTime:   time.Now(),
	}
	Init()
	fmt.Println(tVideo)
	CreateVideo(tVideo)
	fmt.Printf("创建新视频数据后，数据库返回的自增video_id：%d \n", tVideo.VideoID)
	//fmt.Printf("插入成功， %d \n", &video.VideoID)
}
*/

func TestQuerySingleVideo(t *testing.T) {
	Init()
	fmt.Println("开始测试——————TestQuerySingleVideo——————————————————")
	var videoId uint64 = 3
	video, err := QuerySingleVideo(videoId)
	if err != nil {
		fmt.Println("查询失败")
	} else {
		fmt.Println("查询成功，video对象结果：")
		fmt.Println(video)
	}
	fmt.Println("结束测试——————TestQuerySingleVideo——————————————————")
}

func TestGetVideosByIds(t *testing.T) {
	fmt.Println("开始测试——————GetVideosByIds——————————————————")
	//var videoIds []uint64 := {1, 2, 3, 4}
	videoIds := []uint64{1, 2, 3, 4}
	videos, err := GetVideosByIds(videoIds)

	if err != nil {
		fmt.Println("查询错误")
	} else {
		for _, video := range videos {
			fmt.Println(video)
		}
	}
	fmt.Println("结束测试——————GetVideosByIds——————————————————")
}

func TestQueryVideoFeedByLastTimeAndLimit(t *testing.T) {
	fmt.Println("开始测试——————TestQueryVideoFeedByLastTimeAndLimit——————————————————")
	tt := time.Now()
	format := "2006-01-02 15:04:05"
	searchTime := tt.Format(format)
	fmt.Println(searchTime)

	ormVideos, err := QueryVideoFeedByLastTimeAndLimit(&searchTime, 30)

	fmt.Println(ormVideos)
	fmt.Println(err)

	for _, video := range ormVideos {
		fmt.Println(video)
	}

	fmt.Println("结束测试——————TestQueryVideoFeedByLastTimeAndLimit—————————————————")
}

func TestQueryAuthorWorkCount(t *testing.T) {
	fmt.Println("开始测试——————TestQueryAuthorWorkCount———————————————")

	var authorID uint64 = 1
	count, _ := QueryAuthorWorkCount(authorID)

	fmt.Println(count)

	fmt.Println("结束测试—————TestQueryAuthorWorkCount——————————————")
}
