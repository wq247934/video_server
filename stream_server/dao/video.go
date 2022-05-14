package dao

import (
	"gorm.io/gorm"
	"video_server/stream_server/common"
)

type Video struct {
	gorm.Model
	Title        string `gorm:"title"`
	HashID       uint   `gorm:"hash_id"`
	Hash         Hash
	User         User   `gorm:"foreignKey:AuthorID"`
	AuthorID     uint   `gorm:"author_id"`
	Introduction string `gorm:"introduction"`
	Public       bool   `gorm:"public"`
	Comments     []Comment
}

func (v *Video) TableName() string {
	return "videos"
}

func (v *Video) Find() *Video {
	out := &Video{}
	common.MySqlDB.Where(v).Find(out)
	return out
}

func (v *Video) Save() error {
	return common.MySqlDB.Save(v).Error

}
