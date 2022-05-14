package common

import (
	"fmt"
	"testing"
)

func TestGetDatabaseConn(t *testing.T) {
	db, err := getDatabaseConn()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db.Name())
}
