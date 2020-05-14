package main

import (
	"fmt"
)

type Post struct {
	Id int
	Content string
	Author string
}

func test(data *Post) {
	// fmt.Println(data.Id)
	// data = Post{Id: 2, Content:"hello2", Author:"my2"}
	data.Id = 2
}

func (post *Post) test2() {
	// post.Id = 2
	test3(&post.Id)
}

func test3(id *int) {
	id = 3
}


func main() {
	post := Post{Id:1, Content:"hello", Author:"my"}
	// test(&post)
	post.test2()
	fmt.Println(post.Id)

	// var post1 Post
	// test(post1)

	// fmt.Println(&post)
}