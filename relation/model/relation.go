package model

import (
	"database/sql"
	"time"
)

type Relation struct {
	UserId          int64        `gorm:"primaryKey;autoIncrement:false"`
	FollowingUserId int64        `gorm:"primaryKey;autoIncrement:false"`
	IsFriend        sql.NullBool `gorm:"default: 'false'"`
	CreatedTime     *time.Time
	UpdatedTime     *time.Time
}

func CreateRelation(relation *Relation) error { // 关注
	userId := relation.UserId
	toUserId := relation.FollowingUserId
	result := DB.Create(&relation)
	to_relation := new(Relation)
	err := DB.First(&to_relation, "user_id = ? AND following_user_id = ?", toUserId, userId).Error
	if err == nil { // 对方已经关注当前用户
		// 将两个关系设置为朋友
		DB.Model(&Relation{UserId: userId}).Select("is_friend").Updates(map[string]interface{}{"user_id": toUserId, "following_user_id": userId, "is_friend": sql.NullBool{Valid: true}})
		DB.Model(&Relation{UserId: toUserId}).Select("is_friend").Updates(map[string]interface{}{"user_id": toUserId, "following_user_id": userId, "is_friend": sql.NullBool{Valid: true}})
	}
	return result.Error
}

func CancelFollow(userId, toUserId int64) error { // 取关
	relation := new(Relation)
	err := DB.First(&relation, "user_id = ? AND following_user_id = ?", userId, toUserId).Error
	if err != nil { // 当前关系不存在
		return err
	}
	if relation.IsFriend.Valid {
		// 状态从互相关注转为对方关注 取消两人的朋友状态
		to_relation := new(Relation)
		err := DB.First(&to_relation, "user_id = ? AND following_user_id = ?", toUserId, userId).Error
		if err != nil { // 互关的关系不存在
			result := DB.Delete(relation) // 删除当前关系
			return result.Error
		}
		DB.Model(&to_relation).Select("is_friend").Updates(map[string]interface{}{"user_id": toUserId, "following_user_id": userId, "is_friend": sql.NullBool{Valid: false}})
	}
	result := DB.Delete(relation) // 取关之后删除这段关系，即可以得出如果关系存在 那必然是关注状态
	return result.Error
}

func FindFollowCount(userId int64) (followCount int64) { // 当前用户的关注数
	DB.Model(&Relation{}).Where("user_id = ?", userId).Count(&followCount)
	return
}

func FindFansCount(toUserId int64) (fansCount int64) { // 当前用户的粉丝数
	DB.Model(&Relation{}).Where("following_user_id = ?", toUserId).Count(&fansCount)
	return
}

func FindFollowID(userId int64) (followList []int64, err error) { // 返回所有的关注ID
	var relationList []*Relation
	err = DB.Where("user_id = ?", userId).Find(&relationList).Error
	for _, r := range relationList {
		followList = append(followList, r.FollowingUserId)
	}
	return
}

func FindFansID(userId int64) (fansList []int64, err error) { // 返回所有的粉丝ID
	var relationList []*Relation
	err = DB.Where("following_user_id = ?", userId).Find(&relationList).Error
	for _, r := range relationList {
		fansList = append(fansList, r.UserId) // r.UserId 关注了userId
	}
	return
} //FindFriendID
func FindFriendID(userId int64) (friendList []int64, err error) { // 返回所有的朋友ID
	var relationList []*Relation
	err = DB.Where("user_id = ? AND is_friend = ?", userId, sql.NullBool{Valid: true}).Find(&relationList).Error
	for _, r := range relationList {
		friendList = append(friendList, r.UserId)
	}
	return
}

func FindAllFollowCount(userId []int64) []int64 { // 批量返回关注数量
	followCount := make([]int64, len(userId))
	for _, id := range userId {
		followCount = append(followCount, FindFollowCount(id))
	}
	return followCount
}

func FindAllFansCount(toUserId []int64) []int64 { // 批量返回粉丝数量
	fansCount := make([]int64, len(toUserId))
	for _, id := range toUserId {
		fansCount = append(fansCount, FindFansCount(id))
	}
	return fansCount
}

func FindRelationInfo(userId, searchId int64) (isFollow bool, searchIdFollow, searchIdFans int64) {
	// return parm
	// isFollow: userId是否关注searchId
	// searchIdFollow：searchId的关注数
	// searchIdFans：searchId的粉丝数
	relation := new(Relation)
	err := DB.First(&relation, "user_id = ? AND following_user_id = ?", userId, searchId).Error
	isFollow = err == nil
	searchIdFollow = FindFollowCount(searchId)
	searchIdFans = FindFansCount(searchId)
	return
}

func FindAllRelationInfo(userId, searchId []int64) (isFollow []bool, searchIdFollow, searchIdFans []int64) {
	for i := 0; i < len(userId); i++ {
		uId, sId := userId[i], searchId[i]
		isF, sFol, sFans := FindRelationInfo(uId, sId)
		isFollow = append(isFollow, isF)
		searchIdFollow = append(searchIdFollow, sFol)
		searchIdFans = append(searchIdFans, sFans)
	}
	return
}
