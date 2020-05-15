package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"strconv"
)

type Post struct {
	Id int
	Content string
	Author string
}

func main() {
	//创建
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	//写入
	allPosts := []Post{
		Post{Id:1, Content:"hello1", Author:"wo1"},
		Post{Id:2, Content:"hello2", Author:"wo2"},
		Post{Id:3, Content:"hello3", Author:"wo3"},
		Post{Id:4, Content:"hello4", Author:"wo4"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err !=  nil {
			panic(err)
		}
	}

	writer.Flush()

	//读取
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content:item[1], Author:item[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)





}
