package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"   		
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

	output, err := json.MarshalIndent(&post, "", "\t\t")
	fmt.Println(output)
	if err != nil {
		fmt.Println("err:", err)
	}
	err = ioutil.WriteFile("post3.json", output, 0644)
	if err != nil {
		panic(err)
	}

}