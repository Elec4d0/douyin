package user_info

import (
	"comment/server/protos/kitex_gen/api"
	"comment/server/protos/kitex_gen/api/commentserver"
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"time"
)

type User struct {
	Id               int64
	Name             string
	Follow_count     int64
	Follower_count   int64
	Is_follow        bool
	Avatar           string
	Background_image string
	Signature        string
	Total_favorited  int64
	Work_count       int64
	Favorite_count   int64
}

func UserInfo(userId int64) (user User) {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	cli, err := commentserver.NewClient("CommentServer", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	req := &api.DouyinCommentserverCommentallcountRequest{}
	resp, err := cli.CommentAllCount(ctx, req)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	return resp
}
