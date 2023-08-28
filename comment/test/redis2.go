package main

import (
	"comment/comment_deploy/commentsql"
	"log"
)

func test8() {
	commentsql.RedisCommentCountSet(4, 5)
	count, err := commentsql.RedisCommentCountGet(4)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(count)

}
