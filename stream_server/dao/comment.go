package dao

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Message  string `gorm:"message"`
	VideoID  uint   `gorm:"video_id"`
	AuthorID uint   `gorm:"author_id"`
	UPNum    uint   `gorm:"up_num;default:0"`
}

func (c *Comment) TableName() string {
	return "comments"
}

func (c *Comment) Find() {

}

func (c *Comment) Save() {

}
