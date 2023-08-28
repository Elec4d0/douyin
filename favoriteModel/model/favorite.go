package model

import (
	"errors"
	"log"
	"time"
)

type Favorite struct {
	ID          int64      `gorm:"primary_key"`
	UserID      int64      `gorm:"not null"`
	VideoID     int64      `gorm:"not null"`
	AuthorID    int64      `gorm:"not null"`
	CreatedTime *time.Time `gorm:"not null"`
	UpdatedTime *time.Time `gorm:"not null"`
}

func CreateLikeVideo(userID, videoID, authorID int64) error {
	curTime := time.Now()
	favorite, err := QueryFavorite(userID, videoID)
	if favorite != nil {
		return errors.New("已点赞")
	}
	favorite = &Favorite{
		UserID:      userID,
		VideoID:     videoID,
		AuthorID:    authorID,
		CreatedTime: &curTime,
		UpdatedTime: &curTime,
	}

	user, _ := QueryUser(userID)
	if user == nil {
		user = &User{
			UserID:         userID,
			LikeCount:      0,
			FavoritedCount: 0,
			CreatedTime:    &curTime,
			UpdatedTime:    &curTime,
		}
		DB.Create(&user)
	}
	user.LikeCount += 1

	video, _ := QueryVideo(videoID)
	if video == nil {
		video = &Video{
			VideoID:       videoID,
			FavoriteCount: 0,
			CreatedTime:   &curTime,
			UpdatedTime:   &curTime,
		}
		DB.Create(&video)
	}
	video.FavoriteCount += 1

	//写入更新
	err = DB.Create(&favorite).Error
	err1 := DB.Save(&user).Error
	err2 := DB.Save(&video).Error
	if err != nil || err1 != nil || err2 != nil {
		return errors.New("写入更新失败，请即使清理更新video表和user表数据")
	}
	return nil
}

func RemoveLikeVideo(userID, videoID, authorID int64) error {
	favorite, err := QueryFavorite(userID, videoID)
	if err != nil || favorite == nil {
		return errors.New("取消点赞失败，未查到此前的点赞信息")
	}

	user, err := QueryUser(userID)
	if err != nil || user == nil {
		return errors.New("取消点赞失败，未查到此前User的点赞数量信息")
	}

	video, _ := QueryVideo(videoID)
	if err != nil || video == nil {
		return errors.New("取消点赞失败，未查到此前Video的点赞数量信息")
	}

	//三表查询通过，即可开始更新:
	curTime := time.Now()
	//User更新点赞数量、更新时间
	user.LikeCount -= 1
	user.UpdatedTime = &curTime
	//Video更新点赞数量、更新时间
	video.FavoriteCount -= 1
	video.UpdatedTime = &curTime

	//写入更新
	err = DB.Delete(&Favorite{}, favorite.ID).Error
	err1 := DB.Save(&user).Error
	err2 := DB.Save(&video).Error
	if err != nil || err1 != nil || err2 != nil {
		return errors.New("写入更新失败，请即使清理更新video表和user表数据")
	}
	return nil
}

func QueryUserFavoriteList(userID int64) ([]int64, error) {
	log.Println("gorm 拿到的userID", userID)
	var favoriteList []*Favorite
	err := DB.Where("user_id = ?", userID).Find(&favoriteList).Error
	if err != nil {
		return nil, err
	}
	log.Println("gorm favoriteList :", favoriteList)
	var videoIDList = make([]int64, len(favoriteList))
	for i, favorite := range favoriteList {
		videoIDList[i] = favorite.VideoID
	}
	return videoIDList, nil
}

func QueryFavorite(userID, videoID int64) (*Favorite, error) {
	favorite := new(Favorite)
	err := DB.Where("user_id = ?", userID).Where("video_id = ?", videoID).First(&favorite).Error
	if err != nil {
		return nil, err
	}
	return favorite, nil
}

func QueryIsUserFavorite(userID, videoID int64) (bool, error) {
	favorite := new(Favorite)
	err := DB.Where("user_id = ?", userID).Where("video_id = ?", videoID).First(&favorite).Error
	if err != nil || favorite == nil {
		return false, err
	}

	return true, nil
}

func BatchQueryIsUserFavorite(userID int64, videoIDList []int64) ([]bool, error) {
	var favoriteSet []*Favorite
	err := DB.Where("user_id = ?", userID).Where("video_id in (?)", videoIDList).Find(&favoriteSet).Error
	if err != nil || favoriteSet == nil {
		return nil, err
	}

	//favoriteSet 送入哈希表做映射
	m := make(map[int64]*Favorite)
	for _, favorite := range favoriteSet {
		m[favorite.VideoID] = favorite
	}

	//利用哈希表，查询结果有序返回
	var IsUserFavoriteList = make([]bool, len(videoIDList))
	for i, videoID := range videoIDList {
		if _, ok := m[videoID]; ok {
			IsUserFavoriteList[i] = true
		} else {
			IsUserFavoriteList[i] = false
		}
	}

	return IsUserFavoriteList, nil
}
