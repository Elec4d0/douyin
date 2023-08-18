package model

import (
	"time"
)

type Video struct {
	VideoID     int64  `gorm:"primary_key"`
	AuthorID    int64  `gorm:"not null"`
	PlayUrl     string `gorm:"not null"`
	CoverUrl    string `gorm:"not null"`
	Title       string `gorm:"not null"`
	CreatedTime *time.Time
	DeletedTime *time.Time
	UpdatedTime *time.Time
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

func QuerySingleVideo(videoId int64) (*Video, error) {
	video := new(Video)
	err := DB.First(&video, "video_id = ?", videoId).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}

func QueryAuthorVideoList(authorId int64) ([]*Video, error) {
	var videos []*Video
	err := DB.Where("author_id = ?", authorId).Find(&videos).Error
	if err != nil {
		return nil, err
	}

	return videos, nil

}

func GetVideosByIds(videoIDs []int64) ([]*Video, error) {
	var videos []*Video

	var DBVideosSet []*Video
	err := DB.Where("video_id in (?)", videoIDs).Find(&DBVideosSet).Error
	if err != nil {
		return nil, err
	}

	var m map[int64]*Video
	for _, video := range DBVideosSet {
		m[video.VideoID] = video
	}

	for _, ID := range videoIDs {
		if video, ok := m[ID]; ok {
			videos = append(videos, video)
		} else {
			videos = append(videos, nil)
		}
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

func QueryAuthorWorkCount(authorID int64) (int32, error) {
	//var videos []*Video
	result := DB.Where("author_id = ?", authorID)
	err := result.Error

	if err != nil {
		return 0, err
	}
	return int32(result.RowsAffected), nil

}
