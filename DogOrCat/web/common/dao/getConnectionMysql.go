package dao

import (
	"database/sql"
	"log"
	"projects/DogOrCat/web/config"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnectionMysql() (conn *sql.DB, err error) {

	conn, err = sql.Open("mysql", config.Env.MysqlUser+":"+config.Env.MysqlPass+"@tcp("+config.Env.MysqlHost+":"+config.Env.MysqlPort+")/"+config.Env.MysqlDBName)
	if err != nil {
		log.Println("error: mysql get connetion error.", err.Error())
	}
	return conn, err
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
