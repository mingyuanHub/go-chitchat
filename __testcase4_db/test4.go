package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

type Post struct {
	Id int
	Content string
	Author string `sql:"not null"`
	Comments []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id int
	Content string
	Author string `sql:"not null"`
	PostId int `sql:"index"`
	CreatedAt time.Time
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=postgres dbname=gwp-gorm password=admin sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content:"hello1111111", Author:"mmmmmmmm"}
	fmt.Println(post)
	Db.Create(&post)
	fmt.Println(post)

	comment := Comment{Content:"Good Post", Author:"yyyyyy"}
	Db.Model(&post).Association("Comments").Append(comment)

	var readPost Post
	Db.Where("author = $1", "mmmmmmmm1").First(&readPost)

	fmt.Println(readPost)

	var comments []Comment
	Db.Model(&readPost).Related(&comments)

	fmt.Println(comments[0])
}