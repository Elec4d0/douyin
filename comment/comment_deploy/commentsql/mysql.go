package commentsql

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Comment struct {
	gorm.Model
	Video_id    int64  `gorm:"not null"`
	Id          int64  `gorm:"not null"`
	Content     string `gorm:"not null"`
	User_id     int64  `gorm:"not null"`
	Create_date string
}

type CommentCount struct {
	gorm.Model
	Video_id int64 `gorm:"not null"`
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
	return Info, nil
}

func FindCommentCount(videoId int64) (int64, error) {
	Info, err := RedisCommentCountGet(videoId)
	if err != nil {
		var commentCount CommentCount
		err := DB.First(&commentCount, "Video_id=?", videoId).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return 0, fmt.Errorf("not found for videoId: %d", videoId)
			}
			return 0, err
		}
		RedisCommentCountSet(videoId, commentCount.Count)
		return commentCount.Count, nil
	}
	return Info, nil
}

func FindCommentAllCount(videoIds []int64) ([]int64, error) {
	var commentCounts []int64
	for _, videoId := range videoIds {
		Info, err := RedisCommentCountGet(videoId)
		if err != nil {
			var count CommentCount
			err := DB.First(&count, "Video_id = ?", videoId).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					commentCounts = append(commentCounts, 0)
				} else {
					return nil, err
				}
			} else {
				commentCounts = append(commentCounts, count.Count)
				RedisCommentCountSet(videoId, count.Count)
			}
		} else {
			commentCounts = append(commentCounts, Info)
		}
	}
	return commentCounts, nil
}

func CommentCountAdd(videoId int64) {
	var commentcount CommentCount
	if err := DB.FirstOrCreate(&commentcount, CommentCount{Video_id: videoId}).Error; err != nil {
		log.Fatal(err)
		return
	}
	DB.Model(&commentcount).UpdateColumn("Count", gorm.Expr("Count + ?", 1))
}

func CommentCountDel(videoId int64) {
	var commentcount CommentCount
	if err := DB.Where("Video_id=?", videoId).First(&commentcount).Error; err != nil {
		log.Fatal(err)
		return
	}
	if commentcount.Count > 0 {
		DB.Model(&commentcount).UpdateColumn("Count", gorm.Expr("Count - ?", 1))
	}
}
