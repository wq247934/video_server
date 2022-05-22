package common

import (
	"os"
	"strconv"
)

var (
	MySqlUser     string
	MysqlPassword string
	MysqlUrl      string
	DatabaseName  string
	HttpAddr      string
	ReadTimeout   int
	WriteTimeout  int
	VideoPath     string
)

func init() {
	MySqlUser = GetEnv("MYSQL_USER", "root")
	MysqlPassword = GetEnv("MYSQL_PASSWD", "123456")
	MysqlUrl = GetEnv("MYSQL_URL", "127.0.0.1:3306")
	DatabaseName = GetEnv("DATABASE_NAME", "video")
	HttpAddr = GetEnv("HTTP_ADDR", "127.0.0.1:9001")
	ReadTimeout, _ = strconv.Atoi(GetEnv("READ_TIMEOUT", "30"))
	WriteTimeout, _ = strconv.Atoi(GetEnv("WriteTimeout", "30"))
	VideoPath = GetEnv("VIDEO_PATH", "/Users/wangqian/Documents/video/")
}

func GetEnv(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	} else {
		return val
	}
}
