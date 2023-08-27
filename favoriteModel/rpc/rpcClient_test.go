package rpc

import (
	"fmt"
	"testing"
)

func TestFavoriteAction(t *testing.T) {
	InitFavoriteModelRpcClient()
	FavoriteAction(100, 100, 100, 1)
}

func TestQueryFavoriteList(t *testing.T) {
	videoIDList, _ := QueryFavoriteList(100)
	fmt.Println(videoIDList)
}
