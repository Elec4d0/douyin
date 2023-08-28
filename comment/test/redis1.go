package main

import (
	"comment/comment_deploy/commentsql"
	"log"
)

func test7() {
	data := "10-21"

	commentList := []*commentsql.Comment{
		{
			Video_id:    1,
			Id:          1,
			Content:     "ok",
			User_id:     1,
			Create_date: data,
		},
		{
			Video_id:    1,
			Id:          2,
			Content:     "okok",
			User_id:     2,
			Create_date: data,
		},
	}
	commentsql.RedisCommentAllSet(4, commentList)
	comments, err := commentsql.RedisCommentAllGet(4)
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range comments {
		log.Println(c.Id)
		log.Println(c.Content)
	}

}
