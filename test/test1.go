package main

import (
	"fmt"
)

type Post struct {
	Id int
	Content string
	Author string
}

func test(data interface{}) {
	// fmt.Println(data.Id)
	data = Post{Id: 2, Content:"hello2", Author:"my2"}
}

func test2(n int) {
	n = n + 1
	// fmt.Println(n)
}

func main() {
	// post := Post{Id:1, Content:"hello", Author:"my"}
	// test(post)

	// var post1 Post
	// test(&post1)

	// fmt.Println(&post1)

	m := 1
	test2(&m)
	fmt.Println(m)
}