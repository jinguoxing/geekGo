package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
const (
	USER_NAME = "root"
	PASS_WORD = "111111"
	HOST      = "docker_mysql_1"
	PORT      = "3306"
	DATABASE  = "test"
	CHARSET   = "utf8"
)

var DB *sql.DB

func init() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)

	link, _ := sql.Open("mysql", dbDSN)
	link.SetConnMaxLifetime(100)
	link.SetMaxIdleConns(100)

	if err := link.Ping(); err != nil{
		panic(err)
	}
	DB = link
}