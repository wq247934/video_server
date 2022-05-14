package dao

import (
	"fmt"
	"gorm.io/gorm"
	"video_server/stream_server/common"
)

type User struct {
	gorm.Model
	Username string `gorm:"username;unique;not null"`
	Nickname string `gorm:"nickname"`
	Salt     string `gorm:"salt"`
	Password string `gorm:"password"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) LoginCheck() *User {
	user := &User{Username: u.Username}
	userInfo := user.Find()
	saltPassword := common.GenSaltPassword(userInfo.Salt, u.Password)
	fmt.Println(saltPassword)
	if saltPassword == userInfo.Password {
		return userInfo
	}
	return nil
}

func (u *User) Find() *User {
	out := &User{}
	common.MySqlDB.Where(u).Find(out)
	return out
}

func (u *User) Save() error {
	return common.MySqlDB.Create(u).Error
}
