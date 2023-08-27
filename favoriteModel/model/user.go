package model

import "time"

type User struct {
	UserID         int64      `gorm:"primary_key"`
	LikeCount      int64      `gorm:"not null"`
	FavoritedCount int64      `gorm:"not null"`
	CreatedTime    *time.Time `gorm:"not null"`
	UpdatedTime    *time.Time `gorm:"not null"`
}

func QueryUser(userID int64) (*User, error) {
	user := new(User)
	err := DB.First(&user, "user_id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func QueryUserFavoriteCount(userID int64) (likeCount, favoritedCount int64, err error) {
	user := new(User)
	err = DB.First(&user, "user_id = ?", userID).Error
	if err != nil {
		return 0, 0, err
	}
	return user.LikeCount, user.FavoritedCount, nil
}

func BatchQueryUserFavoriteCount(userIDList []int64) (likeCountList, favoritedCountList []int64, err error) {
	var userSet []*User
	err = DB.Where("user_id in (?)", userIDList).Find(userSet).Error
	if err != nil {
		return nil, nil, err
	}

	//userSet 送入哈希表做映射
	m := make(map[int64]*User)
	for _, user := range userSet {
		m[user.UserID] = user
	}

	likeCountList = make([]int64, len(userIDList))
	favoritedCountList = make([]int64, len(userIDList))
	//利用hash表，将查询结果按原数组的ID顺序返回
	for i, videoID := range userIDList {
		if user, ok := m[videoID]; ok {
			likeCountList[i] = user.LikeCount
			favoritedCountList[i] = user.FavoritedCount
		} else {
			likeCountList[i] = 0
			favoritedCountList[i] = 0
		}
	}

	return likeCountList, favoritedCountList, nil
}
