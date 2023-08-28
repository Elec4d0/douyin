package commentsql

import (
	"gorm.io/gorm"
	"log"
)

type Comment struct {
	Id          int64  `gorm:"primary_key"`
	Video_id    int64  `gorm:"not null"`
	Content     string `gorm:"not null"`
	User_id     int64  `gorm:"not null"`
	Create_date string `gorm:"not null"`
}

type CommentCount struct {
	Video_id int64 `gorm:"primary_key"`
	Count    int64 `gorm:"not null"`
}

func CreatComment(comment *Comment) error {
	result := DB.Create(&comment)
	return result.Error
}

func DeleteComment(comment *Comment) error {
	result := DB.Delete(&comment)
	return result.Error
}

func FindComment(videoId int64, commentId int64) (*Comment, error) {
	comment := new(Comment)
	err := DB.First(&comment, "video_id = ? AND id = ?", videoId, commentId).Error
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func FindCommentAll(videoId int64) ([]*Comment, error) {
	Info, err := RedisCommentAllGet(videoId)
	if err != nil {
		var commentList []*Comment
		err := DB.Where("video_id = ?", videoId).Find(&commentList).Error
		if err != nil {
			return nil, err
		}
		RedisCommentAllSet(videoId, commentList)
		return commentList, nil
	}
	log.Println(Info)
	return Info, nil
}

func FindCommentCount(videoId int64) (int64, error) {
	var commentCount *CommentCount
	err := DB.First(&commentCount, "Video_id=?", videoId).Error
	if err != nil || commentCount == nil {
		return 0, err
	}
	return commentCount.Count, nil
}

func FindCommentAllCount(videoIds []int64) ([]int64, error) {
	var CommentCountSet []*CommentCount
	err := DB.Where("video_id in (?)", videoIds).Find(&CommentCountSet).Error
	if err != nil {
		return nil, err
	}

	//mysql查询结果为无序集合，送入hashmap
	mp := make(map[int64]int64)
	for _, commentCount := range CommentCountSet {
		mp[commentCount.Video_id] = commentCount.Count
	}

	//利用哈希表，按原videoID顺序返回
	var commentCountList = make([]int64, len(videoIds))
	for i, videoId := range videoIds {
		if commentCount, ok := mp[videoId]; ok {
			commentCountList[i] = commentCount
			log.Println(videoId, commentCount)
		} else {
			commentCountList[i] = 0
		}
	}
	return commentCountList, nil
}

func CommentCountAdd(videoId int64) {
	var commentcount CommentCount
	if err := DB.FirstOrCreate(&commentcount, CommentCount{Video_id: videoId}).Error; err != nil {
		log.Fatal(err)
		return
	}
	log.Println(commentcount)
	DB.Model(&commentcount).UpdateColumn("count", gorm.Expr("count + ?", 1))
}

func CommentCountDel(videoId int64) {
	var commentcount CommentCount
	if err := DB.Where("video_id=?", videoId).First(&commentcount).Error; err != nil {
		log.Fatal(err)
		return
	}
	if commentcount.Count > 0 {
		DB.Model(&commentcount).UpdateColumn("count", gorm.Expr("count - ?", 1))
	}
}
