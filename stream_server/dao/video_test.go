package dao

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestSave(t *testing.T) {
	//u := User{
	//	Model:    gorm.Model{},
	//	Username: "wangqian",
	//	Nickname: "王乾",
	//	Salt:     "dkaslgehjyy&##",
	//	Password: "kingqian999",
	//}
	//err := u.Save()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	v := Video{
		Model:        gorm.Model{},
		Title:        "视频1",
		HashID:       1,
		AuthorID:     1,
		Introduction: "视频简介",
		Public:       false,
		Comments:     nil,
	}

	if err := v.Save(); err != nil {
		log.Println(err)
		return
	}
}

func TestFind(t *testing.T) {
	v := Video{
		Model: gorm.Model{ID: 3},
	}
	fmt.Println(v.Title)
}
