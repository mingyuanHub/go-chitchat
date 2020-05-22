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
	post = Post{}
	err := db.QueryRow("select id, content, author, create_time from post where id = ?", id).Scan(&post.Id, &post.Content, &post.Author, &post.CreateTime)
	if err != nil {
		fmt.Println(err)
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
	res, err := db.Exec("insert into post (content, author, create_time) values (1, 2, '2020-05-21 18:22:22')")
	if err != nil {
		fmt.Println(err)
	}
	id, err := res.LastInsertId()
	post.Id = int(id)
}

func delete(id int) {
	_, err := db.Exec("delete from post where id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	post := Post{}
	post.add()
	fmt.Println(post)
	post = get(post.Id)
	fmt.Println(post)

	post.Content += "8888"
	post.update()
	fmt.Println(post)

	delete(3)
}