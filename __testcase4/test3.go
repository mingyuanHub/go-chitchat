package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Post struct {
	Id int
	Content string
	AuthorName string `db:"author"`
}

var Db *sqlx.DB

func init() {
	var err error 
	Db, err = sqlx.Open("postgres", "user=postgres dbname=gwp password=admin sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func (post *Post) Create() (err error) {
	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.AuthorName).Scan(&post.Id)
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRowx("select id, content, author from posts where id = $1", id).StructScan(&post)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func main() {
	post := Post{Content:"hello111", AuthorName:"mmmmmm"}
	err := post.Create()
	fmt.Println(err)
	fmt.Println(post)
	fmt.Println(post.Id)

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost) 
}

