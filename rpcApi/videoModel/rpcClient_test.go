package videoModel

import (
	"fmt"
	"testing"
)

func TestQueryVideo(t *testing.T) {
	InitVideoModelRpcClient()

	video, _ := QueryVideo(1)
	fmt.Println(video)
}

func TestQueryAuthorWorkCount(t *testing.T) {
	count, _ := QueryAuthorWorkCount(1000006)
	fmt.Println(count)
}
