package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
    db, _ = sql.Open("mysql", "root:123456@tcp(39.100.133.182:3306)/goChitChat?charset=utf8")
}

func add() {
	_, err := db.Exec("insert into post (content, author, create_time) values (1, 2, '2020-05-21 18:22:22')")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	add()
}