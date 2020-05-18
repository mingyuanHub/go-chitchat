package main

import (
	"encoding/json"
	"fmt"
	// "io"
	"os"		
)

type Post struct {
	Id int `json:"id"`
	Content string	`json:"content"`
	Author	Author
	Comments []Comment
}

type Author struct {
	Id int
	Name string
}

type Comment struct {
	Id	int
	Content string
	Author string
}

func main() {
	post := Post{
		Id: 1,
		Content:"hello",
		Author:Author{
			Id:2,
			Name:"mmt",
		},
		Comments:[]Comment{
			Comment{
				Id:3,
				Content:"hava a good dat",
				Author:"adin",
			},
			Comment{
				Id: 4,
				Content:"how are you today",
				Author:"liu",
			},
		},
	}

	jsonFile, err := os.Create("post.json4");
	if err != nil {
		fmt.Println("error create file:", err)
	}

	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("error recoder post", err)
	}

}