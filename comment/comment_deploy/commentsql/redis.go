package commentsql

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
)

func RedisCommentAllSet(videoId int64, commentList []*Comment) {
	RedisInit()
	defer RedisClose()

	Info, err := json.Marshal(commentList)
	if err != nil {
		log.Fatal(err)
	}
	key := int642string(videoId)
	_, err = rds.Do("SET", key, Info, "NX", "EX", 60)
	if err != nil {
		log.Fatal(err)
	}
}

func RedisCommentAllGet(videoId int64) ([]*Comment, error) {
	RedisInit()
	defer RedisClose()

	key := int642string(videoId)

	Info, err := redis.Bytes(rds.Do("GET", key))
	if err != nil {
		return nil, err
	}
	commentInfo := new([]Comment)
	err = json.Unmarshal(Info, &commentInfo)
	if err != nil {
		return nil, err
	}
	var comments []*Comment
	for _, comment := range *commentInfo {
		comments = append(comments, ToMysqlComment(comment))
	}
	return comments, err
}

func ToMysqlComment(comment Comment) *Comment {
	return &Comment{
		Video_id:    comment.Video_id,
		Id:          comment.Id,
		Content:     comment.Content,
		User_id:     comment.User_id,
		Create_date: comment.Create_date,
	}
}

func int642string(num int64) string {
	str := strconv.FormatInt(num, 10)
	return str
}
