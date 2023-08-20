package comment_mysql

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Video_id    int64  `gorm:"not null"`
	Id          int64  `gorm:"not null"`
	Content     string `gorm:"not null"`
	User_id     int64  `gorm:"not null"`
	Create_date string
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
	var commentList []*Comment
	err := DB.Where("video_id = ?", videoId).Find(&commentList).Error
	if err != nil {
		return nil, err
	}
	return commentList, nil
}

func FindCommentCount(videoId int64) (CommentCount int64) {
	var commentCount int64
	DB.Model(&Comment{}).Where("video_id = ?", videoId).Count(&commentCount)
	return commentCount
}

func FindCommentAllCount(videoIds []int64) (CommentCounts []int64) {
	commentCounts := make([]int64, len(videoIds))
	for num, videoId := range videoIds {
		var count int64
		DB.Model(&Comment{}).Where("video_id = ?", videoId).Count(&count)
		commentCounts[num] = count
	}
	return commentCounts
}
