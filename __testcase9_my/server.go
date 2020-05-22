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

type Post struct {
	Id int
	Content string
	Author string
	CreateTime string
}

func get(id int) (post Post) {
	post1 := Post{}
	stmt, _ := db.Prepare("select id from post where id = ?")
	query := stmt.QueryRow(12)
	query.Scan(&post1.Id)
	fmt.Println(post1)


	post = Post{}
	err := db.QueryRow("select id, content, author, create_time from post where id = ?", id).Scan(&post.Id, &post.Content, &post.Author, &post.CreateTime)
	if err != nil {
		fmt.Println(err)
	}
	// return

	post = Post{}
	rows, err := db.Query("SELECT id FROM post where id < ?", id)
	if err != nil {        
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&post.Id)        
		if err != nil {
			panic(err)
		}
		fmt.Println(post.Id)
	}

	return
}

func (post *Post) update() {
	fmt.Println(post)
	res, err := db.Exec("update post set content = ? , author = ? , create_time = ?  where id = ?", post.Content, post.Author, post.CreateTime, post.Id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func (post *Post) add(){
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return
	}

	res, err := tx.Exec("insert into post (content, author, create_time) values (1, 2, '2020-05-21 18:22:22')")
	if err != nil {
		fmt.Println(err)
	}
	id, err := res.LastInsertId()
	post.Id = int(id)

	tx.Rollback()
}

func delete(id int) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return
	}
	_, err = tx.Exec("delete from post where id = ?", id)
	if err != nil {
		fmt.Println(err)
	}

	tx.Commit()
	// tx.Rollback()
}

func main() {
	post := Post{}
	post.add()
	// fmt.Println(post)
	post = get(post.Id)
	// fmt.Println(post)

	// post.Content += "8888"
	// post.update()
	// fmt.Println(post)

	delete(10)
}