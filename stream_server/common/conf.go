package common

import "os"

var (
	MySqlUser     string
	MysqlPassword string
	MysqlUrl      string
	DatabaseName  string
)

func init() {
	MySqlUser = GetEnv("MYSQL_USER", "root")
	MysqlPassword = GetEnv("MYSQL_PASSWD", "123456")
	MysqlUrl = GetEnv("MYSQL_URL", "127.0.0.1:3306")
	DatabaseName = GetEnv("DATABASE_NAME", "video")
}

func GetEnv(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	} else {
		return val
	}
}
