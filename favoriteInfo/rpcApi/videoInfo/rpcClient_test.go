package videoInfo

import (
	"bytes"
	"encoding/json"
	"favoriteInfo/rpcApi/videoInfo/api"
	"fmt"
	"testing"
)

func TestGetVideoInfo(t *testing.T) {
	InitVideoInfoRpcClient()

	video, _ := GetVideoInfo(0, 31)

	printVideo(video)
}

func TestGetVideoInfoList(t *testing.T) {
	var videoIDList = []int64{32, 33, 34, 50, 40}
	videos, _ := GetVideoInfoList(0, videoIDList)

	for _, video := range videos {
		printVideo(video)
	}
}

func printVideo(video *api.Video) {

	bs, _ := json.Marshal(video)
	var out bytes.Buffer
	json.Indent(&out, bs, "", "\t")
	fmt.Printf("video=%v\n", out.String())
}
