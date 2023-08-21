package main

import (
	"comment/server/protos/kitex_gen/api"
	"comment/server/protos/kitex_gen/api/commentserver"
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"time"
)

func test5() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	cli, err := commentserver.NewClient("CommentServer", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	var ids []int64 = []int64{1, 2}
	req := &api.DouyinCommentserverCommentallcountRequest{
		VideoIds: ids,
	}
	resp, err := cli.CommentAllCount(ctx, req)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	time.Sleep(time.Second)
}
