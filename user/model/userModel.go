package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID              int64  `gorm:"primary_key"`
	Name            string `gorm:"not null"`
	Password        string `gorm:"not null"`
	FollowCount     int64  `gorm:"default:(0)"`
	FollowerCount   int64  `gorm:"default:(0)"`
	Avatar          string `gorm:"default:(-)"`
	BackgroundImage string `gorm:"default:(-)"`
	Signature       string `gorm:"default:(-)"`
	WorkCount       int64  `gorm:"default:(0)"`
	TotalFavorited  int64  `gorm:"default:(0)"`
	FavoriteCount   int64  `gorm:"default:(0)"`
	CreatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	UpdatedAt       time.Time
	Token           string
}

func CreateUser(user *User) (int64, error) {
	result := DB.Create(&user)
	return user.ID, result.Error
}

func FindName(name string) error {
	user := User{Name: name}
	result := DB.Where("name = ?", name).First(&user)
	return result.Error
}

func FindUserByName(name string) (User, error) {
	user := User{Name: name}
	result := DB.Where("name = ?", name).First(&user)
	return user, result.Error
}

func FindUserByID(id int64) (User, error) {
	user := User{ID: id}
	result := DB.Where("id = ?", id).First(&user)
	return user, result.Error
}
