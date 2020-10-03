package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
)

type envList struct {
	RabbitUrl   string
	RabbitUser  string
	RabbitPass  string
	RabbitQName string
	MysqlHost   string
	MysqlPort   string
	MysqlUser   string
	MysqlPass   string
	MysqlDBName string
}

var env = envList{
	RabbitUrl:   "localhost:5672/",
	RabbitUser:  "admin",
	RabbitPass:  "password",
	RabbitQName: "dogOrCatQ",
	MysqlHost:   "localhost",
	MysqlPort:   "3306",
	MysqlUser:   "admin",
	MysqlPass:   "password",
	MysqlDBName: "dogorcat",
}

/*
var env = envList{
	RabbitUrl:   os.Getenv("RABBIT_URL"),
	RabbitUser:  os.Getenv("RABBIT_USER"),
	RabbitPass:  os.Getenv("RABBIT_PASSWORD"),
	RabbitQName: os.Getenv("RABBIT_QUEUE_NAME"),
	MysqlHost:   os.Getenv("MYSQL_HOST"),
	MysqlPort:   os.Getenv("MYSQL_PORT"),
	MysqlUser:   os.Getenv("MYSQL_USER"),
	MysqlPass:   os.Getenv("MYSQL_PASSWORD"),
	MysqlDBName: os.Getenv("MYSQL_DBNAME"),
}
*/

func main() {
	subscribeMessage()
}

func subscribeMessage() {

	conn, err := amqp.Dial("amqp://" + env.RabbitUser + ":" + env.RabbitPass + "@" + env.RabbitUrl)
	if err != nil {
		log.Fatalf("Failed connect to RabbitMQ: %s", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed open to Channel: %s", err.Error())
		os.Exit(1)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		env.RabbitQName,
		false,
		false,
		false,
		false,
		nil,
	)

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed subscribe message: %s", err.Error())
		os.Exit(1)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			voteWorker(string(d.Body))
		}
	}()

	log.Printf("Wating for message. To exit press CTRL + C")
	<-forever
}

func voteWorker(body string) {
	conn, err := getConnectionMysql()
	if err != nil {
		return
	}
	defer conn.Close()

	_, err = votingData(body, conn)
	if err != nil {
		log.Fatalf("Failed voting %s", err.Error())
		return
	}
	log.Printf("vote a message: %s", body)
}

func getConnectionMysql() (conn *sql.DB, err error) {

	conn, err = sql.Open("mysql", env.MysqlUser+":"+env.MysqlPass+"@tcp("+env.MysqlHost+":"+env.MysqlPort+")/"+env.MysqlDBName)
	if err != nil {
		log.Println("error: mysql get connetion error.", err.Error())
	}
	return conn, err
}

func votingData(body string, conn *sql.DB) (int64, error) {
	data, err := transactAndReturnData(conn, func(tx *sql.Tx) (interface{}, error) {

		sql := "INSERT INTO VOTEBOX (CREATE_DATE,VOTE) VALUES(CURRENT_TIMESTAMP,?)"
		stmt, err := tx.Prepare(sql)
		if err != nil {
			return nil, err
		}
		result, err := stmt.Exec(body)
		if err != nil {
			return nil, err
		}
		count, _ := result.RowsAffected()

		return count, err
	})
	if err != nil {
		return 0, err
	}
	return data.(int64), nil
}

func transactAndReturnData(conn *sql.DB, txFunc func(*sql.Tx) (interface{}, error)) (data interface{}, err error) {
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
