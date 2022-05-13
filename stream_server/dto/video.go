package dto

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Title        string `gorm:"title"`
	HashID       uint   `gorm:"hash_id"`
	Hash         Hash
	AuthorID     uint   `gorm:"author_id"`
	Introduction string `gorm:"introduction"`
	Public       bool   `gorm:"public"`
	Comments     []Comment
}

type Comment struct {
	gorm.Model
	Message  string `gorm:"message"`
	VideoID  uint   `gorm:"video_id"`
	AuthorID uint   `gorm:"author_id"`
	UPNum    uint   `gorm:"up_num;default:0"`
}

type Hash struct {
	gorm.Model
	Path    string `gorm:"path"`
	HashSum string `gorm:"hash_sum;type:text"`
	Active  bool   `gorm:"active;default:true"`
}
