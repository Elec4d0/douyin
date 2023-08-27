package model

import "time"

type Video struct {
	VideoID       int64      `gorm:"primary_key"`
	FavoriteCount int64      `gorm:"not null"`
	CreatedTime   *time.Time `gorm:"not null"`
	UpdatedTime   *time.Time `gorm:"not null"`
}

func QueryVideo(videoID int64) (*Video, error) {
	video := new(Video)
	err := DB.First(&video, "video_id = ?", videoID).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}

func QueryVideoFavoriteCount(videoID int64) (int64, error) {
	video := new(Video)
	err := DB.First(&video, "video_id = ?", videoID).Error
	if err != nil {
		return 0, err
	}
	return video.FavoriteCount, nil
}

func BatchQueryVideoFavoriteCount(videoIDList []int64) ([]int64, error) {
	var videoSet []*Video
	err := DB.Where("video_id in (?)", videoIDList).Find(videoSet).Error
	if err != nil {
		return nil, err
	}

	//videoSet 送入哈希表做映射
	m := make(map[int64]*Video)
	for _, video := range videoSet {
		m[video.VideoID] = video
	}

	//利用hash表，将查询结果按原数组的ID顺序返回
	var videoFavoriteCountList = make([]int64, len(videoIDList))
	for i, videoID := range videoIDList {
		if video, ok := m[videoID]; ok {
			videoFavoriteCountList[i] = video.FavoriteCount
		} else {
			videoFavoriteCountList[i] = 0
		}
	}

	return videoFavoriteCountList, nil
}
