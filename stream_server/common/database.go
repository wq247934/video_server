package common

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"video_server/stream_server/dto"
)

var (
	MySqlDB *gorm.DB
)

func init() {
	var err error
	MySqlDB, err = getDatabaseConn()
	if err != nil {
		panic(err)
	}
	// 迁移 schema

	if err := MySqlDB.AutoMigrate(
		&dto.Video{},
		&dto.Hash{},
		&dto.Comment{},
	); err != nil {
		panic(err)
	}
}

func getDatabaseConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", MySqlUser, MysqlPassword, MysqlUrl, DatabaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "open database error")
	}
	return db, nil
}
