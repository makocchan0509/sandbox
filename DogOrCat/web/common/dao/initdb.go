package dao

import (
	"fmt"
	"log"
)

func SetUpDB() {

	DbConnection, err := GetConnectionMysql()
	defer DbConnection.Close()
	if err != nil {
		log.Fatalln(err)
	}
	cmd := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			CREATE_DATE TIMESTAMP NOT NULL,
            VOTE VARCHAR(10))`, "VOTEBOX")
	DbConnection.Exec(cmd)
}
