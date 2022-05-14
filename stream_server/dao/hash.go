package dao

import "gorm.io/gorm"

type Hash struct {
	gorm.Model
	Path    string `gorm:"path"`
	HashSum string `gorm:"hash_sum;type:text"`
	Active  bool   `gorm:"active;default:true"`
}

func (h *Hash) TableName() string {
	return "hashes"
}
