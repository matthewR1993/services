package models

import (
	"time"
)

type UserInfo struct {
    	Id 		uint `gorm:"primary_key"`
	Login 		string `gorm:"not null;unique"`
	Email 		string `gorm:"not null;unique"`
	PasswordHash 	[]byte `gorm:"not null"`
	CreatedAt 	time.Time `gorm:"not null"`
	IsActive 	bool
	IsVerified 	bool
}

// set table name
func (u UserInfo) TableName() string {
  return "UserInfo"
}

type UserHistory struct {
	Id 		uint `gorm:"primary_key"`
	UserId		uint
	Urls		string
	History		string
}

func (u UserHistory) TableName() string {
  return "UserHistory"
}
