package model

import (
	"time"
)

type Video struct {
	VideoID       uint64    `gorm:"primary_key"`
	AuthorID      uint64    `gorm:"not null"`
	PlayUrl       string    `gorm:"not null"`
	CoverUrl      string    `gorm:"not null"`
	FavoriteCount uint64    `gorm:"not null"`
	CommentCount  uint64    `gorm:"not null"`
	Title         string    `gorm:"not null"`
	CreatedTime   time.Time `gorm:"not null"`
	DeletedTime   *time.Time
	UpdatedTime   *time.Time
}

/*
	func (v Video) String() string {
		return fmt.Sprintf("{"+
			"VideoID: %v"+
			"AuthorID: %v", v.VideoID, v.AuthorID)
	}
*/
func CreateVideo(video *Video) error {
	result := DB.Create(&video)
	return result.Error
}

func QuerySingleVideo(videoId uint64) (*Video, error) {
	video := new(Video)
	err := DB.First(&video, "video_id = ?", videoId).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}

func QueryAuthorVideoList(authorId uint64) ([]*Video, error) {
	var videos []*Video
	err := DB.Where("author_id = ?", authorId).Find(&videos).Error
	if err != nil {
		return nil, err
	}

	return videos, nil

}

func GetVideosByIds(videoIDs []uint64) ([]*Video, error) {
	var videos []*Video

	err := DB.Where("video_id in (?)", videoIDs).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, err
}

func QueryVideoFeedByLastTimeAndLimit(lastTime *string, limit int) ([]*Video, error) {
	//fmt.Println(*lastTime)
	var videos []*Video
	err := DB.Where("created_time < ?", *lastTime).Order("created_time desc").Limit(limit).Find(&videos).Error
	//err := DB.Where("created_time < ?", *lastTime).Find(&videos).Error
	if err != nil {
		return nil, err
	}

	//fmt.Println(videos)
	return videos, nil
}
