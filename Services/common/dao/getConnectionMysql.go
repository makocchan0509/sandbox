package dao

import (
	"database/sql"
	"log"
	"projects/Services/common/properties"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnectionMysql() *sql.DB {

	prop := properties.GetProp()

	dbHost := prop.Mysql.MysqlHost
	dbPort := prop.Mysql.MysqlPort
	dbUser := prop.Mysql.MysqlUser
	dbPass := prop.Mysql.MysqlPass
	dbName := prop.Mysql.MysqlDBName

	conn, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		log.Println("error: mysql get connetion error.", err.Error())
	}

	return conn
}

func CloseConnetionMysql(conn *sql.DB) {
	defer conn.Close()
}

func TransactAndReturnData(conn *sql.DB, txFunc func(*sql.Tx) (interface{}, error)) (data interface{}, err error) {
	tx, err := conn.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	data, err = txFunc(tx)
	return
}
