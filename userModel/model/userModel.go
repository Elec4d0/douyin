package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID              int64  `gorm:"primary_key"`
	Name            string `gorm:"not null"`
	Password        string `gorm:"not null"`
	Avatar          string `gorm:"default:(-)"`
	BackgroundImage string `gorm:"default:(-)"`
	Signature       string `gorm:"default:(-)"`
	CreatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	UpdatedAt       time.Time
}

func CreateUser(user *User) (int64, error) {
	result := DB.Create(&user)
	return user.ID, result.Error
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

func FindUserByIDs(ids []int64) ([]*User, error) {
	var users []*User
	result := DB.Where("id = ?", ids).First(&users)
	return users, result.Error
}
