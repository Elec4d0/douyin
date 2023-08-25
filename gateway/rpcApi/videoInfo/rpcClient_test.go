package videoInfo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/rpcApi/videoInfo/api"
	"testing"
)

func TestGetVideoInfo(t *testing.T) {
	InitVideoInfoRpcClient()

	video, _ := GetVideoInfo(31)

	printVideo(video)
}

func TestGetVideoInfoList(t *testing.T) {
	var videoIDList = []int64{32, 33, 34, 50, 40}
	videos, _ := GetVideoInfoList(videoIDList)

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
