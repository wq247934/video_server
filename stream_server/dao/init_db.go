package dao

import (
	"gorm.io/gorm"
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
	initUser()
}

func initUser() {
	user := User{
		Username: "admin",
	}
	if user.Find().Username == "" {
		log.Println("init user")
		u := User{
			Model:    gorm.Model{},
			Username: "admin",
			Nickname: "admin",
			Salt:     "diisadsytw70un02gh8822a",
			Password: "2371686ba2742845c7ca628e90c84078d60bd51d52a571c7f174ecfbcddf6095", // password : 123456
		}
		err := u.Save()
		if err != nil {
			log.Println("init user failed", err)
			return
		}
		return
	}
	log.Println("user admin is exists!")
}
