package dao

import (
	"gorm.io/gorm"
	"video_server/stream_server/common"
)

type Hash struct {
	gorm.Model
	Path    string `gorm:"path"`
	HashSum string `gorm:"hash_sum;type:text"`
	Active  bool   `gorm:"active;default:true"`
}

func (h *Hash) TableName() string {
	return "hashes"
}

func (h *Hash) Find() *Hash {
	out := &Hash{}
	common.MySqlDB.Where(h).Find(out)
	return out
}
