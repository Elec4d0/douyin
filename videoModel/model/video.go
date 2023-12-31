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
	//IDs 送入gorm查询得到videoSet
	var VideosSet []*Video
	err := DB.Where("video_id in (?)", videoIDs).Find(&VideosSet).Error
	if err != nil {
		return nil, err
	}

	//videoSet 送入哈希表做映射
	m := make(map[int64]*Video)
	for _, video := range VideosSet {
		m[video.VideoID] = video
	}

	//利用hash表，将查询结果按原数组的ID顺序返回
	var videos []*Video
	for _, ID := range videoIDs {
		if video, ok := m[ID]; ok {
			videos = append(videos, video)
		} else {
			videos = append(videos, nil)
		}
	}

	return videos, err
}

func QueryVideoFeedByLastTimeAndLimit(lastTime *string, limit int64) ([]*Video, error) {
	//fmt.Println(*lastTime)
	var videos []*Video
	err := DB.Where("created_time < ?", *lastTime).Order("created_time desc").Limit(int(limit)).Find(&videos).Error
	//err := DB.Where("created_time < ?", *lastTime).Find(&videos).Error
	if err != nil {
		return nil, err
	}

	//fmt.Println(videos)
	return videos, nil
}

func QueryAuthorWorkCount(authorID int64) (int32, error) {
	var videos []*Video
	result := DB.Where("author_id = ?", authorID).Find(&videos)
	err := result.Error

	if err != nil {
		return 0, err
	}
	return int32(result.RowsAffected), nil

}
