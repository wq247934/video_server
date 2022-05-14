package dao

import (
	"log"
	"video_server/stream_server/common"
)

func init() {
	// 迁移 schema
	if err := common.MySqlDB.AutoMigrate(
		&Video{},
		&Hash{},
		&Comment{},
		&User{},
	); err != nil {
		panic(err)
	}
	log.Println("Migrate DB successfully!")
}
