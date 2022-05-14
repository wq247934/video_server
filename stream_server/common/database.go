package common

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
}

func getDatabaseConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", MySqlUser, MysqlPassword, MysqlUrl, DatabaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "open database error")
	}
	return db, nil
}
