/*
* @Author: Administrator
* @Date:   2019-06-30 12:01:57
* @Last Modified by:   Administrator
* @Last Modified time: 2019-07-20 19:18:01
*/
package main

import (
	"database/sql"
	// "github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


func main() {
	 db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
    // beego.Run()
	if err != nil {
	    log.Println(err)
	}



	

	log.Println(db)
	insert(db)

	rows, err := db.Query("select id, name from user where id = ?", 1)

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	var id int
	var name string

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func insert(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO user(name, age) VALUES(?, ?)")
	defer stmt.Close()
	if err != nil {
		log.Println(err)
		return
	}
	stmt.Exec("jeanron",12)
	stmt.Exec("jianrong",23)
}