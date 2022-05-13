package common

import (
	"fmt"
	"testing"
	"video_server/stream_server/dto"
)

func TestGetDatabaseConn(t *testing.T) {
	db, err := getDatabaseConn()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db.Name())
	db.Create(&dto.Video{
		Title:        "视频标题1",
		HashID:       1,
		AuthorID:     1,
		Introduction: "视频简介,.....",
	})
}
